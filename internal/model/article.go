package model

type Article struct {
	Base
	Title   string `json:"title"`
	Content string `json:"content"`
}
