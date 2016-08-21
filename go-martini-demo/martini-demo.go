package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

type User struct {
  ID string
  Password string
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
  Directory: "templates", // Specify what path to load the templates from.
  Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
  Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
  IndentJSON: true, // Output human readable JSON
  IndentXML: true, // Output human readable XML
}))
  m.Get("/hello",func (r render.Render)  {
    r.HTML(200,"hello",nil)
  })
  m.Get("/user/:userId",func (parmas martini.Params, r render.Render)  {
    u := User {
      ID: parmas["userId"],
    }
    r.JSON(200,u)
  })
  m.Run()
}
