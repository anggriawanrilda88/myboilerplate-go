package models

import "gorm.io/gorm"

type User struct {
	Id        uint           `valid:"-" json:"id" gorm:"primarykey"`
	Name      string         `valid:"alphanum" json:"name"`
	Password  string         `valid:"alphanum" json:"password"`
	Email     string         `valid:"email" json:"email"`
	RoleID    string         `valid:"int" json:"roleId"`
	CreatedAt string         `valid:"-" json:"createdAt"`
	UpdatedAt string         `valid:"-" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `valid:"-" json:"deletedAt" gorm:"index"`
	Role      Role           `valid:"-" json:"role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
