package factorymethod

type AndroidFactory struct {
  *PhoneFactory
}

func (this *AndroidFactory) CreatePhone(ptype string) phoneProduct {
  if ptype == "samsung" {
    return NewSamsung()
  } else if ptype == "xiaomi"{
    return NewXiaoMi()
  } else {
    return nil
  }
}
