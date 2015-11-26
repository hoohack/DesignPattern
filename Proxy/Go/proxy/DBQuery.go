package proxy

import "fmt"

type DBQuery struct {
  IDBQuery
}

func NewDBQuery() *DBQuery {
  fmt.Printf("Connecting...\n")
  return &DBQuery{}
}

func (this *DBQuery) Request() {
  fmt.Printf("Requesting...\n")
}
