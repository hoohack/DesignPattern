package factorymethod

type BlackBerry struct {
  *Phone
}

func NewBlackBerry() *BlackBerry {
  return &BlackBerry{&Phone{name: "BlackBerry"}}
}
