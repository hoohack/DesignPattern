package factorymethod

type IPhoneFactory struct {
  PhoneFactory
}

func (this *IPhoneFactory) CreatePhone() phoneProduct {
  return NewIPhone()
}
