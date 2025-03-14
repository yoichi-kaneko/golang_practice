// テンプレートエンジンでコンテキスト依存処理な出力を行う
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_08.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}


func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
    server.ListenAndServe()
}
