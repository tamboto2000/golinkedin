package simple

type Position struct {
	Title        string     `json:"title"`
	CompanyName  string     `json:"companyName"`
	DateRange    *DateRange `json:"dateRange"`
	LocationName string     `json:"locationName"`
	Description  string     `json:"description"`
	Company      *Company   `json:"company"`
}
