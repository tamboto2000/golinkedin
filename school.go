package golinkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type SchoolNode struct {
	Keywords string   `json:"keywords,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
	Elements []School `json:"elements,omitempty"`
	Paging   Paging   `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type School struct {
	// Element contains schools from search school result
	Elements                                                []School                    `json:"elements,omitempty"`
	ExtendedElements                                        []interface{}               `json:"extendedElements,omitempty"`
	ObjectUrn                                               string                      `json:"objectUrn,omitempty"`
	EntityUrn                                               string                      `json:"entityUrn,omitempty"`
	Active                                                  bool                        `json:"active,omitempty"`
	Logo                                                    *Logo                       `json:"logo,omitempty"`
	SchoolName                                              string                      `json:"schoolName,omitempty"`
	TrackingID                                              string                      `json:"trackingId,omitempty"`
	Image                                                   Image                       `json:"image,omitempty"`
	Subtext                                                 Text                        `json:"subtext,omitempty"`
	TargetUrn                                               string                      `json:"targetUrn,omitempty"`
	Text                                                    Text                        `json:"text,omitempty"`
	DashTargetUrn                                           string                      `json:"dashTargetUrn,omitempty"`
	Type                                                    string                      `json:"type,omitempty"`
	TrackingUrn                                             string                      `json:"trackingUrn,omitempty"`
	Title                                                   Title                       `json:"title,omitempty"`
	AntiAbuseAnnotations                                    []AntiAbuseAnnotation       `json:"$anti_abuse_annotations"`
	Name                                                    string                      `json:"name"`
	RecipeType                                              string                      `json:"$recipeType"`
	URL                                                     string                      `json:"url"`
	Industry                                                map[string]Industry         `json:"industry,omitempty"`
	IndustryUrns                                            []string                    `json:"industryUrns,omitempty"`
	MiniCompany                                             *MiniCompany                `json:"miniCompany,omitempty"`
	EmployeeCountRange                                      *EmployeeCountRange         `json:"employeeCountRange,omitempty"`
	Industries                                              []string                    `json:"industries,omitempty"`
	UniversalName                                           string                      `json:"universalName,omitempty"`
	Showcase                                                bool                        `json:"showcase,omitempty"`
	Headline                                                *Text                       `json:"headline,omitempty"`
	Subline                                                 *Text                       `json:"subline,omitempty"`
	VideosTabVisible                                        bool                        `json:"videosTabVisible,omitempty"`
	StaffCount                                              int                         `json:"staffCount,omitempty"`
	CompanyEmployeesSearchPageURL                           string                      `json:"companyEmployeesSearchPageUrl,omitempty"`
	ViewerFollowingJobsUpdates                              bool                        `json:"viewerFollowingJobsUpdates,omitempty"`
	MyCompanyVisible                                        bool                        `json:"myCompanyVisible,omitempty"`
	Permissions                                             *Permissions                `json:"permissions,omitempty"`
	EligibleForOnlineJobPostingEntry                        bool                        `json:"eligibleForOnlineJobPostingEntry,omitempty"`
	FollowingInfo                                           *FollowingInfo              `json:"followingInfo,omitempty"`
	ViewerEmployee                                          bool                        `json:"viewerEmployee,omitempty"`
	AffiliatedCompaniesWithEmployeesRollup                  []interface{}               `json:"affiliatedCompaniesWithEmployeesRollup,omitempty"`
	AffiliatedCompaniesWithJobsRollup                       []interface{}               `json:"affiliatedCompaniesWithJobsRollup,omitempty"`
	OrganizationLixes                                       []interface{}               `json:"organizationLixes,omitempty"`
	Tagline                                                 string                      `json:"tagline,omitempty"`
	ViewerCurrentEmployee                                   bool                        `json:"viewerCurrentEmployee,omitempty"`
	MultiLocaleTaglines                                     *MultiLocaleTaglines        `json:"multiLocaleTaglines,omitempty"`
	Headquarter                                             *Location                   `json:"headquarter,omitempty"`
	PublishedProductsOwner                                  bool                        `json:"publishedProductsOwner,omitempty"`
	CompanyPageURL                                          string                      `json:"companyPageUrl,omitempty"`
	ViewerConnectedToAdministrator                          bool                        `json:"viewerConnectedToAdministrator,omitempty"`
	DataVersion                                             int                         `json:"dataVersion,omitempty"`
	AssociatedHashtags                                      []string                    `json:"associatedHashtags,omitempty"`
	ShowcasePages                                           []interface{}               `json:"showcasePages,omitempty"`
	ClaimableByViewer                                       bool                        `json:"claimableByViewer,omitempty"`
	JobSearchPageURL                                        string                      `json:"jobSearchPageUrl,omitempty"`
	AutoGenerated                                           bool                        `json:"autoGenerated,omitempty"`
	ViewerPermissions                                       *ViewerPermissions          `json:"viewerPermissions,omitempty"`
	StaffingCompany                                         bool                        `json:"staffingCompany,omitempty"`
	CompanyIndustries                                       []Industry                  `json:"companyIndustries,omitempty"`
	CallToAction                                            *CallToAction               `json:"callToAction,omitempty"`
	AdsRule                                                 string                      `json:"adsRule,omitempty"`
	StaffCountRange                                         *CountRange                 `json:"staffCountRange,omitempty"`
	Claimable                                               bool                        `json:"claimable,omitempty"`
	Specialities                                            []string                    `json:"specialities,omitempty"`
	ConfirmedLocations                                      []Location                  `json:"confirmedLocations,omitempty"`
	VersionTag                                              string                      `json:"versionTag,omitempty"`
	LcpTreatment                                            bool                        `json:"lcpTreatment,omitempty"`
	AssociatedHashtagsResolutionResults                     *json.RawMessage            `json:"associatedHashtagsResolutionResults,omitempty"`
	EmployeeExperienceSettings                              *EmployeeExperienceSettings `json:"employeeExperienceSettings,omitempty"`
	Description                                             string                      `json:"description,omitempty"`
	PaidCompany                                             bool                        `json:"paidCompany,omitempty"`
	ViewerPendingAdministrator                              bool                        `json:"viewerPendingAdministrator,omitempty"`
	AffiliatedCompanies                                     []interface{}               `json:"affiliatedCompanies,omitempty"`
	FoundedOn                                               *Date                       `json:"foundedOn,omitempty"`
	CompanyType                                             *CompanyType                `json:"companyType,omitempty"`
	Groups                                                  []interface{}               `json:"groups,omitempty"`
	EventsTabVisible                                        bool                        `json:"eventsTabVisible,omitempty"`
	BackgroundCoverImage                                    *BackgroundCoverImage       `json:"backgroundCoverImage,omitempty"`
	AffiliatedCompaniesWithEmployeesRollupResolutionResults map[string]Company          `json:"affiliatedCompaniesWithEmployeesRollupResolutionResults,omitempty"`
	Phone                                                   Phone                       `json:"phone,omitempty"`
	OverviewPhoto                                           Photo                       `json:"overviewPhoto,omitempty"`
	CoverPhoto                                              Photo                       `json:"coverPhoto,omitempty"`
	School                                                  string                      `json:"school,omitempty"`
	AffiliatedCompaniesResolutionResults                    map[string]Company          `json:"affiliatedCompaniesResolutionResults,omitempty"`
	ShowcasePagesResolutionResults                          map[string]Company          `json:"showcasePagesResolutionResults,omitempty"`
	GroupsResolutionResults                                 map[string]Group
}

// SchoolByName lookup school by universal name.
func (ln *Linkedin) SchoolByName(name string) (*SchoolNode, error) {
	raw, err := ln.get("/organization/companies", url.Values{
		"decorationId":  {"com.linkedin.voyager.deco.organization.web.WebFullCompanyMain-28"},
		"q":             {"universalName"},
		"universalName": {name},
	})

	if err != nil {
		return nil, err
	}

	schNode := new(SchoolNode)
	if err := json.Unmarshal(raw, schNode); err != nil {
		return nil, err
	}

	return schNode, nil
}

func (sch *SchoolNode) SetLinkedin(ln *Linkedin) {
	sch.ln = ln
}

func (sch *SchoolNode) Next() bool {
	if sch.stopCursor {
		return false
	}

	start := strconv.Itoa(sch.Paging.Start)
	count := strconv.Itoa(sch.Paging.Count)
	raw, err := sch.ln.get("/search/blended", url.Values{
		"keywords":     {sch.Keywords},
		"origin":       {OriginSwitchSearchVertical},
		"q":            {QAll},
		"start":        {start},
		"count":        {count},
		"filters":      {composeFilter(DefaultSearchSchoolFilter)},
		"queryContext": {composeFilter(DefaultSearchSchoolQueryContext)},
	})

	if err != nil {
		sch.err = err
		return false
	}

	schNode := new(SchoolNode)
	if err := json.Unmarshal(raw, schNode); err != nil {
		sch.err = err
		return false
	}

	sch.Elements = schNode.Elements
	sch.Paging.Start = schNode.Paging.Start + schNode.Paging.Count

	if len(sch.Elements) == 0 {
		return false
	}

	if len(sch.Elements[0].Elements) < sch.Paging.Count {
		sch.stopCursor = true
	}

	return true
}

func (sch *SchoolNode) Error() error {
	return sch.err
}
