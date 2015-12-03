package interpreter

type ModExpression struct {
  *SymbolExpression
}

func NewModExpression(Left Expression, Right Expression) *ModExpression {
  return &ModExpression{&SymbolExpression{Left: Left, Right: Right}}
}

func (this *ModExpression) Calculate() int {
  return this.SymbolExpression.Left.Calculate() % this.SymbolExpression.Right.Calculate()
}
