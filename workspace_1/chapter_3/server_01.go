// 最も単純なwebサーバ
package main

import (
	"net/http"
)

func main() {
    http.ListenAndServe("", nil)
}
