package iterator

type GoodsList struct {
  goodses []*Goods
}

func (this *GoodsList) GetGoods(idx int) *Goods {
  if this.goodses[idx] != nil {
    return this.goodses[idx]
  }

  return nil
}

func (this *GoodsList) AddGoods(goods *Goods) int {
  this.goodses = append(this.goodses, goods)
  return this.Count()
}

func (this *GoodsList) RemoveGoods(goods *Goods) int {
  t_idx := -1
  for idx, g := range this.goodses {
    if g.GetName() == goods.GetName() {
      t_idx = idx
      break
    }
  }

  if t_idx != -1 {
    this.goodses = append(this.goodses[:t_idx], this.goodses[t_idx+1:]...)
  }

  return this.Count()
}

func (this *GoodsList) Count() int {
  return len(this.goodses)
}
