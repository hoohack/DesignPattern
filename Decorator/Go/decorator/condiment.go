package decorator

type Condiment struct {
  *Coffee
  beverage beverage
  price float64
  description string
}

func (self *Condiment) Cost() float64 {
  return self.price
}

func (self *Condiment) GetDescription() string {
  return self.description
}
