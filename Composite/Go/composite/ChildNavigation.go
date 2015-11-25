package composite

type ChildNavigation struct {
  *BaseNavigation
}

func NewChildNavigation(name string) *ChildNavigation {
  return &ChildNavigation{&BaseNavigation{Name: name}}
}
