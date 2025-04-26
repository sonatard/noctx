package noctx

import (
	"github.com/sonatard/noctx/ngfunc"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var Analyzer = &analysis.Analyzer{
	Name:             "noctx",
	Doc:              "noctx finds sending http request without context.Context",
	Run:              ngfunc.Run,
	RunDespiteErrors: false,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
	ResultType: nil,
	FactTypes:  nil,
}
