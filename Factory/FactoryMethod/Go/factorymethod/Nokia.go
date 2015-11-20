package factorymethod

type Nokia struct {
  *Phone
}

func NewNokia() *Nokia {
  return &Nokia{&Phone{name: "Nokia"}}
}
