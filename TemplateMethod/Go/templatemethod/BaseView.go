package templatemethod

import "fmt"

type BaseView struct {
}

func (bv *BaseView) GetData() [2]string {
  return [2]string{"200", "success"}
}

func (self *BaseView) Render(viewengine viewengine) {
  data := self.GetData()
  result := viewengine.PackData(data)
  fmt.Printf(result + "\n")
}
