package main

import (
    "fmt"

    "github.com/codegangsta/martini"
    "github.com/martini-contrib/render"
    "github.com/nicksnyder/go-i18n/i18n"
)

type Site struct {
    Lang        string
    Title       string
    Description string
    LearnMore   string
}

func NewSite(locale string) *Site {
    locale_path := fmt.Sprintf("locale/%s.json", locale)
    i18n.MustLoadTranslationFile(locale_path)
    T, _ := i18n.Tfunc(locale)

    lang := locale
    title := "Lozoya Valley"
    description := T("description")
    learnMore := T("learn_more")
    return &Site{lang, title, description, learnMore}
}

func main() {
    m := martini.Classic()
    m.Use(martini.Static("assets"))
    m.Use(render.Renderer(render.Options{
        Layout: "layout",
    }))

    m.Get("/", func(r render.Render) {
        site := NewSite("en-US")
        r.HTML(200, "index", site)
    })

    m.Get("/es", func(r render.Render) {
        site := NewSite("es-ES")
        r.HTML(200, "index", site)
    })

    m.NotFound(func(r render.Render) {
        site := NewSite("en-US")
        r.HTML(404, "404", site)
    })

    m.Run()
}
