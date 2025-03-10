// 配列をテンプレートに渡し、イテレータを実行する
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_3.html")
	daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	t.Execute(w, daysOfWeek)
}


func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
    server.ListenAndServe()
}
