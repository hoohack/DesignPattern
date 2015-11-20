package factorymethod

type IPhone struct {
  *Phone
}

func NewIPhone() *IPhone {
  return &IPhone{&Phone{name: "IPhone"}}
}
