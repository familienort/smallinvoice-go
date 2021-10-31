package smallinvoice

// Address is the main contact address.
type Address struct {
	Country  string `json:"country"`
	Street   string `json:"street"`
	Street2  string `json:"street2"`
	Postcode string `json:"postcode"`
	City     string `json:"city"`
}

// Contact is the contact entity.
type Contact struct {
	Gender      string  `json:"gender"`
	Email       string  `json:"email"`
	Name        string  `json:"name"`
	MainAddress Address `json:"main_address"`
}
