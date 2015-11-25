package iterator

type GoodsIterator struct {
  goods_list *GoodsList
  Current_idx int
}

func NewGoodsIterator(goods_list *GoodsList) *GoodsIterator {
  return &GoodsIterator{goods_list, 0}
}

func (this *GoodsIterator) Current() *Goods {
  return this.goods_list.GetGoods(this.Current_idx)
}

func (this *GoodsIterator) Next() {
  this.Current_idx++
}

func (this *GoodsIterator) Key() int {
  return this.Current_idx
}

func (this *GoodsIterator) Valid() bool {
  if this.Current_idx >= this.goods_list.Count() {
    return false
  }

  return true
}

func (this *GoodsIterator) Rewind() {
  this.Current_idx = 0
}
