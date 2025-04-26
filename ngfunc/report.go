package ngfunc

import (
	"fmt"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ssa"
)

type Report struct {
	instruction ssa.Instruction
	function    *types.Func
	message     string
}

func NewReport(instruction ssa.Instruction, function *types.Func, message string) *Report {
	return &Report{
		instruction: instruction,
		function:    function,
		message:     message,
	}
}

func (n *Report) Pos() token.Pos {
	return n.instruction.Pos()
}

func (n *Report) Message() string {
	return fmt.Sprintf("%s %s", n.function.FullName(), n.message)
}

func report(pass *analysis.Pass, reports []*Report) {
	for _, report := range reports {
		pass.Reportf(report.Pos(), "%s", report.Message())
	}
}
