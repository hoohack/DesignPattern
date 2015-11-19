package decorator

type DarkRoast struct {
	*Coffee
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{&Coffee{price: 2.0, description: "DarkRoast"}}
}

func (self *DarkRoast) Cost() float64 {
	return self.price
}

func (self *DarkRoast) GetDescription() string {
	return self.description
}
