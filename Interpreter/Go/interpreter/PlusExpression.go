package interpreter

type PlusExpression struct {
  *SymbolExpression
}

func NewPlusExpression(Left Expression, Right Expression) *PlusExpression {
  return &PlusExpression{&SymbolExpression{Left: Left, Right: Right}}
}

func (this *PlusExpression) Calculate() int {
  return this.SymbolExpression.Left.Calculate() + this.SymbolExpression.Right.Calculate()
}
