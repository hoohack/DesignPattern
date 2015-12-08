package prototype

type ArticleApi interface {
  Clone() ArticleApi
  GetTitle() string
  SetTitle(name string)
  SetAuthor(author *Author)
  GetAuthor() *Author
}
