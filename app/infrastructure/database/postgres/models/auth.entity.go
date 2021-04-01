package models

type UserLogin struct {
	Password string `validate:"alphanum" json:"password"`
	Email    string `validate:"email" json:"email"`
	Token    string `json:"token"`
}
