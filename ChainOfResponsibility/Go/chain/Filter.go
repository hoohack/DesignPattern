package chain

type Filter interface {
  HandleFilter(str string) string
}
