package composite

type navigationcomponent interface {
  GetName() string
  Add(navigationcomponent navigationcomponent)
  Remove(navigationcomponent navigationcomponent)
  GetChild() string
}
