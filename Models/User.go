package models

// User model
type User struct {
	Name string `schema:"name"`
	Age  int    `schema:"age"`
}
