package decorator

type Espresso struct {
  *Coffee
}

func NewEspresso() *Espresso {
  return &Espresso{&Coffee{price: 1.0, description: "Espresso"}}
}

func (self *Espresso) Cost() float64 {
  return self.price
}

func (self Espresso) GetDescription() string {
  return self.description
}
