package chain

import "strings"

type HTMLFilter struct {
}

func (this *HTMLFilter) HandleFilter(str string) string {
  result := strings.Replace(str, "<", "&lt;", -1)
  result = strings.Replace(result, ">", "&gt;", -1)
  return result
}
