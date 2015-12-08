package prototype

type Author struct {
  name string
}

func NewAuthor(name string) *Author {
  return &Author{name}
}

func (this *Author) SetName(name string) {
  this.name = name
}

func (this *Author) GetName() string {
  return this.name
}

func (this *Author) Clone() *Author {
  return &Author{name: this.name}
}
