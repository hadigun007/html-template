package main

import "net/http"

func main() {
	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, `<h1 style="color:green;">Go world !</h1>`)
	// 	fmt.Fprintf(w, "<p>Golang is a wonderful languagepppp</p>")

	// })
	// http.ListenAndServe(":3030", nil)
	static := http.FileServer(http.Dir("./static"))
	http.Handle("/", static)

	http.ListenAndServe(":8081", nil)
}
