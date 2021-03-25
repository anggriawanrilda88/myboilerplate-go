package models

// User model
type User struct {
	Name     string `valid:"alphanum" json:"name"`
	Password string `valid:"alphanum" json:"password"`
	Email    string `valid:"email"  json:"email"`
	RoleID   string `valid:"int"  json:"role_id"`
}
