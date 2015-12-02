package interpreter

type SymbolExpression struct {
  Expression
  Left Expression
  Right Expression
}

func NewSymbolExpression(Left Expression, Right Expression) *SymbolExpression {
  return &SymbolExpression{Left: Left, Right: Right}
}
