package analysis

import (
  "go/ast"
  "go/parser"
  "go/token"
)

func cyclomatic(node ast.Stmt) uint {
  panic("TODO: implement this!")
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
