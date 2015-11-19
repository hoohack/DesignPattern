package decorator

type Soy struct {
  *Condiment
  price float64
  description string
}

func NewSoy(beverage beverage) *Soy {
  return &Soy{&Condiment{beverage: beverage}, 1.2, "Soy"}
}

func (self *Soy) Cost() float64 {
  return self.beverage.Cost() + self.price
}

func (self *Soy) GetDescription() string {
  return self.beverage.GetDescription() + " with " + self.description
}
