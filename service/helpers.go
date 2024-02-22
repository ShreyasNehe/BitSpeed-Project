package service

import (
	"github.com/suresh024/identity_reconciliation/model"
	"github.com/suresh024/identity_reconciliation/sql_models"
)

func PrepareResponseStructure(contacts []sql_models.Contact) model.ContactResponse {
	var output model.ContactResponse

	var primaryEmail, primaryPhoneNumber string
	var secondaryEmails, secondaryPhoneNumber []string
	emailMap := map[string]struct{}{}
	phoneMap := map[string]struct{}{}

	for _, contact := range contacts {
		if contact.LinkPrecedence == "primary" {
			output.Contact.PrimaryContactID = contact.ID
			primaryEmail = contact.Email
			primaryPhoneNumber = contact.PhoneNumber
		} else if contact.LinkPrecedence == "secondary" {
			if _, ok := emailMap[contact.Email]; !ok && contact.Email != primaryEmail && contact.Email != "" {
				secondaryEmails = append(secondaryEmails, contact.Email)
				emailMap[contact.Email] = struct{}{}
			}
			if _, ok := phoneMap[contact.PhoneNumber]; !ok && contact.PhoneNumber != primaryPhoneNumber && contact.PhoneNumber != "" {
				secondaryPhoneNumber = append(secondaryPhoneNumber, contact.PhoneNumber)
				phoneMap[contact.PhoneNumber] = struct{}{}
			}
			output.Contact.SecondaryContactIds = append(output.Contact.SecondaryContactIds, contact.ID)
		}
	}

	output.Contact.Emails = append([]string{primaryEmail}, secondaryEmails...)
	output.Contact.PhoneNumbers = append([]string{primaryPhoneNumber}, secondaryPhoneNumber...)

	return output
}
