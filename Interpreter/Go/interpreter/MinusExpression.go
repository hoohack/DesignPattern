package interpreter

type MinusExpression struct {
  *SymbolExpression
}

func NewMinusExpression(Left Expression, Right Expression) *MinusExpression {
  return &MinusExpression{&SymbolExpression{Left: Left, Right: Right}}
}

func (this *MinusExpression) Calculate() int {
  return this.SymbolExpression.Left.Calculate() + this.SymbolExpression.Right.Calculate()
}
