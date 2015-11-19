package decorator

type Decaf struct {
	*Coffee
	price float64
	description string
}

func NewDecaf() *Decaf {
	return &Decaf{&Coffee{}, 2.5, "Decaf"}
}

func (self *Decaf) Cost() float64 {
	return self.price + self.Coffee.Cost()
}

func (self *Decaf) GetDescription() string {
	return self.description + self.Coffee.GetDescription()
}
