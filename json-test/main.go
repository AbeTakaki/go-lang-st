package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}

	comment2 := Comment{
		CommentID: 2,
		ArticleID: 2,
		Message:   "test comment2",
		CreatedAt: time.Now(),
	}

	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article.",
		UserName:    "Mike",
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	output, err := json.MarshalIndent(article, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}
