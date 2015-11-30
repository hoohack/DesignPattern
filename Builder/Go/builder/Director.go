package builder

type Director struct {
  builder builder
}

func NewDirector(builder builder) *Director {
  return &Director{builder: builder}
}

func (this *Director) Construct(header string, body string, footer string) {
  this.builder.buildHeader(header)
  this.builder.buildBody(body)
  this.builder.buildFooter(footer)
}
