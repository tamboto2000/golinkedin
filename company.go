package linkedin

type CompanyNode struct {
	Metadata Metadata  `json:"metadata,omitempty"`
	Elements []Company `json:"elements,omitempty"`
	Paging   Paging    `json:"paging,omitempty"`
	Keywords string    `json:"keywords,omitempty"`

	err error
	ln  *Linkedin
}

// Company contain information about a company
type Company struct {
	Industry             map[string]Industry   `json:"industry,omitempty"`
	IndustryUrns         []string              `json:"industryUrns,omitempty"`
	AntiAbuseAnnotations []AntiAbuseAnnotation `json:"$anti_abuse_annotations,omitempty"`
	EntityUrn            string                `json:"entityUrn,omitempty"`
	MiniCompany          *MiniCompany          `json:"miniCompany,omitempty"`
	EmployeeCountRange   *EmployeeCountRange   `json:"employeeCountRange,omitempty"`
	Industries           []string              `json:"industries,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Logo                 *Logo                 `json:"logo,omitempty"`
	RecipeType           string                `json:"$recipeType,omitempty"`
	UniversalName        string                `json:"universalName,omitempty"`
	URL                  string                `json:"url,omitempty"`
	ObjectUrn            string                `json:"objectUrn,omitempty"`
	Showcase             bool                  `json:"showcase,omitempty"`
	Active               bool                  `json:"active,omitempty"`
	TrackingID           string                `json:"trackingId,omitempty"`
	Image                Image                 `json:"image,omitempty"`
	Subtext              Text                  `json:"subtext,omitempty"`
	TargetUrn            string                `json:"targetUrn,omitempty"`
	Text                 Text                  `json:"text,omitempty"`
	DashTargetUrn        string                `json:"dashTargetUrn,omitempty"`
	Type                 string                `json:"type,omitempty"`
}

type Image struct {
	Attributes                  []Attribute   `json:"attributes,omitempty"`
	AccessibilityTextAttributes []interface{} `json:"accessibilityTextAttributes,omitempty"`
}
type Attribute struct {
	MiniCompany *MiniCompany `json:"miniCompany,omitempty"`
	SourceType  string       `json:"sourceType,omitempty"`
}
