package template

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed templates/components/*.html
var templateComponents embed.FS

//go:embed templates/pages/*.html
var templatePages embed.FS

//go:embed static
var Statics embed.FS

type Engine struct {
	templates map[string]*template.Template
	funcMap   *funcMap
}

func NewEngine(basePath string) *Engine {
	return &Engine{
		templates: make(map[string]*template.Template),
		funcMap:   &funcMap{basePath: basePath},
	}
}

func (e *Engine) ParseTemplates() {
	funcMap := e.funcMap.Map()
	templates := map[string][]string{
		"index.html": {"posts.html", "filters.html", "generals.html", "profile.html", "iconic.html", "login.html", "comment.html"},
	}

	for name, components := range templates {
		tpl := template.New("").Funcs(funcMap)
		for _, component := range components {
			template.Must(tpl.ParseFS(templateComponents, "templates/components/"+component))
		}
		e.templates[name] = template.Must(tpl.ParseFS(templatePages, "templates/pages/"+name))
	}

	if pages, err := templatePages.ReadDir("templates/pages"); err == nil {
		for _, page := range pages {
			if _, ok := e.templates[page.Name()]; !ok {
				panic("Template " + page.Name() + " isn't declared in ParseTemplates")
			}
		}
	} else {
		panic("Unable to read all embedded pages templates")
	}
}

func (e *Engine) Render(name string, data map[string]any) []byte {
	tpl, ok := e.templates[name]
	if !ok {
		panic("The template " + name + " does not exists.")
	}

	var b bytes.Buffer
	if err := tpl.ExecuteTemplate(&b, "base", data); err != nil {
		panic(err)
	}

	return b.Bytes()
}
