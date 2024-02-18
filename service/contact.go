package service

import "github.com/suresh024/identity_reconciliation/store"

type ContactService interface {
	FetchContacts()
}
type contactService struct {
	contactStore store.ContactStore
}

func NewContactService(repo store.Store) ContactService {
	return &contactService{
		contactStore: repo.ContactStore,
	}
}

func (s *contactService) FetchContacts() {

}
