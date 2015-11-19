package decorator

type Decaf struct {
	*Coffee
}

func NewDecaf() *Decaf {
	return &Decaf{&Coffee{price: 2.5, description: "Decaf"}}
}

func (self *Decaf) Cost() float64 {
	return self.price
}

func (self *Decaf) GetDescription() string {
	return self.description
}
