package prototype

type LifeArticle struct {
  title string
  author *Author
}

func (this *LifeArticle) Clone() ArticleApi {
  life_article := &LifeArticle{title: this.title}
  life_article.SetAuthor(this.author.Clone())
  return life_article
}

func (this *LifeArticle) SetTitle(title string) {
  this.title = title
}

func (this *LifeArticle) GetTitle() string {
  return this.title
}

func (this *LifeArticle) SetAuthor(author *Author) {
  this.author = author
}

func (this *LifeArticle) GetAuthor() *Author {
  return this.author
}
