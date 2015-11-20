package factorymethod

type IPhoneFactory struct {
  *PhoneFactory
}

func (this *IPhoneFactory) CreatePhone(ptype string) phoneProduct {
  if ptype == "6" {
    return NewIPhone6()
  } else if ptype == "6s"{
    return NewIPhone6s()
  } else {
    return nil
  }
}
