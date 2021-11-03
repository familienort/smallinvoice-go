package smallinvoice

// ContactType is either a commercial or private customer type.
type ContactType string

// ContactRelation is either a client or creditor.
type ContactRelation string

// ContactGender is either F or M.
type ContactGender string

const (
	// ContactTypePrivate is a private customer.
	ContactTypePrivate ContactType = "P"

	// ContactTypeCompany is a commercial customer/company.
	ContactTypeCompany ContactType = "C"

	// ContactRelationClient is a client.
	ContactRelationClient ContactRelation = "CL"

	// ContactRelationCreditor is a creditor.
	ContactRelationCreditor ContactRelation = "CR"

	// ContactGenderFemale ...
	ContactGenderFemale ContactGender = "F"

	// ContactGenderMale ...
	ContactGenderMale ContactGender = "M"
)

// Address is the main contact address.
type Address struct {
	Country  string `json:"country,omitempty"`
	Street   string `json:"street,omitempty"`
	Street2  string `json:"street2,omitempty"`
	StreetNo string `json:"street_no,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	City     string `json:"city,omitempty"`
}

// Contact is the contact entity.
type Contact struct {
	Gender      ContactGender     `json:"gender"`
	Type        ContactType       `json:"type"`
	Relation    []ContactRelation `json:"relation"`
	Email       string            `json:"email"`
	Name        string            `json:"name"`
	MainAddress Address           `json:"main_address"`
}
