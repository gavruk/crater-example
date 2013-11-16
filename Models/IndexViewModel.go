package models

type IndexViewModel struct {
	User User
}

func NewIndexViewModel(userFromSession interface{}) *IndexViewModel {
	username := userFromSession.(string)
	model := new(IndexViewModel)
	model.User = User{Name: username}
	return model
}
