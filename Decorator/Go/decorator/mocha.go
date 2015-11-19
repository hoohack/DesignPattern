package decorator

type Mocha struct {
  *Condiment
}

func NewMocha(beverage beverage) *Mocha {
  return &Mocha{&Condiment{beverage: beverage, price: 0.2, description: "Mocha"}}
}

func (self *Mocha) Cost() float64 {
  return self.beverage.Cost() + self.price
}

func (self *Mocha) GetDescription() string {
  return self.beverage.GetDescription() + " with " + self.description
}
