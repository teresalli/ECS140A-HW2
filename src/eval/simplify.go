package eval

func (v Var) Simplify(env Env) Expr {
  // check if v is in the provided environment
  if _, ok := env[v]; ok{
    // found in the environment, replace with the value
    return Literal(env[v])
  }
  return v
}

func (f Literal) Simplify(env Env) Expr {
  // already a Literal nothing more to do
  return f
}

func (u unary) Simplify(env Env) Expr {
  // first simplify the expression
  u.x = u.x.Simplify(env)
  // check now if the expression is a literal
  switch u.x.(type) {
  case Literal: // a literal so evaluate
    return Literal(u.Eval(env))
  }
  return u
}

func (b binary) Simplify(env Env) Expr {
  // first simplify the expressions
  b.x = b.x.Simplify(env)
  b.y = b.y.Simplify(env)
  // check their type to see if evaluation needed
  switch x := b.x.(type) {
  case Literal:
    switch b.y.(type) {
    case Literal: // both literals, evaluate
      return Literal(b.Eval(env))
    default:
      switch x {
      case 0:
        switch b.op{
        case '*':
          return Literal(0)
        case '+':
          return b.y
        }
      case 1:
        switch b.op{
          case '*':
            return b.y
        }
      }

    }
  default:
    switch y := b.y.(type) {
    case Literal:
      switch y {
      case 0:
        switch b.op{
        case '*':
          return Literal(0)
        case '+':
          return b.x
        }
      case 1:
        switch b.op{
        case '*':
          return b.x
        }
      }
    }
  }
  return b
}

func (m measure) Simplify(env Env) Expr {
  // don't need to implement this
  panic("cannot simplify measure expression")
}
