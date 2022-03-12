package exitanalyser

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const exitName = "Exit"

var ExitAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check for os.Exit calls",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.SelectorExpr:
				if x.Sel.Name == exitName {
					pass.Reportf(x.Pos(), "os.Exit called")
				}
			}
			return true
		})
	}
	return nil, nil
}
