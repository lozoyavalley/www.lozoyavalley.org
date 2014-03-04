package main

import (
    "github.com/codegangsta/martini"
    "github.com/martini-contrib/render"
)

func main() {
    m := martini.Classic()
    m.Use(martini.Static("assets"))
    m.Use(render.Renderer(render.Options{
        Layout: "layout",
    }))

    m.Get("/", func(r render.Render) {
        r.HTML(200, "index", nil)
    })

    m.Run()
}
