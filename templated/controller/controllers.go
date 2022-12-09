package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Id   int
	Name string //exported field since it begins with a capital letter
}

func SimpleIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}
func HttpFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templated/index.html")
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("hello template")
	tmpl, _ = tmpl.Parse("Top Student: {{.Id}} - {{.Name}}!")

	p := Student{Id: 1, Name: "Hadi"}

	tmpl.Execute(w, p)
}

func HtmlTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("/Users/tennetdepositoryfullstack1/Hadi/Golang/9dec2022html/templated/html/templated.html")
	tmpl, _ = tmpl.ParseFiles("/Users/tennetdepositoryfullstack1/Hadi/Golang/9dec2022html/templated/html/templated.html")
	fmt.Println("ppp")
	p := Student{Id: 1, Name: "Hadi"}

	tmpl.Execute(w, p)
}
