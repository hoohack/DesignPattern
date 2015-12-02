package interpreter

type ValueExpression struct {
  value int
}

func NewValueExpression(value int) *ValueExpression {
  return &ValueExpression{value}
}

func (this *ValueExpression) Calculate() int {
  return this.value
}
