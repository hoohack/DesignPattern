package composite

import "fmt"

type BaseNavigation struct {
  Name string
  navigationcomponent
}

func (bn *BaseNavigation) Add(navigationcomponent navigationcomponent) {
  fmt.Printf("Unsupported operation\n")
}

func (bn *BaseNavigation) Remove(navigationcomponent navigationcomponent) {
  fmt.Printf("Unsupported operation\n")
}

func (bn *BaseNavigation) GetChild() string {
  fmt.Printf("Unsupported operation\n")
  return ""
}

func (this *BaseNavigation) GetName() string {
  return this.Name
}
