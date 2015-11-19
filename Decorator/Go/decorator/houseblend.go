package decorator

type HouseBlend struct {
  *Coffee
}

func NewHouseBlend() *HouseBlend {
  return &HouseBlend{&Coffee{price: 3.0, description: "HouseBlend"}}
}

func (self *HouseBlend) Cost() float64 {
  return self.price
}

func (self *HouseBlend) GetDescription() string {
  return self.description
}
