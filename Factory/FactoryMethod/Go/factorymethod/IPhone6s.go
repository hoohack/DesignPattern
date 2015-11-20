package factorymethod

type IPhone6s struct {
  *Phone
}

func NewIPhone6s() *IPhone6s {
  return &IPhone6s{&Phone{name: "IPhone6s"}}
}
