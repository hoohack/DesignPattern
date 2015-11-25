package main

import "./iterator"
import "fmt"

func main() {
  goods_list := new(iterator.GoodsList)
  goods_list.AddGoods(iterator.NewGoods("tshirt"))
  goods_list.AddGoods(iterator.NewGoods("jean"))
  goods_list.AddGoods(iterator.NewGoods("phone"))
  goods_list.AddGoods(iterator.NewGoods("cup"))

  for goods_iterator := iterator.NewGoodsIterator(goods_list);goods_iterator.Valid(); {
    fmt.Printf("%s\n", goods_iterator.Current().GetName())
    goods_iterator.Next()
  }

}
