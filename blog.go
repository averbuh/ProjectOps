package main

import "fmt"

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Blog struct {
	Articles []Article
}

func New() *Blog {
	fmt.Println("New blog created")
	return &Blog{}
}

func (b *Blog) SaveArticle(article Article) {
	b.Articles = append(b.Articles, article)
}

func (b *Blog) FetchAll() []Article {
	return b.Articles
}
