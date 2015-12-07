package main

import "./prototype"
import "fmt"

func main() {
  life_article := prototype.NewLifeArticle("test")
  author_hector := prototype.NewAuthor("hector")
  life_article.SetAuthor(author_hector)

  article := life_article.Clone()
  article.SetTitle("Programmer life")

  fmt.Printf("%s\n", article.GetTitle())
  article.GetAuthor().SetName("haha")
  // fmt.Printf("%s\n", article.GetAuthor.GetName())
  // fmt.Printf("%s\n", life_article.GetAuthor.GetName())
  // fmt.Printf("%s\n", life_article.GetTitle())
}
