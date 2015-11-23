package abstractfactory

type OS struct {
  Name string
  OSI
}

func (self *OS) Create() string {
  return self.Name
}
