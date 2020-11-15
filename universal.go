package linkedin

// Paging control resource cursoring
type Paging struct {
	// set entities per page
	Count int `json:"count,omitempty"`
	// set start cursor
	Start int `json:"start,omitempty"`
	// total entities in this node
	Total int `json:"total,omitempty"`
	// Linkedin cursor link?
	Links      []interface{} `json:"links,omitempty"`
	RecipeType string        `json:"$recipeType,omitempty"`
}

type MultiLocale struct {
	EnUS string `json:"en_US,omitempty"`
}

type DisplayImageReference struct {
	VectorImage *VectorImage `json:"vectorImage,omitempty"`
	URL         string       `json:"url,omitempty"`
}

// VectorImage store images of multiple qualities
type VectorImage struct {
	RecipeType string `json:"$recipeType,omitempty"`
	// use RootURL + Artifact.FileIdentifyingURLPathSegment for creating url to an image
	RootURL   string     `json:"rootUrl,omitempty"`
	Artifacts []Artifact `json:"artifacts,omitempty"`
}

// Artifact store the suffix path of an image
type Artifact struct {
	Width                         int    `json:"width,omitempty"`
	FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment,omitempty"`
	RecipeType                    string `json:"$recipeType,omitempty"`
	ExpiresAt                     int64  `json:"expiresAt,omitempty"`
	Height                        int    `json:"height,omitempty"`
}

type PhotoFilterEditInfo struct {
	BottomLeft      *Coordinate `json:"bottomLeft,omitempty"`
	Vignette        float64     `json:"vignette,omitempty"`
	BottomRight     *Coordinate `json:"bottomRight,omitempty"`
	TopRight        *Coordinate `json:"topRight,omitempty"`
	Saturation      float64     `json:"saturation,omitempty"`
	Brightness      float64     `json:"brightness,omitempty"`
	PhotoFilterType string      `json:"photoFilterType,omitempty"`
	Contrast        float64     `json:"contrast,omitempty"`
	TopLeft         *Coordinate `json:"topLeft,omitempty"`
	RecipeType      string      `json:"$recipeType,omitempty"`
}

type Coordinate struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

// Logo ususally used in Company or School for their logo image
type Logo struct {
	// Sometimes the images stored in here
	COMLinkedinCommonVectorImage *VectorImage `json:"com.linkedin.common.VectorImage,omitempty"`
	// Sometimes in here, make sure to check both of these fields
	VectorImage *VectorImage `json:"vectorImage,omitempty"`
}

// TimePeriod represent a period of time
type TimePeriod struct {
	EndDate   *Date `json:"endDate,omitempty"`
	StartDate *Date `json:"startDate,omitempty"`
}

type Date struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

// DateRange is kinda like TimePeriod, don't know why linkedin have to use different approach for different object
type DateRange struct {
	Start *Date `json:"start,omitempty"`
	End   *Date `json:"end,omitempty"`
}

type AntiAbuseAnnotation struct {
	AttributeID int `json:"attributeId,omitempty"`
	EntityID    int `json:"entityId,omitempty"`
}

type Locale struct {
	Country              string                `json:"country,omitempty"`
	Language             string                `json:"language,omitempty"`
	RecipeType           string                `json:"$recipeType,omitempty"`
	AntiAbuseAnnotations []AntiAbuseAnnotation `json:"$anti_abuse_annotations,omitempty"`
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
}

type EmployeeCountRange struct {
	Start      int    `json:"start,omitempty"`
	RecipeType string `json:"$recipeType,omitempty"`
	End        int    `json:"end,omitempty"`
}

type Industry struct {
	Name                string `json:"name,omitempty"`
	RecipeType          string `json:"$recipeType,omitempty"`
	EntityUrn           string `json:"entityUrn,omitempty"`
	CompanyNameRequired *bool  `json:"companyNameRequired,omitempty"`
}

type Country struct {
	RecipeType           string `json:"$recipeType,omitempty"`
	EntityUrn            string `json:"entityUrn,omitempty"`
	DefaultLocalizedName string `json:"defaultLocalizedName,omitempty"`
}

type Location struct {
	PreferredGeoPlace string `json:"preferredGeoPlace,omitempty"`
	CountryCode       string `json:"countryCode,omitempty"`
}

type MiniCompany struct {
	ObjectUrn     string `json:"objectUrn,omitempty"`
	EntityUrn     string `json:"entityUrn,omitempty"`
	Name          string `json:"name,omitempty"`
	Showcase      bool   `json:"showcase,omitempty"`
	Active        bool   `json:"active,omitempty"`
	Logo          *Logo  `json:"logo,omitempty"`
	UniversalName string `json:"universalName,omitempty"`
	TrackingID    string `json:"trackingId,omitempty"`
}
