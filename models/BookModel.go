package models

type BookModel struct {
	Id     string `schema:"id"`
	Title  string `schema:"title"`
	Author string `schema:"author"`
}
