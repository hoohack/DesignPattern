package main

import "./composite"
import "fmt"

func main() {
  store_navigation := composite.NewParentNavigation("store")

  nike := composite.NewChildNavigation("nike")
  adidas := composite.NewChildNavigation("adidas")
  new_balance := composite.NewChildNavigation("new balance")
  asscis := composite.NewChildNavigation("asscis")

  store_navigation.Add(nike)
  store_navigation.Add(adidas)
  store_navigation.Add(new_balance)
  store_navigation.Add(asscis)

  fmt.Printf("%s\n", store_navigation.GetChild())

  food := composite.NewChildNavigation("food");
  food.Add(composite.NewChildNavigation("kfc"));
}
