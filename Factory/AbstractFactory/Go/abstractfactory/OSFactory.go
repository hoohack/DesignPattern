package abstractfactory

type OSFactory struct {
  Name string
}

func NewOSFactory(name string) *OSFactory {
  osf := new(OSFactory)
  osf.Name = name
  return osf
}

func (this *OSFactory) GetName() string {
  return this.Name
}
