package store

import (
	"github.com/suresh024/identity_reconciliation/db"
	"gorm.io/gorm"
)

type ContactStore interface {
	FetchPrimaryContacts()
}

type contactRepo struct {
	db *gorm.DB
}

func NewContactRepo() ContactStore {
	return &contactRepo{
		db: db.DBClient,
	}
}

func (s *contactRepo) FetchPrimaryContacts() {

}
