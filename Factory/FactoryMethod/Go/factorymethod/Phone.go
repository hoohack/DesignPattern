package factorymethod

type Phone struct {
  phoneProduct
  name string
}

func (this *Phone) GetName() string {
  return this.name
}
