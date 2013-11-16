package models

type BookListViewModel struct {
	User  User
	Books []*BookModel
}

func NewBookListViewModel(userFromSession interface{}) *BookListViewModel {
	username := userFromSession.(string)
	model := new(BookListViewModel)
	model.User = User{Name: username}
	model.Books = make([]*BookModel, 0)
	return model
}
