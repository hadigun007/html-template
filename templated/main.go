package main

import (
	"fmt"
	"net/http"

	"github.com/hadigun007/html-template/templated/controller"
)

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/", controller.SimpleIndexHandler)
	http.HandleFunc("/index", controller.HttpFileHandler)
	http.HandleFunc("/tmpl", controller.TemplateHandler)
	http.HandleFunc("/htm", controller.HtmlTemplateHandler)

	http.ListenAndServe(":8081", nil)
}
