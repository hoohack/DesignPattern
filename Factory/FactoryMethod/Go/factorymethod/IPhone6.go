package factorymethod

type IPhone6 struct {
  *Phone
}

func NewIPhone6() *IPhone6 {
  return &IPhone6{&Phone{name: "IPhone6"}}
}
