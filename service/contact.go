package service

import (
	"github.com/suresh024/identity_reconciliation/consts"
	"github.com/suresh024/identity_reconciliation/model"
	"github.com/suresh024/identity_reconciliation/sql_models"
	"github.com/suresh024/identity_reconciliation/store"
	"log"
)

type ContactService interface {
	FetchContacts(payload model.ContactFilter) (model.ContactResponse, error)
}
type contactService struct {
	contactStore store.ContactStore
}

func NewContactService(repo store.Store) ContactService {
	return &contactService{
		contactStore: repo.ContactStore,
	}
}

func (s *contactService) FetchContacts(payload model.ContactFilter) (model.ContactResponse, error) {
	funcName := "FetchContacts | Service"

	var response model.ContactResponse
	contacts, err := s.contactStore.FetchContacts(payload)
	if err != nil {
		log.Printf("%s | errMsg: %s", funcName, err.Error())
		return response, err
	}
	switch len(contacts) {
	case 0:
		contact, err := s.contactStore.CreateContact(payload, consts.PrimaryPrecedence, 0)
		if err != nil {
			log.Printf("%s | errMsg: %s", funcName, err.Error())
			return response, err
		}
		return PrepareResponseStructure([]sql_models.Contact{contact}), nil
	case 1:
		if (payload.Email != nil && contacts[0].Email == *payload.Email) &&
			(payload.PhoneNumber != nil && contacts[0].PhoneNumber == *payload.PhoneNumber) {
			return PrepareResponseStructure(contacts), nil
		}

		var needToCreate bool
		if (payload.Email != nil && contacts[0].Email != *payload.Email) || (payload.PhoneNumber != nil && contacts[0].PhoneNumber != *payload.PhoneNumber) {
			needToCreate = true
		}

		if needToCreate {
			newContactResponse, err := s.contactStore.CreateContact(payload, consts.SecondaryPrecedence, contacts[0].ID)
			if err != nil {
				log.Printf("%s | errMsg: %s", funcName, err.Error())
				return response, err
			}
			contacts = append(contacts, newContactResponse)
		}
		if contacts[0].LinkPrecedence == consts.SecondaryPrecedence {

		}
		return PrepareResponseStructure(contacts), nil
	default:
		primaryContactID := contacts[0].ID
		count := 0
		for _, val := range contacts {
			if val.LinkPrecedence == consts.PrimaryPrecedence {
				if count == 0 {
					primaryContactID = val.ID
				}
				count += 1
			}
		}
		if count > 1 {
			updatedContacts, err := s.contactStore.TogglePrimaryContact(payload, primaryContactID)
			if err != nil {
				log.Printf("%s | errMsg: %s", funcName, err.Error())
				return response, err
			}
			if len(updatedContacts) > 0 {
				contacts = updatedContacts
			}
		}
		var emailFlag, phoneFalg bool
		for _, contact := range contacts {
			if payload.Email != nil && contact.Email == *payload.Email {
				emailFlag = true
			}
			if payload.PhoneNumber != nil && contact.PhoneNumber == *payload.PhoneNumber {
				phoneFalg = true
			}
		}

		if emailFlag && phoneFalg {
			return PrepareResponseStructure(contacts), nil
		}

		if !emailFlag || !phoneFalg {
			newSecondaryContact, err := s.contactStore.CreateContact(payload, consts.SecondaryPrecedence, primaryContactID)
			if err != nil {
				log.Printf("%s | errMsg: %s", funcName, err.Error())
				return response, err
			}
			contacts = append(contacts, newSecondaryContact)
		}
		return PrepareResponseStructure(contacts), nil
	}
}
