package analysis

import (
  "go/ast"
  "go/parser"
  "go/token"
)

func branchCount(fn *ast.FuncDecl) uint {
  var answer uint
  answer = 0
  for _, item := range fn.Body.List {
      switch i := item.(type) {
      case *ast.ForStmt, *ast.RangeStmt:
        answer += 1
      case *ast.SelectStmt:
        answer += 1
      case *ast.SwitchStmt, *ast.TypeSwitchStmt:
        answer += 1
      case *ast.IfStmt:
        answer += 1 + ElseCount(i)
      }
    }
    return answer
}

func ElseCount(i *ast.IfStmt) uint {
  if i.Else == nil {
    return 0
  }
  return 1
}

func ComputeBranchFactors(src string) map[string]uint {
  fset := token.NewFileSet()
  f, err := parser.ParseFile(fset, "src.go", src, 0)
  if err != nil {
    panic(err)
  }

  m := make(map[string]uint)
  for _, decl := range f.Decls {
    switch fn := decl.(type) {
    case *ast.FuncDecl:
      m[fn.Name.Name] = branchCount(fn)
    }
  }

  return m
}
