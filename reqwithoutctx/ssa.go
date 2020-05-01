package reqwithoutctx

import (
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

func requestsByNewRequest(pass *analysis.Pass) map[*ssa.Call]*ssa.Extract {
	reqs := make(map[*ssa.Call]*ssa.Extract)

	newRequestType := analysisutil.TypeOf(pass, "net/http", "NewRequest")

	srcFuncs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, f := range srcFuncs {
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				if ext, ok := instr.(*ssa.Extract); ok {
					if ext.Type().String() == "*net/http.Request" {
						operands := ext.Operands([]*ssa.Value{})
						if len(operands) == 1 {
							operand := *operands[0]
							if f, ok := operand.(*ssa.Call); ok {
								if f.Call.Value.Type().String() == newRequestType.String() {
									reqs[f] = ext
								}
							}
						}
					}
				}
			}
		}
	}

	return reqs
}

func usedReqs(pass *analysis.Pass) map[string]*ssa.Extract {
	reqExts := make(map[string]*ssa.Extract)

	srcFuncs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, sf := range srcFuncs {
		for _, b := range sf.Blocks {
			for _, instr := range b.Instrs {
				switch i := instr.(type) {
				case *ssa.Call:
					exts := extractByCall(i)
					for _, ext := range exts {
						key := i.String() + ext.String()
						reqExts[key] = ext
					}
				case *ssa.UnOp:
					ext := extractByUnOp(i)
					if ext != nil {
						key := i.String() + ext.String()
						reqExts[key] = ext
					}
				case *ssa.Return:
					exts := extractByReturn(i)
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

func extractByCall(call *ssa.Call) []*ssa.Extract {
	var exts []*ssa.Extract

	fType := call.String()

	// skip net/http.Request method call
	if strings.Contains(fType, "net/http.Request).") {
		return exts
	}

	args := call.Common().Args
	if len(args) == 0 {
		return exts
	}

	for _, arg := range args {
		if ext, ok := arg.(*ssa.Extract); ok && strings.Contains(ext.Type().String(), "net/http.Request") {
			exts = append(exts, ext)
		}
	}

	return exts
}

func extractByUnOp(op *ssa.UnOp) *ssa.Extract {
	if ext, ok := op.X.(*ssa.Extract); ok && strings.Contains(ext.Type().String(), "net/http.Request") {
		return ext
	}

	return nil
}

func extractByReturn(ret *ssa.Return) []*ssa.Extract {
	rets := ret.Results
	exts := make([]*ssa.Extract, 0, len(rets))

	for _, ret := range rets {
		if ext, ok := ret.(*ssa.Extract); ok && strings.Contains(ext.Type().String(), "net/http.Request") {
			exts = append(exts, ext)
		}
	}

	return exts
}
