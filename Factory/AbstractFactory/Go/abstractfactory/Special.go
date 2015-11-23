package abstractfactory

type Special struct {
  Content string
  SpecialI
}

func (self *Special) Create() string {
  return self.Content
}
