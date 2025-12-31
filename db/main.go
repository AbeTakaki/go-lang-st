package main

import (
	"database/sql"
	"fmt"

	"github.com/AbeTakaki/go-lang-st/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "dockertest"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// DB の接続
	// if err := db.Ping(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("connect to DB")
	// }

	// select 文の実装
	const sqlStr = `
		select *
		from articles;
	`

	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime

		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}
	fmt.Printf("%+v\n", articleArray)
}
