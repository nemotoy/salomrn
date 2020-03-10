package mreclen

import (
	"fmt"
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

const Doc = "mreclen is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, d := range f.Decls {
			// prints x to stdout for debug
			ast.Print(nil, d)
			if v, ok := d.(*ast.FuncDecl); ok {
				// verify v.Recv if not nil
				if v.Recv != nil {
					r := v.Recv.List[0].Names[0].Name
					if len(r) > 2 {
						fmt.Printf("[Warning] method's receiver(%s) length overs max length\n", r)
					}
				}
			}
		}
	}
	return nil, nil
}
