package linkedin

import "encoding/json"

type PositionGroupNode struct {
	Paging     Paging          `json:"paging"`
	RecipeType string          `json:"$recipeType"`
	Elements   []PositionGroup `json:"elements"`
}

type ProfilePositionInPositionGroupElement struct {
	DateRange                    DateRange       `json:"dateRange"`
	MultiLocaleCompanyName       MultiLocale     `json:"multiLocaleCompanyName"`
	CompanyName                  string          `json:"companyName"`
	Title                        string          `json:"title"`
	CompanyUrn                   string          `json:"companyUrn"`
	EmploymentTypeUrn            *string         `json:"employmentTypeUrn,omitempty"`
	EntityUrn                    string          `json:"entityUrn"`
	MultiLocaleGeoLocationName   MultiLocale     `json:"multiLocaleGeoLocationName"`
	Company                      Company         `json:"company"`
	ShouldShowSourceOfHireBadge  bool            `json:"shouldShowSourceOfHireBadge"`
	LocationName                 *string         `json:"locationName,omitempty"`
	MultiLocaleTitle             MultiLocale     `json:"multiLocaleTitle"`
	EmploymentType               *Industry       `json:"employmentType,omitempty"`
	GeoUrn                       string          `json:"geoUrn"`
	ProfileTreasuryMediaPosition json.RawMessage `json:"profileTreasuryMediaPosition"`
	RegionUrn                    string          `json:"regionUrn"`
	GeoLocationName              string          `json:"geoLocationName"`
	RecipeType                   string          `json:"$recipeType"`
	MultiLocaleLocationName      *MultiLocale    `json:"multiLocaleLocationName,omitempty"`
	Description                  *string         `json:"description,omitempty"`
	MultiLocaleDescription       *MultiLocale    `json:"multiLocaleDescription,omitempty"`
}

type ProfilePositionInPositionGroup struct {
	Paging     Paging                                  `json:"paging"`
	RecipeType string                                  `json:"$recipeType"`
	Elements   []ProfilePositionInPositionGroupElement `json:"elements"`
}

type PositionGroup struct {
	EntityUrn                      string                         `json:"entityUrn"`
	DateRange                      DateRange                      `json:"dateRange"`
	MultiLocaleCompanyName         MultiLocale                    `json:"multiLocaleCompanyName"`
	CompanyName                    string                         `json:"companyName"`
	ProfilePositionInPositionGroup ProfilePositionInPositionGroup `json:"profilePositionInPositionGroup"`
	Company                        Company                        `json:"company"`
	RecipeType                     string                         `json:"$recipeType"`
	CompanyUrn                     string                         `json:"companyUrn"`
}

type EmployeeCountRange struct {
	Start      int64  `json:"start"`
	RecipeType string `json:"$recipeType"`
	End        *int64 `json:"end,omitempty"`
}
