package main

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/ToruMakabe/wdp-container-handson-part2/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// クエリ文字列"name"を取得する
	name := r.URL.Query().Get("name")
	// Hello関数を呼び出し、結果をレスポンスライターに書き込む
	io.WriteString(w, hello.Hello(name))
}

func main() {
	// リッスンするポート番号を指定する
	portNumber := "9000"
	addr := ":" + portNumber
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening on port ", portNumber)
	// /へのアクセスでhandle関数を呼び出す
	http.HandleFunc("/", handle)
	srv := http.Server{}
	srv.Serve(l)
}
