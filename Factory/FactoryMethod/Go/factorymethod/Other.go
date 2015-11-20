package factorymethod

type Other struct {
  *Phone
}

func NewOther() *Other {
  return &Other{&Phone{name: "Other"}}
}
