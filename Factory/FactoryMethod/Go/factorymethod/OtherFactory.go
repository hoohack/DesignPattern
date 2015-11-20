package factorymethod

type OtherFactory struct {
  *PhoneFactory
}

func (this *OtherFactory) CreatePhone(ptype string) phoneProduct {
  if ptype == "nokia" {
    return NewNokia()
  } else if ptype == "blackberry"{
    return NewBlackBerry()
  } else {
    return nil
  }
}
