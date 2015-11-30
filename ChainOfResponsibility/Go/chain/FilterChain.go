package chain

type FilterChain struct {
  filters []Filter
}

func NewFilterChain() *FilterChain {
  return &FilterChain{}
}

func (this *FilterChain) AddFilter(filter Filter) *FilterChain {
  this.filters = append(this.filters, filter)
  return this
}

func (this *FilterChain) HandleFilter(str string) string {
  for _, filter := range this.filters {
    str = filter.HandleFilter(str)
  }

  return str
}
