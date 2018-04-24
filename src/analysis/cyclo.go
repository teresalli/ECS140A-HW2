package analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
)


func cyclo(start ast.Stmt) uint {
  var answer uint
	switch fn := start.(type){
  case *ast.BlockStmt:
      answer = 1
		for _,item := range fn.List {
			answer *= cyclo(item)
		}
    return answer;
  case *ast.CaseClause:
		answer = 1
		for _,item := range fn.Body {
			answer *= cyclo(item)
		}
    return answer;
  case *ast.SwitchStmt:
		answer = 0;
		var flag = false
    if fn.Body.List != nil {
		  for _, item := range fn.Body.List {
			  answer += cyclo(item)
				if item1, ok := item.(*ast.CaseClause); ok{
					if item1.List == nil {
						flag = true
					}
				}
		  }
  	}
		if !flag {
			answer++
		}
    return answer;
	case  *ast.TypeSwitchStmt:
		answer = 0;
		var flag = false
    if fn.Body.List != nil {
		  for _, item := range fn.Body.List {
			  answer += cyclo(item)
				if item1,ok := item.(*ast.CaseClause); ok{
					if item1.List == nil {
						flag = true
					}
				}
		  }
  	}
		if !flag {
			answer++
		}
    return answer;
	case *ast.IfStmt:
		return cyclo(fn.Body) + cyclo(fn.Else);
	case *ast.ForStmt:
    return cyclo(fn.Body) + 1;
	case *ast.RangeStmt:
		return cyclo(fn.Body) + 1;
	default:
		return 1;
	}
}

func cyclomatic(node ast.Stmt) uint {
	return cyclo(node)
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
