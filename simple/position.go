package simple

type Position struct {
	Title        string     `json:"title,omitempty"`
	CompanyName  string     `json:"companyName,omitempty"`
	DateRange    *DateRange `json:"dateRange,omitempty"`
	LocationName string     `json:"locationName,omitempty"`
	Description  string     `json:"description,omitempty"`
	Company      *Company   `json:"company,omitempty"`
}
