// クライアントにレスポンスを送信する処理。WriteHeaderによるレスポンスヘッダの書き込み
package main

import (
	"fmt"
	"net/http"
)

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


func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", write)
	http.HandleFunc("/writeheader", writeHeader)
    server.ListenAndServe()
}
