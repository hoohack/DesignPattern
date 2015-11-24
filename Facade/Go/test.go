package main

import "fmt"
import "./facade"

func main() {
  store_dao := &facade.StoreDao{}
  goods_dao := &facade.GoodsDao{}

  goods_service := &facade.ServiceFacade{store_dao, goods_dao}
  goods_detail := goods_service.GoodsDetail()

  for i := range goods_detail {
    fmt.Printf("%s\n", goods_detail[i])
  }
}
