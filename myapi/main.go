package main

import (
	"fmt"
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

	// Posting Article ハンドラの定義
	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article…\n")
	}
	// article listハンドラの定義
	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List…\n")
	}
	// article No.1 ハンドラの定義
	article1Handler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	}
	// article nice
	articleNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice…\n")
	}
	// comment
	commentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment…\n")
	}

	// 定義したハンドラを、サーバーで使用するように登録
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", article1Handler)
	http.HandleFunc("/article/nice", articleNiceHandler)
	http.HandleFunc("/comment", commentHandler)

	// サーバー起動時のログを出力
	log.Println("server start at port 8080")
	// ListenAndServe関数にてサーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
