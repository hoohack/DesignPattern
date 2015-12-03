package interpreter

type MulExpression struct {
  *SymbolExpression
}

func NewMulExpression(Left Expression, Right Expression) *MulExpression {
  return &MulExpression{&SymbolExpression{Left: Left, Right: Right}}
}

func (this *MulExpression) Calculate() int {
  return this.SymbolExpression.Left.Calculate() * this.SymbolExpression.Right.Calculate()
}
