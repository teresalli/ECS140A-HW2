package analysis

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
  "strings"
  "fmt"
	"eval"
	"strconv"
)

type FuncVisitor struct {
}

func (v *FuncVisitor) Visit(node ast.Node) (w ast.Visitor)  {
  switch call := node.(type) {
  case *ast.CallExpr:
    if name,ok := call.Fun.(*ast.SelectorExpr); ok{
      // function call found
      if name.Sel.Name == "ParseAndEval" {
        // not valid call
        if len(call.Args) == 1 {
          return v
        }
        if item,ok := call.Args[0].(*ast.BasicLit); ok{
          src := strings.Trim(item.Value, "\"")
          num,err := eval.ParseAndEval(src, eval.Env{})
          if err == nil {
            item.Value = fmt.Sprintf("\"%s\"", strconv.Itoa(int(num)))
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
