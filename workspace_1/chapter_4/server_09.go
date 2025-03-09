// クライアントにレスポンスを送信する処理。json出力の書き込み
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func write(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Yoichi Kaneko",
		Threads: []string{"1番目", "2番目", "3番目"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", write)
	http.HandleFunc("/writeheader", writeHeader)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", jsonResponse)
    server.ListenAndServe()
}
