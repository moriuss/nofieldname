package nofieldname

import (
	"go/ast"

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

	inspect.Preorder(nil, func(n ast.Node) {
		c, ok := n.(*ast.CompositeLit)
		if !ok {
			return
		}
		if !isDeclTypeStruct(c) {
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

func isDeclTypeStruct(c *ast.CompositeLit) bool {
	switch t := c.Type.(type) {
	case *ast.Ident:
		if t.Obj == nil {
			return false
		}
		d, ok := t.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return false
		}
		if _, ok := d.Type.(*ast.StructType); !ok {
			return false
		}

		return true
	default:
		return false
	}
}
