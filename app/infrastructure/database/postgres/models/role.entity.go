package models

import "gorm.io/gorm"

// Role model
type Role struct {
	Id          uint           `valid:"-" json:"id" gorm:"primarykey"`
	Name        string         `valid:"alphanum" gorm:"unique;not null" json:"name" xml:"name" form:"name" query:"name"`
	Description string         `valid:"alphanum" gorm:"type:varchar(100);" json:"description"`
	CreatedAt   string         `valid:"-" json:"created_at"`
	UpdatedAt   string         `valid:"-" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `valid:"-" json:"deleted_at" gorm:"index"`
}
