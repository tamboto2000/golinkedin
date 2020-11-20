package simple

type Education struct {
	Activities   string     `json:"activities"`
	School       *School    `json:"school"`
	Description  string     `json:"description"`
	DegreeName   string     `json:"degreeName"`
	SchoolName   string     `json:"schoolName"`
	FieldOfStudy string     `json:"fieldOfStudy"`
	DateRange    *DateRange `json:"dateRange"`
}
