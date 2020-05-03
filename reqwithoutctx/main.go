package reqwithoutctx

import (
	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

func Run(pass *analysis.Pass) (interface{}, error) {
	newRequestType := analysisutil.TypeOf(pass, "net/http", "NewRequest")
	requestType := analysisutil.TypeOf(pass, "net/http", "*Request")

	analyzer := &analyzer{
		pass:           pass,
		newRequestType: newRequestType,
		requestType:    requestType,
	}

	reports := analyzer.Exec()

	report(pass, reports)

	return nil, nil
}
