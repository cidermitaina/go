package main

import (
	"io"
	"net/http"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	// HTML テキストをhttp.ResponseWriteへ
	io.WriteString(w, `
  <!DOCTYPE html>
  <html lang="ja">
  <head>
  <meta charset="UTF-8">
  <title>hello</title>
  </head>
  <body>
  <h1>Hello World</h1>
  </body></html>
  `)
}

func main() {
	// /hello にアクセスした際に処理するハンドラーを登録
	http.HandleFunc("/hello", infoHandler)

	// サーバーをポート8080で起動
	http.ListenAndServe(":8080", nil)
}
