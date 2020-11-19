package simple

// Profile is simplified Profile object
type Profile struct {
	FirstName       string `json:"firstName,omitempty"`
	LastName        string `json:"lastName,omitempty"`
	Headline        string `json:"headline,omitempty"`
	GeoLocationName string `json:"geoLocationName,omitempty"`
	About           string `json:"aboout,omitempty"`
}

type Experience struct {
}
