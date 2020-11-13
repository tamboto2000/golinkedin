package linkedin

type CertificationNode struct {
	Paging     Paging          `json:"paging,omitempty"`
	RecipeType string          `json:"$recipeType,omitempty"`
	Elements   []Certification `json:"elements,omitempty"`
}

type Certification struct {
	DateRange                DateRange   `json:"dateRange,omitempty"`
	MultiLocaleLicenseNumber MultiLocale `json:"multiLocaleLicenseNumber,omitempty"`
	CompanyUrn               string      `json:"companyUrn,omitempty"`
	URL                      string      `json:"url,omitempty"`
	MultiLocaleAuthority     MultiLocale `json:"multiLocaleAuthority,omitempty"`
	EntityUrn                string      `json:"entityUrn,omitempty"`
	Authority                string      `json:"authority,omitempty"`
	Name                     string      `json:"name,omitempty"`
	MultiLocaleName          MultiLocale `json:"multiLocaleName,omitempty"`
	LicenseNumber            string      `json:"licenseNumber,omitempty"`
	Company                  Company     `json:"company,omitempty"`
	RecipeType               string      `json:"$recipeType,omitempty"`
	DisplaySource            string      `json:"displaySource,omitempty"`
}
