package builder

type TextBuilder struct {
  content string
}

func (this *TextBuilder) buildHeader(header string) {
  this.content += "header: " + header + "\n"
}

func (this *TextBuilder) buildBody(body string) {
  this.content += "body: " + body + "\n"
}

func (this *TextBuilder) buildFooter(footer string) {
  this.content += "footer: " + footer + "\n"
}

func (this *TextBuilder) GetResult() string {
  return this.content
}
