package observer

import "fmt"

type ObserverB struct {
	value string
}

func (ob ObserverB) Update(val string) {
	ob.value = val
	ob.Display()
}

func (ob ObserverB) Display() {
	fmt.Printf("observer B value is " + ob.value + "\n")
}
