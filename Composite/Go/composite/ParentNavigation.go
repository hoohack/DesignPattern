package composite

type ParentNavigation struct {
  Childs []navigationcomponent
  *BaseNavigation
}

func NewParentNavigation(name string) *ParentNavigation {
  return &ParentNavigation{BaseNavigation: &BaseNavigation{Name: name}}
}

func (this *ParentNavigation) Add(navigationcomponent navigationcomponent) {
  this.Childs = append(this.Childs, navigationcomponent)
}

func (this *ParentNavigation) Remove(navigationcomponent navigationcomponent) {
  c_idx := -1
  for idx, c := range this.Childs {
    if c.GetName() == navigationcomponent.GetName() {
      c_idx = idx
    }
  }

  this.Childs = append(this.Childs[:c_idx], this.Childs[c_idx+1:]...)
}

func (this *ParentNavigation) GetChild() string {
  result := ""
  for _, c := range this.Childs {
    result += " " + c.GetName()
  }
  return result
}
