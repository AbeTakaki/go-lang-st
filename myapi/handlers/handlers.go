package handlers

import (
	"fmt"
	"io"
	"net/http"
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
	// GET メソッド時のみ通常通りのレスポンスを返す
	if req.Method == http.MethodGet {
		// 通常通りレスポンスを返す
		io.WriteString(w, "Hello, World!\n")
	} else {
		// メソッドがGETではなかった場合は、Invalid method というレスポンスを405で返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// Posting Article ハンドラの定義
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article…\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// article listハンドラの定義
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List…\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// article No.1 ハンドラの定義
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// article nice
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice…\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// comment
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment…\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
