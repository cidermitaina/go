package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// /now にアクセスした際に処理するハンドラーを登録
	http.HandleFunc("/now", handleClock)

	// サーバーをポート8080で起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleClock(w http.ResponseWriter, r *http.Request) {
	// 現在時刻をHTMLで出力
	fmt.Fprintf(w, `
    <!DOCTYPE html>
    <html><body>
        %s
    </body></html>
    `, time.Now().Format("2006-01-02 15:04:05"))
}
