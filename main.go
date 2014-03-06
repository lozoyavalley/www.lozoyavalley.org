package main

import (
    "github.com/codegangsta/martini"
    "github.com/martini-contrib/render"
)

type Site struct {
    Title       string
    Description string
}

func NewSite() *Site {
    title := "Lozoya Valley"
    description := `The Lozoya Valley initiative aims to associate product and media professional
to launch and adapt a technology and creative ecosystem within a rural, natural environment,
encouraging a healthy and sustainable way of life.`
    return &Site{title, description}
}

func main() {
    m := martini.Classic()
    m.Use(martini.Static("assets"))
    m.Use(render.Renderer(render.Options{
        Layout: "layout",
    }))

    site := NewSite()

    m.Get("/", func(r render.Render) {
        r.HTML(200, "index", site)
    })

    m.Run()
}
