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
					// prints method name and it recever name
					fmt.Println(v.Name, v.Recv.List[0].Names[0])
				}
			}
		}
	}
	return nil, nil
}
