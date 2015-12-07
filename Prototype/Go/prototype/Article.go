package prototype

type Article struct {
  title string
  author *Author
}

func (this *Article) SetTitle(title string) {
  this.title = title
}

func (this *Article) GetTitle() string {
  return this.title
}
