package simple

type Location struct {
	CountryCode string `json:"countryCode,omitempty"`
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	Description string `json:"description,omitempty"`
	Headquarter bool   `json:"headquarter,omitempty"`
	Line2       string `json:"line2,omitempty"`
	Line1       string `json:"line1,omitempty"`
}
