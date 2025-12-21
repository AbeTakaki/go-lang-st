package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ハンドラの宣言
/*
	サーバーで使用するハンドラを定義
	ハンドラ = 「HTTPリクエストを受け取って、そてに対するHTTPレスポンスの内容をコネクションに書き込む」関数の事
	引数： http.ResponseWriter型 と http.Request型
	戻り値： なし

	w http.ResponseWriter に "Hello, Workd!" という文字列を、io.WriteString関数で書き込む流れ
*/
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// ハンドラの処理内容:
	io.WriteString(w, "Hello, World!\n")
}

// Posting Article ハンドラの定義
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article…\n")
}

// article listハンドラの定義
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List…\n")
}

// GET /article/{id} ハンドラの定義
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Lnvalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// article nice
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice…\n")
}

// comment
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment…\n")
}
