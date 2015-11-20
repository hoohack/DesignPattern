package factorymethod

type XiaoMi struct {
  *Phone
}

func NewXiaoMi() *XiaoMi {
  return &XiaoMi{&Phone{name: "Xiaomi"}}
}
