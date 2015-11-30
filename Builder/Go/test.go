package main

import "./builder"
import "fmt"

func main() {
  header := "title->test"
  body := "builder"
  footer := "end"

  text_builder := &builder.TextBuilder{}
  text_director := builder.NewDirector(text_builder)
  text_director.Construct(header, body, footer)
  fmt.Println(text_builder.GetResult())

  xml_builder := &builder.XmlBuilder{}
  xml_director := builder.NewDirector(xml_builder)
  xml_director.Construct(header, body, footer)
  fmt.Println(xml_builder.GetResult())
}
