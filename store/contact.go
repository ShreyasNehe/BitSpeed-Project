package store

import (
	"fmt"
	"github.com/suresh024/identity_reconciliation/consts"
	"github.com/suresh024/identity_reconciliation/db"
	"github.com/suresh024/identity_reconciliation/model"
	"github.com/suresh024/identity_reconciliation/sql_models"
	"gorm.io/gorm"
)

type ContactStore interface {
	FetchContacts(payload model.ContactFilter) ([]sql_models.Contact, error)
	CreateContact(payload model.ContactFilter, precedence string, linkedId uint) (sql_models.Contact, error)
	TogglePrimaryContact(payload model.ContactFilter, primaryID uint) ([]sql_models.Contact, error)
	FetchAllByLinkedID(linkedID uint) (sql_models.Contact, []sql_models.Contact, error)
}

type contactRepo struct {
	db *gorm.DB
}

func NewContactRepo() ContactStore {
	return &contactRepo{
		db: db.DBClient,
	}
}

func (s *contactRepo) FetchContacts(payload model.ContactFilter) ([]sql_models.Contact, error) {
	funcDesc := "FetchPrimaryContact | Repo"

	var primaryContact []sql_models.Contact
	query := s.db.Debug().Where("email = ? OR phone_number = ?",
		payload.Email, payload.PhoneNumber).Order("created_at ASC")
	result := query.Find(&primaryContact)
	if result.Error != nil && result.Error.Error() != "record not found" {
		fmt.Printf("%s | errMsg: %s", funcDesc, result.Error)
		return primaryContact, fmt.Errorf("error while searching contacts")
	}

	return primaryContact, nil

}
func (s *contactRepo) FetchAllByLinkedID(linkedID uint) (sql_models.Contact, []sql_models.Contact, error) {
	funcDesc := "FetchAllByLinkedID | Repo"

	var linkedContacts []sql_models.Contact
	var primaryContact sql_models.Contact

	res := s.db.Where("linkedId = ?", linkedID).
		Find(&linkedContacts).Order("created_at ASC")
	if res.Error != nil {
		fmt.Printf("%s | errMsg: %s", funcDesc, res.Error)
		return primaryContact, linkedContacts, fmt.Errorf("error while fetching linked contact")
	}
	res = s.db.Where("id = ?", linkedID).
		First(&primaryContact).Order("created_at ASC")
	if res.Error != nil {
		fmt.Printf("%s | errMsg: %s", funcDesc, res.Error)
		return primaryContact, linkedContacts, fmt.Errorf("error while fetching linked contact")
	}
	return primaryContact, linkedContacts, nil
}
func (s *contactRepo) CreateContact(payload model.ContactFilter, precedence string, linkedId uint) (sql_models.Contact, error) {
	funcDesc := "CreateContact | Repo"

	var response sql_models.Contact
	var contactData = sql_models.Contact{
		LinkPrecedence: consts.PrimaryPrecedence,
	}
	if precedence == consts.SecondaryPrecedence {
		contactData.LinkPrecedence = consts.SecondaryPrecedence
		contactData.LinkedID = linkedId
	}
	if payload.PhoneNumber != nil {
		contactData.PhoneNumber = *payload.PhoneNumber
	}

	if payload.Email != nil {
		contactData.Email = *payload.Email
	}

	res := s.db.Debug().Create(&contactData)
	if res.Error != nil {
		fmt.Printf("%s | errMsg: %s", funcDesc, res.Error)
		return contactData, fmt.Errorf("error while insert contact")
	}

	res = s.db.Debug().Where("email=? AND phone_number=?",
		contactData.Email, contactData.PhoneNumber).
		Order("created_at ASC").
		First(&response)
	if res.Error != nil {
		fmt.Printf("%s | errMsg: %s", funcDesc, res.Error)
		return response, fmt.Errorf("error while insert contact")
	}

	return response, nil

}

func (s *contactRepo) TogglePrimaryContact(payload model.ContactFilter, primaryID uint) ([]sql_models.Contact, error) {
	funcDesc := "TogglePrimaryContact | Repo"

	contacts := make([]sql_models.Contact, 0)
	subquery := s.db.Debug().Model(&sql_models.Contact{}).
		Select("id"). // Select only the IDs
		Where("(email = ? OR phone_number = ?) AND link_precedence = ?", payload.Email, payload.PhoneNumber, "primary").
		Order("created_at ASC").
		Limit(-1). // Set Limit to -1 to remove any limit on the number of records
		Offset(1)  // Apply the desired offset

	res := s.db.Debug().Model(&sql_models.Contact{}).
		Where("id IN (?)", subquery).
		Updates(map[string]interface{}{
			"linked_id":       primaryID,
			"link_precedence": "secondary",
			"updated_at":      gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if res.Error != nil && res.Error.Error() != "record not found" {
		fmt.Printf("%s | errMsg: %s", funcDesc, res.Error)
		return contacts, fmt.Errorf("error while switching secondary contact")
	}

	if res.RowsAffected > 0 {
		res = s.db.Debug().Where("id = ? OR linked_id = ?",
			primaryID, primaryID).Order("created_at ASC").
			Find(&contacts)
		if res.Error != nil {
			fmt.Printf("%s | errMsg: %s", funcDesc, res.Error)
			return contacts, fmt.Errorf("error while fetched switched contacts")
		}
		return contacts, nil
	}
	return contacts, nil
}
