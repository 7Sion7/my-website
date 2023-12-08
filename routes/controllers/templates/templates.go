package templates

import (
	"html/template"
)

type PageVariables struct {
	Title string
	Body  template.HTML
	Style template.CSS
}

const (
	baseTmpl = `
	{{define "base"}}
		<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" href="/static/style.css">
			<script src="/static/script.js" defer></script>
			<title>My Website</title>
		</head>
		<header>
			<nav id="navbar">
				<a class="links" href="/">Home</a>
				<a  class="links" id="about" href="/about">About</a>
				<div class="dropdown" id="projects-dropdown" onmouseleave="fadeOut()">
					<a class="links last" id="projects" href="/projects" onmouseover="fadeIn()">Projects</a>
					<div class="projects" id="dropdown">
						<a id="projects/recipes" class="fades" href="/projects/recipes" >Recipes</a>
						<a id="projects/forum"  class="fades" href="/projects/forum" >Forum</a>
						<a id="projects/currency-converter" class="fades" href="/projects/currency-converter" >Currency Converter</a>
						<a id="projects/artists-tracker" class="fades" href="/projects/artists-tracker" >Artist Tracker</a>
				</div>
			</nav>
		</header>
		{{template "content" .}}
		
		</html>
		{{end}}
`
	contentTmpl = `
		{{define "content"}}
	{{.Body}}
	{{end}}
	`
)

func RenderTemplate(name string) (*template.Template, error) {

	tplBase := template.Must(template.New("base").Parse(baseTmpl))

	tplContent, err := tplBase.Clone()
	if err != nil {
		return nil, err
	}
	tplContent, err = tplContent.Parse(contentTmpl)
	if err != nil {
		return nil, err
	}

	tplContent = tplContent.Lookup("base").Funcs(
		template.FuncMap{
			"html": func(s string) template.HTML { return template.HTML(s) },
			"css":  func(s string) template.CSS { return template.CSS(s) },
		},
	)

	return tplContent, nil
}
