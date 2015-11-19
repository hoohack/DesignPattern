package decorator

type Whip struct {
  *Condiment
  price float64
  description string
}

func NewWhip(beverage beverage) *Whip {
  return &Whip{&Condiment{beverage: beverage}, 0.9, "Whip"}
}

func (self *Whip) Cost() float64 {
  return self.beverage.Cost() + self.price
}

func (self *Whip) GetDescription() string {
  return self.beverage.GetDescription() + " with " + self.description
}
