package decorator

type Mocha struct {
  *Condiment
  price float64
  description string
}

func NewMocha(beverage beverage) *Mocha {
  return &Mocha{&Condiment{beverage: beverage}, 0.2, "Mocha"}
}

func (self *Mocha) Cost() float64 {
  return self.beverage.Cost() + self.price
}

func (self *Mocha) GetDescription() string {
  return self.beverage.GetDescription() + " with " + self.description
}
