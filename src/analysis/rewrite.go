package analysis

import (
  "go/ast"
  "go/parser"
  "go/token"
  "go/format"
  "bytes"

  // "eval"
  // "strconv"
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
  
}

func SimplifyParseAndEval(src string) string {
  fset := token.NewFileSet()
  f, err := parser.ParseFile(fset, "src.go", src, 0)
  if err != nil {
    panic(err)
  }

  rewriteCalls(f)

  var buf bytes.Buffer
  format.Node(&buf, fset, f)
  return buf.String()
}
