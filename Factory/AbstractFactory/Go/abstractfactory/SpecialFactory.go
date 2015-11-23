package abstractfactory

type SpecialFactory struct {
  Content string
}

func NewSpecialFactory(content string) *SpecialFactory {
  sf := new(SpecialFactory)
  sf.Content = content
  return sf
}

func (sf *SpecialFactory) GetContent() string {
  return sf.Content
}
