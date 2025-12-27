package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/AbeTakaki/go-lang-st/models"
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
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// GET /article/listハンドラの定義
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// 変数 page
	var page int
	// パラメータ page が1個以上ある場合
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		// パラメータ page に対応する１つ目の値を採用し、数値に変換する
		var err error
		page, err = strconv.Atoi(p[0])

		// 数値に変換できない値だった場合は 400 エラーを返す
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		// パラメータ page が存在しなかった場合
	} else {
		// パラメータ page=1 と同じ処理を行う
		page = 1
	}

	log.Println(page)

	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id} ハンドラの定義
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Lnvalid query parameter", http.StatusBadRequest)
		return
	}

	log.Panicln(articleID)
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// article nice
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// comment
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
