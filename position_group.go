package linkedin

import "encoding/json"

type PositionGroupNode struct {
	Paging     Paging          `json:"paging,omitempty"`
	RecipeType string          `json:"$recipeType,omitempty"`
	Elements   []PositionGroup `json:"elements,omitempty"`
}

type ProfilePositionInPositionGroupElement struct {
	DateRange                    DateRange       `json:"dateRange,omitempty"`
	MultiLocaleCompanyName       MultiLocale     `json:"multiLocaleCompanyName,omitempty"`
	CompanyName                  string          `json:"companyName,omitempty"`
	Title                        string          `json:"title,omitempty"`
	CompanyUrn                   string          `json:"companyUrn,omitempty"`
	EmploymentTypeUrn            *string         `json:"employmentTypeUrn,omitempty"`
	EntityUrn                    string          `json:"entityUrn,omitempty"`
	MultiLocaleGeoLocationName   MultiLocale     `json:"multiLocaleGeoLocationName,omitempty"`
	Company                      Company         `json:"company,omitempty"`
	ShouldShowSourceOfHireBadge  bool            `json:"shouldShowSourceOfHireBadge,omitempty"`
	LocationName                 *string         `json:"locationName,omitempty"`
	MultiLocaleTitle             MultiLocale     `json:"multiLocaleTitle,omitempty"`
	EmploymentType               *Industry       `json:"employmentType,omitempty"`
	GeoUrn                       string          `json:"geoUrn,omitempty"`
	ProfileTreasuryMediaPosition json.RawMessage `json:"profileTreasuryMediaPosition,omitempty"`
	RegionUrn                    string          `json:"regionUrn,omitempty"`
	GeoLocationName              string          `json:"geoLocationName,omitempty"`
	RecipeType                   string          `json:"$recipeType,omitempty"`
	MultiLocaleLocationName      *MultiLocale    `json:"multiLocaleLocationName,omitempty"`
	Description                  *string         `json:"description,omitempty"`
	MultiLocaleDescription       *MultiLocale    `json:"multiLocaleDescription,omitempty"`
}

type ProfilePositionInPositionGroup struct {
	Paging     Paging                                  `json:"paging,omitempty"`
	RecipeType string                                  `json:"$recipeType,omitempty"`
	Elements   []ProfilePositionInPositionGroupElement `json:"elements,omitempty"`
}

type PositionGroup struct {
	EntityUrn                      string                         `json:"entityUrn,omitempty"`
	DateRange                      DateRange                      `json:"dateRange,omitempty"`
	MultiLocaleCompanyName         MultiLocale                    `json:"multiLocaleCompanyName,omitempty"`
	CompanyName                    string                         `json:"companyName,omitempty"`
	ProfilePositionInPositionGroup ProfilePositionInPositionGroup `json:"profilePositionInPositionGroup,omitempty"`
	Company                        Company                        `json:"company,omitempty"`
	RecipeType                     string                         `json:"$recipeType,omitempty"`
	CompanyUrn                     string                         `json:"companyUrn,omitempty"`
}

type EmployeeCountRange struct {
	Start      int64  `json:"start,omitempty"`
	RecipeType string `json:"$recipeType,omitempty"`
	End        *int64 `json:"end,omitempty"`
}
