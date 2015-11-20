package factorymethod

type Samsung struct {
  *Phone
}

func NewSamsung() *Samsung {
  return &Samsung{&Phone{name: "Samsung"}}
}
