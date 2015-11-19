package decorator

type Espresso struct {
  *Coffee
  price float64
  description string
}

func NewEspresso() *Espresso {
  return &Espresso{&Coffee{}, 1.0, "Espresso"}
}

func (self *Espresso) Cost() float64 {
  return self.price + self.Coffee.Cost()
}

func (self Espresso) GetDescription() string {
  return self.description + self.Coffee.GetDescription()
}
