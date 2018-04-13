package eval

func (v Var) Simplify(env Env) Expr {
  if _, ok := env[v]; ok{
    return Literal(env[v])
  }
  return v
}

func (f Literal) Simplify(env Env) Expr {
  return f
}

func (u unary) Simplify(env Env) Expr {
  u.x = u.x.Simplify(env)
  return u
}

func (b binary) Simplify(env Env) Expr {
  b.x = b.x.Simplify(env)
  b.y = b.y.Simplify(env)
  switch x := b.x.(type) {
  case Literal:
    switch b.y.(type) {
    case Literal:
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
