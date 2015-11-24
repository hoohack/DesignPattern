package facade

type GoodsDao struct {
}

func (gd *GoodsDao) Info() string {
  return "goods_name: long jeans, goods_size: [29, 30], goods_color[blue, black]"
}
