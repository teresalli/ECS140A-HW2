package analysis

import (
  // "fmt"
  "testing"
)

//!+CyclomaticComplexity
func TestCyclomaticComplexity(t *testing.T) {
  var test_code = `
    package main

    import (
      "fmt"
      "eval"
    )

    func f() {
      for x := range []uint{1, 2, 3, 4} {
        if x == 1 {
          fmt.Println("A good number")
        } else {
          fmt.Println("Not so great")
        }
      }
    }

    func g(x int) {
      if x < 0 {
        x = 0;
      } else if x > 0 {
        x = 1;
      }

      if x == 1 {
        x = 100
      }

      return x;
    }

    func h() {
      switch 5 {
      case 0:
        fmt.Println("It's zero!")
      case 5:
        fmt.Println("It's five!")
      default:
        fmt.Println("It isn't five...")
      }
    }

    func foo(x uint) uint {
      if x == 4 {
        return 2
      }else {
        if x == 5 {
          return 42
        } else {
          return 101
        }
      }
    }

    func fooo() {
      for x := 0; x < 5; x += 1 {
        if x == 1 {
          fmt.Println("A good number")
        } else {
          fmt.Println("Not so great")
        }
      }
    }

    func typee() {
      var x = 0
      switch x.(type) {
      case int:
        fmt.Println("It's int!")
      case string:
        fmt.Println("It's string!")
      default:
        fmt.Println("It's !")
      }
    }

    func hh() {
      switch 5 {
      }
    }


    func hhh() {
      var x = 0
      switch x.(type) {
      case int:
        fmt.Println("It's int!")
      case string:
        fmt.Println("It's string!")
      }
    }
`


  tests := []struct {
    name string
    cyclo uint
  }{
    {"f", 3},
    {"g", 6},
    {"h", 3},
    {"foo", 3},
    {"fooo", 3},
    {"typee", 3},
    {"hh", 1},
    {"hhh", 3},
  }

  cyclos := CyclomaticComplexity(test_code)

  for _, test := range tests {
    if cyclos[test.name] != test.cyclo {
      t.Errorf("CyclomaticComplexity()[%v] = %d, want %d\n",
        test.name, cyclos[test.name], test.cyclo)
    }
  }
}

func TestCyclomaticComplexity_Failure(t *testing.T) {
  var test_code = `
    package main

    func f() {
      x := 42
      fmt.Println(x
      return x
    }

  `
      defer func() {
        if recover() == nil {
          t.Errorf("ComputeBranchFactors() did not panic, but should\n")
        }
      }()

      CyclomaticComplexity(test_code)
}
//!-CyclomaticComplexity
