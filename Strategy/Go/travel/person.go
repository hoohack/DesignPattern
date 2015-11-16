package travel

type Person struct {
  _strategy TravelStrategy
}

func SetStrategy(strategy TravelStrategy) Person {
	person := new(Person)
  
  person._strategy = strategy

	return *person
}

func (p Person) Travel() string {
  return p._strategy.traveler();
}
