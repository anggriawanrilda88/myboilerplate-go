package models

import "gorm.io/gorm"

// Role model
type Role struct {
	Id          uint           `valid:"-" json:"id" gorm:"primarykey"`
	Name        string         `valid:"alphanum" gorm:"unique;not null" json:"name" xml:"name" form:"name" query:"name"`
	Description string         `valid:"alphanum" gorm:"type:varchar(100);" json:"description"`
	CreatedAt   string         `valid:"-" json:"createdAt"`
	UpdatedAt   string         `valid:"-" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `valid:"-" json:"deletedAt" gorm:"index"`
}
