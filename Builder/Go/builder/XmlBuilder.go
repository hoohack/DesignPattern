package builder

type XmlBuilder struct {
  content string
}

func (this *XmlBuilder) buildHeader(header string) {
  this.content += "<Header>" + header + "</Header>\n"
}

func (this *XmlBuilder) buildBody(body string) {
  this.content += "<Body>" + body + "</Body>\n"
}

func (this *XmlBuilder) buildFooter(footer string) {
  this.content += "<Footer>" + footer + "</Footer>\n"
}

func (this *XmlBuilder) GetResult() string {
  return this.content
}
