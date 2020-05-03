package reqwithoutctx

import (
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

type analyzer struct {
	pass           *analysis.Pass
	newRequestType types.Type
	requestType    types.Type
}

func (a *analyzer) Exec() []*Report {
	usedReqs := a.usedReqs()
	newReqs := a.requestsByNewRequest()

	return a.report(usedReqs, newReqs)
}

func (a *analyzer) report(usedReqs map[string]*ssa.Extract, newReqs map[*ssa.Call]*ssa.Extract) []*Report {
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

func (a *analyzer) usedReqs() map[string]*ssa.Extract {
	reqExts := make(map[string]*ssa.Extract)

	srcFuncs := a.pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, sf := range srcFuncs {
		for _, b := range sf.Blocks {
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

func (a *analyzer) usedReqByCall(call *ssa.Call) []*ssa.Extract {
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

func (a *analyzer) usedReqByUnOp(op *ssa.UnOp) *ssa.Extract {
	if ext, ok := op.X.(*ssa.Extract); ok && types.Identical(ext.Type(), a.requestType) {
		return ext
	}

	return nil
}

func (a *analyzer) usedReqByReturn(ret *ssa.Return) []*ssa.Extract {
	rets := ret.Results
	exts := make([]*ssa.Extract, 0, len(rets))

	for _, ret := range rets {
		if ext, ok := ret.(*ssa.Extract); ok && types.Identical(ext.Type(), a.requestType) {
			exts = append(exts, ext)
		}
	}

	return exts
}

func (a *analyzer) requestsByNewRequest() map[*ssa.Call]*ssa.Extract {
	reqs := make(map[*ssa.Call]*ssa.Extract)

	srcFuncs := a.pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, f := range srcFuncs {
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				if ext, ok := instr.(*ssa.Extract); ok {
					if types.Identical(ext.Type(), a.requestType) {
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
		}
	}

	return reqs
}
