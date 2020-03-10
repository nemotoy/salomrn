package mreclen

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "mreclen",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const (
	Doc       = "mreclen is ..."
	defMaxLen = 2
)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, d := range f.Decls {
			if v, ok := d.(*ast.FuncDecl); ok {
				// verify v.Recv if not nil
				if v.Recv != nil {
					r := v.Recv.List[0].Names[0].Name
					if len(r) > defMaxLen {
						pass.Reportf(v.Pos(), "%s should be a one or two letter", r)
					}
				}
			}
		}
	}
	return nil, nil
}
