package analysis

import (
	"bytes"
	"eval"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strconv"
)

type FuncVisitor struct {
}

func (v *FuncVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch call := node.(type) {
	case *ast.CallExpr:
		if name, ok := call.Fun.(*ast.SelectorExpr); ok {
			// function call found
			if x, ok := name.X.(*ast.Ident); ok {
				if name.Sel.Name == "ParseAndEval" && x.Name == "eval" {
					// not valid call
					if len(call.Args) == 1 {
						return v
					}
					if item, ok := call.Args[0].(*ast.BasicLit); ok {
						src, _ := strconv.Unquote(item.Value)
						expr, err := eval.Parse(src)
						// if valid string
						if err == nil {
							anw := expr.Simplify(eval.Env{})
							item.Value = strconv.Quote(eval.Format(anw))
						}
					}
				}
			}
		}
	}
	return v
}

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	ast.Walk(new(FuncVisitor), node)
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}
	//ast.Print(fset, f)
	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
