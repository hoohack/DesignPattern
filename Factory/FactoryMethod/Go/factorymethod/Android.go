package factorymethod

type Android struct {
  *Phone
}

func NewAndroid() *Android {
  return &Android{&Phone{name: "Android"}}
}
