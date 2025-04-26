package ngfunc

import (
	"fmt"
	"go/types"
	"maps"
	"slices"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

func Run(pass *analysis.Pass) (interface{}, error) {
	ngFuncMessages := map[string]string{
		// net/http
		"net/http.Get":                    "must not be called. use net/http.NewRequestWithContext and (*net/http.Client).Do(*http.Request)",
		"net/http.Head":                   "must not be called. use net/http.NewRequestWithContext and (*net/http.Client).Do(*http.Request)",
		"net/http.Post":                   "must not be called. use net/http.NewRequestWithContext and (*net/http.Client).Do(*http.Request)",
		"net/http.PostForm":               "must not be called. use net/http.NewRequestWithContext and (*net/http.Client).Do(*http.Request)",
		"(*net/http.Client).Get":          "must not be called. use (*net/http.Client).Do(*http.Request)",
		"(*net/http.Client).Head":         "must not be called. use (*net/http.Client).Do(*http.Request)",
		"(*net/http.Client).Post":         "must not be called. use (*net/http.Client).Do(*http.Request)",
		"(*net/http.Client).PostForm":     "must not be called. use (*net/http.Client).Do(*http.Request)",
		"net/http.NewRequest":             "must not be called. use net/http.NewRequestWithContext",
		"(*net/http.Request).WithContext": "must not be called. use net/http.NewRequestWithContext",

		// database/sql
		"(*database/sql.DB).Exec":     "must not be called. use (*database/sql.DB).ExecContext",
		"(*database/sql.DB).Ping":     "must not be called. use (*database/sql.DB).PingContext",
		"(*database/sql.DB).Prepare":  "must not be called. use (*database/sql.DB).PrepareContext",
		"(*database/sql.DB).Query":    "must not be called. use (*database/sql.DB).QueryContext",
		"(*database/sql.DB).QueryRow": "must not be called. use (*database/sql.DB).QueryRowContext",
		"(*database/sql.Tx).Exec":     "must not be called. use (*database/sql.Tx).ExecContext",
		"(*database/sql.Tx).Prepare":  "must not be called. use (*database/sql.Tx).PrepareContext",
		"(*database/sql.Tx).Query":    "must not be called. use (*database/sql.Tx).QueryContext",
		"(*database/sql.Tx).QueryRow": "must not be called. use (*database/sql.Tx).QueryRowContext",
	}

	ngFuncs := typeFuncs(pass, slices.Collect(maps.Keys(ngFuncMessages)))
	if len(ngFuncs) == 0 {
		return nil, nil
	}

	reportFuncs := ngCalledFuncs(pass, ngFuncs, ngFuncMessages)
	report(pass, reportFuncs)

	return nil, nil
}

func ngCalledFuncs(pass *analysis.Pass, ngFuncs []*types.Func, ngFuncMessages map[string]string) []*Report {
	var reports []*Report

	ssa, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	if !ok {
		panic(fmt.Sprintf("%T is not *buildssa.SSA", pass.ResultOf[buildssa.Analyzer]))
	}

	for _, sf := range ssa.SrcFuncs {
		for _, b := range sf.Blocks {
			for _, instr := range b.Instrs {
				for _, ngFunc := range ngFuncs {
					if analysisutil.Called(instr, nil, ngFunc) {
						message := ngFuncMessages[ngFunc.FullName()]
						ngCalledFunc := NewReport(instr, ngFunc, message)
						reports = append(reports, ngCalledFunc)

						break
					}
				}
			}
		}
	}

	return reports
}
