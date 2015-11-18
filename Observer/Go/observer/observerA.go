package observer

import "fmt"

type ObserverA struct {
	value string
}

func (ob ObserverA) Update(val string) {
	ob.value = val
	ob.Display()
}

func (ob ObserverA) Display() {
	fmt.Printf("observer A value is " + ob.value + "\n")
}
