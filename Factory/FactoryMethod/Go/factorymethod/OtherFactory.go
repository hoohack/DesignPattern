package factorymethod

type OtherFactory struct {
  PhoneFactory
}

func (this *OtherFactory) CreatePhone() phoneProduct {
  return NewOther()
}
