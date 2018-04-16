package eval

import (
  "fmt"
  "testing"
)

//!+Simplify
func TestSimplify(t *testing.T) {
  tests := []struct {
    expr string
    env  Env
    want string
  } {
    {"0 * X", Env{}, "0"},
    {"X * 0", Env{}, "0"},
    {"0 + X", Env{}, "X"},
    {"X + 0", Env{}, "X"},
    {"1 * X", Env{}, "X"},
    {"X * 1", Env{}, "X"},
    {"-X", Env{}, "(-X)"},
    {"-X", Env{"X": 1}, "-1"},
    {"+X", Env{"X": 5}, "5"},
    {"-(X + X)", Env{"X": 1}, "-2"},
    {"10 / X", Env{"X": 2}, "5"},
    {"(X + X) - Y", Env{"X": 2}, "(4 - Y)"},
    {"(X + X) - Y", Env{"Y": 8}, "((X + X) - 8)"},
    {"5 + 2", Env{}, "7"},
    {"10 - 1 + X - Y", Env{}, "((9 + X) - Y)"},
    {"X + 3 + 5", Env{}, "((X + 3) + 5)"},
  }

  for _, test := range tests {
    expr, err := Parse(test.expr)
    if err != nil {
      t.Error(err) // parse error
      continue
    }

    fmt.Printf("\n%s\n", test.expr)

    // Run the method
    result := expr.Simplify(test.env)

    // Display the result
    got := Format(result)
    fmt.Printf("\t%s, %v => %s\n", Format(expr), test.env, got)

    // Check the result
    if got != test.want {
      t.Errorf("(%s).Simplify() in %v = %q, want %q\n",
        test.expr, test.env, got, test.want)
    }
  }
}

//Test failure case
func TestSimplify_Failure(t *testing.T) {
  tests := []struct {
    expr string
    env  Env
  } {
    {"km(km(10))", Env{}},
    {"mi(10)", Env{}},
  }

  for _, test := range tests {
    expr, err := Parse(test.expr)
    if err != nil {
      t.Error(err) // parse error
      continue
    }
    func() {
      defer func() {
        if recover() == nil {
          t.Errorf("(%s).Simplify() did not panic, but should\n", test.expr)
        }
      }()
      // The following is the code under test
      expr.Simplify(test.env)
    }()
  }
}
//!-Simplify
