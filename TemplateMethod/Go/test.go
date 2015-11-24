package main

import "./templatemethod"

func main() {
  json_view := &templatemethod.JsonView{}
  json_view.Render(json_view)

  html_view := &templatemethod.HtmlView{}
  html_view.Render(html_view)
}
