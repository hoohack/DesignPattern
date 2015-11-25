package iterator

type Goods struct {
  Name string
}

func NewGoods(name string) *Goods {
  return &Goods{Name: name}
}

func (this *Goods) GetName() string {
  return this.Name
}
