package abstractfactory

type IPhoneFactory struct {
  IngredientFactory
  PhoneFactory
}

func (ipf *IPhoneFactory) CreateOS(name string) *IPhoneOS {
  return &IPhoneOS{&OS{Name: name}}
}

func (ipf *IPhoneFactory) CreateSpecial(content string) *IPhoneSpecial {
  return &IPhoneSpecial{&Special{Content: content}}
}
