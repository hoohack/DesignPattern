package templatemethod

type viewengine interface {
  PackData(data [2]string) string
}
