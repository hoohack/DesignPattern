package prototype

type LifeArticle struct {
  *Article
}

func NewLifeArticle(title string) *LifeArticle {
  return &LifeArticle{&Article{title: title}}
}

func (this *LifeArticle) Clone() ArticleApi {
  author := NewAuthor(this.author.GetName())
  title := this.Article.GetTitle()
  return &LifeArticle{&Article{title: title, author: author}}
}

func (this *LifeArticle) SetAuthor(author *Author) {
  this.Article.author = author
}

func (this *LifeArticle) GetAuthor() *Author {
  return this.Article.author
}
