package chain

type LengthFilter struct {
}

func (this *LengthFilter) HandleFilter(str string) string {
  return str[0:20]
}
