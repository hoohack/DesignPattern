package decorator

type Soy struct {
  *Condiment
}

func NewSoy(beverage beverage) *Soy {
  return &Soy{&Condiment{beverage: beverage, price: 1.2, description: "Soy"}}
}

func (self *Soy) Cost() float64 {
  return self.beverage.Cost() + self.price
}

func (self *Soy) GetDescription() string {
  return self.beverage.GetDescription() + " with " + self.description
}
