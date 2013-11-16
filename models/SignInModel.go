package models

type SignInModel struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
}
