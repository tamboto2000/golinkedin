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

type VectorImage struct {
	RecipeType string     `json:"$recipeType,omitempty"`
	RootURL    string     `json:"rootUrl,omitempty"`
	Artifacts  []Artifact `json:"artifacts,omitempty"`
}

type Artifact struct {
	Width                         int    `json:"width,omitempty"`
	FileIdentifyingURLPathSegment string `json:"fileIdentifyingUrlPathSegment,omitempty"`
	RecipeType                    string `json:"$recipeType,omitempty"`
	ExpiresAt                     int64  `json:"expiresAt,omitempty"`
	Height                        int    `json:"height,omitempty"`
}

type PhotoFilterEditInfo struct {
	BottomLeft      *BottomLeft `json:"bottomLeft,omitempty"`
	Vignette        float64     `json:"vignette,omitempty"`
	BottomRight     *BottomLeft `json:"bottomRight,omitempty"`
	TopRight        *BottomLeft `json:"topRight,omitempty"`
	Saturation      float64     `json:"saturation,omitempty"`
	Brightness      float64     `json:"brightness,omitempty"`
	PhotoFilterType string      `json:"photoFilterType,omitempty"`
	Contrast        float64     `json:"contrast,omitempty"`
	TopLeft         *BottomLeft `json:"topLeft,omitempty"`
	RecipeType      string      `json:"$recipeType,omitempty"`
}

type BottomLeft struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type Logo struct {
	COMLinkedinCommonVectorImage *VectorImage `json:"com.linkedin.common.VectorImage,omitempty"`
	VectorImage                  *VectorImage `json:"vectorImage,omitempty"`
}

type TimePeriod struct {
	EndDate   *Date `json:"endDate,omitempty"`
	StartDate *Date `json:"startDate,omitempty"`
}

type Date struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

type DateRange struct {
	Start *Date `json:"start,omitempty"`
	End   *Date `json:"end,omitempty"`
}

type AntiAbuseAnnotation struct {
	AttributeID int64 `json:"attributeId,omitempty"`
	EntityID    int64 `json:"entityId,omitempty"`
}

type Locale struct {
	Country              string                `json:"country,omitempty"`
	Language             string                `json:"language,omitempty"`
	RecipeType           string                `json:"$recipeType,omitempty"`
	AntiAbuseAnnotations []AntiAbuseAnnotation `json:"$anti_abuse_annotations,omitempty"`
}

type Company struct {
	Industry             map[string]Industry   `json:"industry,omitempty"`
	IndustryUrns         []string              `json:"industryUrns,omitempty"`
	AntiAbuseAnnotations []AntiAbuseAnnotation `json:"$anti_abuse_annotations,omitempty"`
	EntityUrn            string                `json:"entityUrn,omitempty"`
	Name                 string                `json:"name,omitempty"`
	Logo                 *Logo                 `json:"logo,omitempty"`
	RecipeType           string                `json:"$recipeType,omitempty"`
	UniversalName        string                `json:"universalName,omitempty"`
	URL                  string                `json:"url,omitempty"`
	EmployeeCountRange   *EmployeeCountRange   `json:"employeeCountRange,omitempty"`
}

type Industry struct {
	Name                string `json:"name,omitempty"`
	RecipeType          string `json:"$recipeType,omitempty"`
	EntityUrn           string `json:"entityUrn,omitempty"`
	CompanyNameRequired *bool  `json:"companyNameRequired,omitempty"`
}

type GeoLocation struct {
	Geo        Geo    `json:"geo,omitempty"`
	GeoUrn     string `json:"geoUrn,omitempty"`
	RecipeType string `json:"$recipeType,omitempty"`
}

type Geo struct {
	CountryUrn                             string   `json:"countryUrn,omitempty"`
	Country                                *Country `json:"country,omitempty"`
	DefaultLocalizedNameWithoutCountryName string   `json:"defaultLocalizedNameWithoutCountryName,omitempty"`
	EntityUrn                              string   `json:"entityUrn,omitempty"`
	RecipeType                             string   `json:"$recipeType,omitempty"`
	DefaultLocalizedName                   string   `json:"defaultLocalizedName,omitempty"`
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
