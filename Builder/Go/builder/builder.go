package builder

type builder interface {
  buildHeader(header string)
  buildBody(body string)
  buildFooter(footer string)
}
