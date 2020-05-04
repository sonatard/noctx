package reqwithoutctx

import (
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/gostaticanalysis/analysisutil"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

type Analyzer struct {
	netHTTPImportFuncs []*ssa.Function
	newRequestType     types.Type
	requestType        types.Type
}

func NewAnalyzer(pass *analysis.Pass) *Analyzer {
	newRequestType := analysisutil.TypeOf(pass, "net/http", "NewRequest")
	requestType := analysisutil.TypeOf(pass, "net/http", "*Request")

	srcFuncs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	skipFile := make(map[*ast.File]bool)
	netHTTPImportFuncs := make([]*ssa.Function, 0, len(srcFuncs))

	for _, f := range srcFuncs {
		if UnImportedNetHTTP(pass, f, skipFile) {
			continue
		}

		netHTTPImportFuncs = append(netHTTPImportFuncs, f)
	}

	return &Analyzer{
		netHTTPImportFuncs: netHTTPImportFuncs,
		newRequestType:     newRequestType,
		requestType:        requestType,
	}
}

func (a *Analyzer) Exec() []*Report {
	usedReqs := a.usedReqs()
	newReqs := a.requestsByNewRequest()

	return a.report(usedReqs, newReqs)
}

func (a *Analyzer) report(usedReqs map[string]*ssa.Extract, newReqs map[*ssa.Call]*ssa.Extract) []*Report {
	var reports []*Report

	for _, fReq := range usedReqs {
		for newRequest, req := range newReqs {
			if fReq == req {
				reports = append(reports, &Report{Instruction: newRequest})
			}
		}
	}

	return reports
}

func (a *Analyzer) usedReqs() map[string]*ssa.Extract {
	reqExts := make(map[string]*ssa.Extract)

	for _, f := range a.netHTTPImportFuncs {
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				switch i := instr.(type) {
				case *ssa.Call:
					exts := a.usedReqByCall(i)
					for _, ext := range exts {
						key := i.String() + ext.String()
						reqExts[key] = ext
					}
				case *ssa.UnOp:
					ext := a.usedReqByUnOp(i)
					if ext != nil {
						key := i.String() + ext.String()
						reqExts[key] = ext
					}
				case *ssa.Return:
					exts := a.usedReqByReturn(i)
					for _, ext := range exts {
						key := i.String() + ext.String()
						reqExts[key] = ext
					}
				}
			}
		}
	}

	return reqExts
}

func (a *Analyzer) usedReqByCall(call *ssa.Call) []*ssa.Extract {
	var exts []*ssa.Extract

	args := call.Common().Args
	if len(args) == 0 {
		return exts
	}

	for _, arg := range args {
		if ext, ok := arg.(*ssa.Extract); ok && types.Identical(ext.Type(), a.requestType) {
			// skip net/http.Request method call
			if strings.Contains(call.String(), "(*net/http.Request).") {
				continue
			}

			exts = append(exts, ext)
		}
	}

	return exts
}

func (a *Analyzer) usedReqByUnOp(op *ssa.UnOp) *ssa.Extract {
	if ext, ok := op.X.(*ssa.Extract); ok && types.Identical(ext.Type(), a.requestType) {
		return ext
	}

	return nil
}

func (a *Analyzer) usedReqByReturn(ret *ssa.Return) []*ssa.Extract {
	rets := ret.Results
	exts := make([]*ssa.Extract, 0, len(rets))

	for _, ret := range rets {
		if ext, ok := ret.(*ssa.Extract); ok && types.Identical(ext.Type(), a.requestType) {
			exts = append(exts, ext)
		}
	}

	return exts
}

func (a *Analyzer) requestsByNewRequest() map[*ssa.Call]*ssa.Extract {
	reqs := make(map[*ssa.Call]*ssa.Extract)

	for _, f := range a.netHTTPImportFuncs {
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				ext, ok := instr.(*ssa.Extract)
				if !ok {
					continue
				}

				if !types.Identical(ext.Type(), a.requestType) {
					continue
				}

				operands := ext.Operands([]*ssa.Value{})
				if len(operands) == 1 {
					operand := *operands[0]
					if f, ok := operand.(*ssa.Call); ok {
						if types.Identical(f.Call.Value.Type(), a.newRequestType) {
							reqs[f] = ext
						}
					}
				}
			}
		}
	}

	return reqs
}

func UnImportedNetHTTP(pass *analysis.Pass, f *ssa.Function, skipFile map[*ast.File]bool) (ret bool) {
	obj := f.Object()
	if obj == nil {
		return false
	}

	file := analysisutil.File(pass, obj.Pos())
	if file == nil {
		return false
	}

	if skip, has := skipFile[file]; has {
		return skip
	}

	defer func() {
		skipFile[file] = ret
	}()

	for _, impt := range file.Imports {
		path, err := strconv.Unquote(impt.Path.Value)
		if err != nil {
			continue
		}

		path = analysisutil.RemoveVendor(path)
		if path == "net/http" {
			return false
		}
	}

	return true
}
