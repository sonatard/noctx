package reqwithoutctx

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"
)

func Run(pass *analysis.Pass) (interface{}, error) {
	usedReqs := usedReqs(pass)
	newReqs := requestsByNewRequest(pass)

	reports := requestWithoutContext(usedReqs, newReqs)

	report(pass, reports)

	return nil, nil
}

func requestWithoutContext(usedReqs map[string]*ssa.Extract, newReqs map[*ssa.Call]*ssa.Extract) []*Report {
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
