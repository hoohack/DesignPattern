package observer

type ConcreteSubject struct {
	value     string
	observers []Observer
}

func (cs *ConcreteSubject) RegisterObserver(ob Observer) {
	cs.observers = append(cs.observers, ob)
}

func (cs *ConcreteSubject) Notify() {
	for _, observer := range cs.observers {
		observer.Update(cs.value)
	}
}

func (cs *ConcreteSubject) SetValue(val string) {
	cs.value = val
	cs.Notify()
}

func (cs *ConcreteSubject) GetValue() string {
	return cs.value
}
