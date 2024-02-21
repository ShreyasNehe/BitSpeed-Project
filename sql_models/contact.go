package sql_models

import (
	"gorm.io/gorm"
	"time"
)

type Contact struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	PhoneNumber    string     `gorm:"column:phone_number;default:null" json:"phoneNumber"`
	Email          string     `gorm:"column:email;default:null" json:"email"`
	LinkedID       uint       `gorm:"column:linked_id;default:null" json:"linkedId"`
	LinkPrecedence string     `gorm:"column:link_precedence;type:enum('primary', 'secondary');default:primary" json:"linkPrecedence"`
	CreatedAt      time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt      time.Time  `gorm:"column:updated_at"json:"updatedAt"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" sql:"index" json:"deletedAt"`
}

func (Contact) TableName() string {
	return "contact"
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Contact{})
}
