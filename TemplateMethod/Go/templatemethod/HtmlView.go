package templatemethod

type HtmlView struct {
  viewengine
  *BaseView
}

func (hv *HtmlView) PackData(data [2]string) string {
  return "code: " + data[0] + " result: " + data[1]
}
