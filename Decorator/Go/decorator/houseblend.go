package decorator

type HouseBlend struct {
  *Coffee
  price float64
  description string
}

func NewHouseBlend() *HouseBlend {
  return &HouseBlend{&Coffee{}, 3.0, "HouseBlend"}
}

func (self *HouseBlend) Cost() float64 {
  return self.price + self.Coffee.Cost()
}

func (self *HouseBlend) GetDescription() string {
  return self.description + self.Coffee.GetDescription()
}
