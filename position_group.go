package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type PositionGroupNode struct {
	ProfileID  string          `json:"profileId,omitempty"`
	Paging     Paging          `json:"paging,omitempty"`
	RecipeType string          `json:"$recipeType,omitempty"`
	Elements   []PositionGroup `json:"elements,omitempty"`

	err error
	ln  *Linkedin
}

type PositionGroup struct {
	EntityUrn                      string             `json:"entityUrn,omitempty"`
	DateRange                      *DateRange         `json:"dateRange,omitempty"`
	MultiLocaleCompanyName         *MultiLocale       `json:"multiLocaleCompanyName,omitempty"`
	CompanyName                    string             `json:"companyName,omitempty"`
	ProfilePositionInPositionGroup *PositionGroupNode `json:"profilePositionInPositionGroup,omitempty"`
	Company                        *Company           `json:"company,omitempty"`
	RecipeType                     string             `json:"$recipeType,omitempty"`
	CompanyUrn                     string             `json:"companyUrn,omitempty"`
	Title                          string             `json:"title,omitempty"`
	EmploymentTypeUrn              string             `json:"employmentTypeUrn,omitempty"`
	MultiLocaleGeoLocationName     *MultiLocale       `json:"multiLocaleGeoLocationName,omitempty"`
	ShouldShowSourceOfHireBadge    bool               `json:"shouldShowSourceOfHireBadge,omitempty"`
	LocationName                   string             `json:"locationName,omitempty"`
	MultiLocaleTitle               *MultiLocale       `json:"multiLocaleTitle,omitempty"`
	EmploymentType                 *Industry          `json:"employmentType,omitempty"`
	GeoUrn                         string             `json:"geoUrn,omitempty"`
	ProfileTreasuryMediaPosition   *PositionGroupNode `json:"profileTreasuryMediaPosition,omitempty"`
	RegionUrn                      string             `json:"regionUrn,omitempty"`
	GeoLocationName                string             `json:"geoLocationName,omitempty"`
	MultiLocaleLocationName        *MultiLocale       `json:"multiLocaleLocationName,omitempty"`
	Description                    string             `json:"description,omitempty"`
	MultiLocaleDescription         *MultiLocale       `json:"multiLocaleDescription,omitempty"`
	MiniCompany                    *MiniCompany       `json:"miniCompany,omitempty"`
	TimePeriod                     *TimePeriod        `json:"timePeriod,omitempty"`
	Name                           string             `json:"name,omitempty"`
	Positions                      []Position         `json:"positions,omitempty"`
	Paging                         *Paging            `json:"paging,omitempty"`
}

type Position struct {
	LocationName    string     `json:"locationName,omitempty"`
	EntityUrn       string     `json:"entityUrn,omitempty"`
	GeoLocationName string     `json:"geoLocationName,omitempty"`
	GeoUrn          string     `json:"geoUrn,omitempty"`
	CompanyName     string     `json:"companyName,omitempty"`
	TimePeriod      TimePeriod `json:"timePeriod,omitempty"`
	Company         Company    `json:"company,omitempty"`
	Title           string     `json:"title,omitempty"`
	Region          string     `json:"region,omitempty"`
	CompanyUrn      string     `json:"companyUrn,omitempty"`
	Recommendations []string   `json:"recommendations,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Honors          []string   `json:"honors,omitempty"`
}

func (post *PositionGroupNode) SetLinkedin(ln *Linkedin) {
	post.ln = ln
}

func (post *PositionGroupNode) Next() bool {
	start := strconv.Itoa(post.Paging.Start)
	count := strconv.Itoa(post.Paging.Count)
	raw, err := post.ln.get("/identity/profiles/"+post.ProfileID+"/positionGroups", url.Values{
		"start": {start},
		"count": {count},
	})

	if err != nil {
		post.err = err
		return false
	}

	postNode := new(PositionGroupNode)
	if err := json.Unmarshal(raw, postNode); err != nil {
		post.err = err
		return false
	}

	post.Elements = postNode.Elements
	post.Paging.Start = postNode.Paging.Start + postNode.Paging.Count

	if len(post.Elements) == 0 {
		return false
	}

	return true
}

func (post *PositionGroupNode) Error() error {
	return post.err
}
