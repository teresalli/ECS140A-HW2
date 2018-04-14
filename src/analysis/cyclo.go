package analysis

import (
  "go/ast"
  "go/parser"
  "go/token"
)

func countStmt(start ast.Stmt, answer uint) uint {
  switch start.(type) {
  case *ast.IfStmt:
    answer++
  default:
    answer++
  }
  return answer
}

func cyclomatic(node ast.Stmt) uint {
  var answer uint
  answer = 0
  if fn,ok := node.(*ast.BlockStmt); ok{
    for _, item := range fn.List {
      answer = countStmt(item, answer)
    }
  }
  return answer
}

func CyclomaticComplexity(src string) map[string]uint {
  fset := token.NewFileSet()
  f, err := parser.ParseFile(fset, "src.go", src, 0)
  if err != nil {
    panic(err)
  }

  m := make(map[string]uint)
  for _, decl := range f.Decls {
    switch fn := decl.(type) {
    case *ast.FuncDecl:
      m[fn.Name.Name] = cyclomatic(fn.Body)
    }
  }

  return m
}
