package golinkedin

// Paging control resource cursoring.
// It is highly advised to NOT CHANGE ANY OF THE VALUES OF Paging as it can cause infinite loop when cursoring
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
	Image       *Image       `json:"image,omitempty"`
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

type EmployeeCountRange struct {
	Start      int    `json:"start,omitempty"`
	RecipeType string `json:"$recipeType,omitempty"`
	End        int    `json:"end,omitempty"`
}

type Country struct {
	RecipeType           string `json:"$recipeType,omitempty"`
	EntityUrn            string `json:"entityUrn,omitempty"`
	DefaultLocalizedName string `json:"defaultLocalizedName,omitempty"`
}

type Location struct {
	PreferredGeoPlace string `json:"preferredGeoPlace,omitempty"`
	CountryCode       string `json:"countryCode,omitempty"`
	Country           string `json:"country,omitempty"`
	City              string `json:"city,omitempty"`
	PostalCode        string `json:"postalCode,omitempty"`
	Description       string `json:"description,omitempty"`
	Headquarter       bool   `json:"headquarter,omitempty"`
	Line2             string `json:"line2,omitempty"`
	Line1             string `json:"line1,omitempty"`
}

type Picture struct {
	COMLinkedinCommonVectorImage *VectorImage `json:"com.linkedin.common.VectorImage,omitempty"`
}

type Text struct {
	TextDirection string        `json:"textDirection,omitempty"`
	Text          string        `json:"text,omitempty"`
	Attributes    []interface{} `json:"attributes,omitempty"`
}

type Attribute struct {
	MiniCompany *MiniCompany `json:"miniCompany,omitempty"`
	MiniProfile *MiniProfile `json:"miniProfile,omitempty"`
	MiniSchool  *MiniSchool  `json:"miniSchool,omitempty"`
	MiniGroup   *MiniGroup   `json:"miniGroup,omitempty"`
	Distance    *Distance    `json:"distance,omitempty"`
	SourceType  string       `json:"sourceType,omitempty"`
	Start       int64        `json:"start,omitempty"`
	Length      int64        `json:"length,omitempty"`
	// can be string or Type
	Type        interface{} `json:"type,omitempty"`
	ArtDecoIcon string      `json:"artDecoIcon,omitempty"`
	VectorImage VectorImage `json:"vectorImage,omitempty"`
}

type Distance struct {
	Value string `json:"value"`
}

type MiniGroup struct {
	GroupName        string `json:"groupName,omitempty"`
	ObjectUrn        string `json:"objectUrn,omitempty"`
	GroupDescription string `json:"groupDescription,omitempty"`
	EntityUrn        string `json:"entityUrn,omitempty"`
	Logo             *Logo  `json:"logo,omitempty"`
	TrackingID       string `json:"trackingId,omitempty"`
}

type Image struct {
	COMLinkedinCommonVectorImage *VectorImage  `json:"com.linkedin.common.VectorImage,omitempty"`
	Attributes                   []Attribute   `json:"attributes,omitempty"`
	AccessibilityTextAttributes  []interface{} `json:"accessibilityTextAttributes,omitempty"`
}

type MiniSchool struct {
	Logo       Logo   `json:"logo,omitempty"`
	SchoolName string `json:"schoolName,omitempty"`
	ObjectUrn  string `json:"objectUrn,omitempty"`
	EntityUrn  string `json:"entityUrn,omitempty"`
	TrackingID string `json:"trackingId,omitempty"`
}

type Metadata struct {
	PaginationToken      string       `json:"paginationToken,omitempty"`
	NewRelevanceFeed     bool         `json:"newRelevanceFeed,omitempty"`
	ID                   string       `json:"id,omitempty"`
	Type                 string       `json:"type,omitempty"`
	QueryAfterTime       int64        `json:"queryAfterTime,omitempty"`
	Urn                  string       `json:"urn,omitempty"`
	ActionsPosition      string       `json:"actionsPosition,omitempty"`
	ActionTriggerEnabled bool         `json:"actionTriggerEnabled,omitempty"`
	DetailPageType       string       `json:"detailPageType,omitempty"`
	ShareAudience        string       `json:"shareAudience,omitempty"`
	ShareUrn             string       `json:"shareUrn,omitempty"`
	ExcludedFromSeen     bool         `json:"excludedFromSeen,omitempty"`
	ActionsUrn           string       `json:"actionsUrn,omitempty"`
	Actions              []Action     `json:"actions,omitempty"`
	TrackingData         TrackingData `json:"trackingData,omitempty"`
}

type Action struct {
	ActionType string      `json:"actionType,omitempty"`
	Subtext    string      `json:"subtext,omitempty"`
	SaveAction *SaveAction `json:"saveAction,omitempty"`
	// can be string or Text
	Text             interface{} `json:"text,omitempty"`
	URL              string      `json:"url,omitempty"`
	TargetUrn        string      `json:"targetUrn,omitempty"`
	ContentSource    string      `json:"contentSource,omitempty"`
	AuthorProfileID  string      `json:"authorProfileId,omitempty"`
	AuthorUrn        string      `json:"authorUrn,omitempty"`
	ConfirmationText *Text       `json:"confirmationText,omitempty"`
}

type SaveAction struct {
	EntityUrn string `json:"entityUrn,omitempty"`
	Saved     bool   `json:"saved,omitempty"`
}

type CountRange struct {
	Start  int64 `json:"start,omitempty"`
	Length int   `json:"length,omitempty"`
	End    int   `json:"end,omitempty"`
}
