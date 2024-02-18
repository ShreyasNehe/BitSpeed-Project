package model

type ContactFilter struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type ContactResponse struct {
	Contact struct {
		PrimaryContactID    uint     `json:"primaryContactId"`
		Emails              []string `json:"emails"`
		PhoneNumbers        []string `json:"phoneNumbers"`
		SecondaryContactIds []uint   `json:"secondaryContactIds"`
	} `json:"contact"`
}
