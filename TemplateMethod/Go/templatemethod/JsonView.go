package templatemethod

import "encoding/json"

type JsonView struct {
  *BaseView
}

type JData struct {
  Code string `json:"code"`
  Result string `json:"result"`
}

func (self *JsonView) PackData(data [2]string) string {
  var j_data JData
  j_data.Code = data[0]
  j_data.Result = data[1]
  js, _ := json.Marshal(j_data)
  return string(js[:len(js)])
}
