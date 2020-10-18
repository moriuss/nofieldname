package nofieldname

import (
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "nofieldname finds struct created without field name"

// Analyzer finds struct created without field name
var Analyzer = &analysis.Analyzer{
	Name: "nofieldname",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	filter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}

	inspect.Preorder(filter, func(n ast.Node) {
		c, ok := n.(*ast.CompositeLit)
		if !ok {
			return
		}
		if !isDeclTypeStruct(pass, c) {
			return
		}

		for _, e := range c.Elts {
			if _, ok := e.(*ast.KeyValueExpr); !ok {
				pass.Reportf(e.Pos(), "field name is missing")
			}
		}
	})

	return nil, nil
}

func isDeclTypeStruct(pass *analysis.Pass, c *ast.CompositeLit) bool {
	_, ok := analysisutil.Under(pass.TypesInfo.TypeOf(c)).(*types.Struct)

	return ok
}
