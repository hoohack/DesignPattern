package abstractfactory

type AndroidFactory struct {
  IngredientFactory
  PhoneFactory
}

func (af *AndroidFactory) CreateOS(name string) *AndroidOS {
  return &AndroidOS{&OS{Name: name}}
}

func (af *AndroidFactory) CreateSpecial(content string) *AndroidSpecial {
  return &AndroidSpecial{&Special{Content: content}}
}
