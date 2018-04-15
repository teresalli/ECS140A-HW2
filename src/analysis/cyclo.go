package analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type visitor struct {
  Complexity uint
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch n := node.(type) {
	case *ast.IfStmt:
		v.Complexity++
	case *ast.ForStmt, *ast.RangeStmt:
		v.Complexity++
	case *ast.CaseClause:
		if n.List != nil {
			v.Complexity++
		}
	case *ast.CommClause:
		if n.Comm != nil {
			v.Complexity++
		}
	}
	return v
}

func countStmt(start ast.Stmt) uint {
  v := new(visitor)
  v.Complexity = 1
  ast.Walk(v, start)
	return v.Complexity
}

func cyclomatic(node ast.Stmt) uint {
	var answer uint
	answer = 0
	if fn, ok := node.(*ast.BlockStmt); ok {
		for _, item := range fn.List {
			answer += countStmt(item)
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
