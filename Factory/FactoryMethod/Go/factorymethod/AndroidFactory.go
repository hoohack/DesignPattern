package factorymethod

type AndroidFactory struct {
  PhoneFactory
}

func (this *AndroidFactory) CreatePhone() phoneProduct {
  return NewAndroid()
}
