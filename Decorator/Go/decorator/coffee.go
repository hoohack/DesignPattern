package decorator

type Coffee struct {
  beverage
  price float64
  description string
}

func (self *Coffee) Cost() float64 {
  return self.price
}

func (self *Coffee) GetDescription() string {
  return self.description
}
