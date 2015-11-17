package observer

import "fmt"

type ObserverB struct {
}

func (ob ObserverB) Update(cs *Subject) {
  fmt.Printf("observer B value is " + cs.GetValue())
}
