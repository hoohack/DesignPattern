package prototype

type LifeArticle struct {
  *Article
}

func (this *LifeArticle) Clone() ArticleApi {
  author := NewAuthor(this.author.GetName())
  this.SetAuthor(author)
  return this
}

func (this *LifeArticle) SetAuthor(author *Author) {
  this.Article.author = author
}

func (this *LifeArticle) GetAuthor() *Author {
  return this.Article.author
}
