package observer

import "fmt"

type ObserverA struct {
}

func (oa ObserverA) Update(cs *Subject) {
  fmt.Printf("observer A value is " + cs.GetValue())
}
