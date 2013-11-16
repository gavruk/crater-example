package models

type AddBookViewModel struct {
	User User
	Book *BookModel
}

func NewAddBookViewModel(userFromSession interface{}) *AddBookViewModel {
	username := userFromSession.(string)
	model := new(AddBookViewModel)
	model.User = User{Name: username}
	model.Book = new(BookModel)
	return model
}
