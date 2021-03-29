package models

import "gorm.io/gorm"

type User struct {
	Id        uint           `json:"id" gorm:"primarykey"`
	Name      string         `validate:"alphanum" json:"name"`
	Password  string         `validate:"alphanum" json:"password"`
	Email     string         `validate:"email" json:"email"`
	RoleID    uint           `validate:"required" json:"roleId"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Role      Role           `json:"role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type UserLogin struct {
	Password string `validate:"alphanum" json:"password"`
	Email    string `validate:"email" json:"email"`
	Token    string `json:"token"`
}
