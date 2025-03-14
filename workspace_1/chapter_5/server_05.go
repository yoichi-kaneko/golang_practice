// テンプレートエンジンで値の代入を行う
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_05_t1.html", "tmpl_05_t2.html")
	t.Execute(w, "Hello World!!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
    server.ListenAndServe()
}
