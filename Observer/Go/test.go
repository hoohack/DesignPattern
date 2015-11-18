package main

import "./observer"

func main() {
	subject := new(observer.ConcreteSubject)
	observerA := new(observer.ObserverA)
	observerB := new(observer.ObserverB)

	subject.RegisterObserver(observerA)
	subject.RegisterObserver(observerB)

	subject.SetValue("5")
}
