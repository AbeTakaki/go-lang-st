package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// ハンドラの宣言
	/*
		サーバーで使用するハンドラを定義
		ハンドラ = 「HTTPリクエストを受け取って、そてに対するHTTPレスポンスの内容をコネクションに書き込む」関数の事
		引数： http.ResponseWriter型 と http.Request型
		戻り値： なし

		w http.ResponseWriter に "Hello, Workd!" という文字列を、io.WriteString関数で書き込む流れ
	*/
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// ハンドラの処理内容:
		// 何が来ても、Hello,World! の文字列を返す
		io.WriteString(w, "Hello, World!\n")
	}

	// 定義したハンドラを、サーバーで使用するように登録
	http.HandleFunc("/", helloHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にてサーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
