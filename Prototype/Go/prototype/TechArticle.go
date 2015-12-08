package prototype

type TechArticle struct {
  title string
  author *Author
}

func (this *TechArticle) Clone() ArticleApi {
  return &TechArticle{title: this.title, author: this.author}
}

func (this *TechArticle) SetTitle(title string) {
  this.title = title
}

func (this *TechArticle) GetTitle() string {
  return this.title
}

func (this *TechArticle) SetAuthor(author *Author) {
  this.author = author
}

func (this *TechArticle) GetAuthor() *Author {
  return this.author
}
