package main

import "./prototype"
import "fmt"

func main() {
  life_article := &prototype.LifeArticle{}
  tech_article := &prototype.TechArticle{}
  life_article.SetTitle("Life")
  tech_article.SetTitle("techology")
  author_hector := prototype.NewAuthor("hector")
  author_tom := prototype.NewAuthor("tom")
  life_article.SetAuthor(author_hector)
  tech_article.SetAuthor(author_tom)

  clone_life := life_article.Clone()
  clone_tech := tech_article.Clone()
  clone_life.SetTitle("Programmer life")

  fmt.Printf("%s\n", clone_life.GetTitle())
  fmt.Printf("%s\n", life_article.GetTitle())
  fmt.Printf("%s\n", clone_tech.GetTitle())

  clone_life.GetAuthor().SetName("haha")
  fmt.Printf("%s\n", clone_life.GetAuthor().GetName())
  fmt.Printf("%s\n", life_article.GetAuthor().GetName())
}
