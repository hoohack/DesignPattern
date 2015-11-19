package decorator

type DarkRoast struct {
	*Coffee
	price float64
	description string
}

func NewDarkRoast() *DarkRoast {
	return &DarkRoast{&Coffee{}, 2.0, "DarkRoast"}
}

func (self *DarkRoast) Cost() float64 {
	return self.price + self.Coffee.Cost()
}

func (self *DarkRoast) GetDescription() string {
	return self.description + self.Coffee.GetDescription()
}
