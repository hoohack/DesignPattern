package decorator

type Whip struct {
  *Condiment
}

func NewWhip(beverage beverage) *Whip {
  return &Whip{&Condiment{beverage: beverage, price: 0.9, description: "Whip"}}
}

func (self *Whip) Cost() float64 {
  return self.beverage.Cost() + self.price
}

func (self *Whip) GetDescription() string {
  return self.beverage.GetDescription() + " with " + self.description
}
