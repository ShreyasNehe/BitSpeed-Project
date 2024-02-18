package sql_models

import (
	"gorm.io/gorm"
	"time"
)

type Contact struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	PhoneNumber    string     `json:"phoneNumber"`
	Email          string     `json:"email"`
	LinkedID       uint       `json:"linkedId"`
	LinkPrecedence string     `gorm:"type:enum('primary', 'secondary')" json:"linkPrecedence"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `sql:"index" json:"deletedAt"`
}

func (Contact) TableName() string {
	return "contact"
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Contact{})
}
