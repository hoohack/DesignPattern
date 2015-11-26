package main

import "./proxy"

func main() {
  idbq := new(proxy.DBQueryProxy)
  idbq.Request()
}
