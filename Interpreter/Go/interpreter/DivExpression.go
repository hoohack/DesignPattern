package interpreter

type DivExpression struct {
  *SymbolExpression
}

func NewDivExpression(Left Expression, Right Expression) *DivExpression {
  return &DivExpression{&SymbolExpression{Left: Left, Right: Right}}
}

func (this *DivExpression) Calculate() int {
  return this.SymbolExpression.Left.Calculate() + this.SymbolExpression.Right.Calculate()
}
