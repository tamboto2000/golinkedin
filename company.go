package golinkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type CompanyNode struct {
	Metadata Metadata  `json:"metadata,omitempty"`
	Elements []Company `json:"elements,omitempty"`
	Paging   Paging    `json:"paging,omitempty"`
	Keywords string    `json:"keywords,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

// Company contain information about a company.
// Company can also represent a school since school is also a company, see Linkedin.SchoolByName
type Company struct {
	// Elements contains companies from search company result
	Elements                                                []Company                   `json:"elements,omitempty"`
	ExtendedElements                                        []interface{}               `json:"extendedElements,omitempty"`
	Industry                                                map[string]Industry         `json:"industry,omitempty"`
	IndustryUrns                                            []string                    `json:"industryUrns,omitempty"`
	AntiAbuseAnnotations                                    []AntiAbuseAnnotation       `json:"$anti_abuse_annotations,omitempty"`
	EntityUrn                                               string                      `json:"entityUrn,omitempty"`
	MiniCompany                                             *MiniCompany                `json:"miniCompany,omitempty"`
	EmployeeCountRange                                      *EmployeeCountRange         `json:"employeeCountRange,omitempty"`
	Industries                                              []string                    `json:"industries,omitempty"`
	Name                                                    string                      `json:"name,omitempty"`
	Logo                                                    *Logo                       `json:"logo,omitempty"`
	RecipeType                                              string                      `json:"$recipeType,omitempty"`
	UniversalName                                           string                      `json:"universalName,omitempty"`
	URL                                                     string                      `json:"url,omitempty"`
	ObjectUrn                                               string                      `json:"objectUrn,omitempty"`
	Showcase                                                bool                        `json:"showcase,omitempty"`
	Active                                                  bool                        `json:"active,omitempty"`
	TrackingID                                              string                      `json:"trackingId,omitempty"`
	Image                                                   *Image                      `json:"image,omitempty"`
	Subtext                                                 *Text                       `json:"subtext,omitempty"`
	TargetUrn                                               string                      `json:"targetUrn,omitempty"`
	Text                                                    *Text                       `json:"text,omitempty"`
	DashTargetUrn                                           string                      `json:"dashTargetUrn,omitempty"`
	Type                                                    string                      `json:"type,omitempty"`
	TrackingUrn                                             string                      `json:"trackingUrn,omitempty"`
	Title                                                   *Title                      `json:"title,omitempty"`
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
	GroupsResolutionResults                                 map[string]Group            `json:"groupsResolutionResults,omitempty"`
}

type Photo struct {
	COMLinkedinVoyagerCommonMediaProcessorImage MediaProcessorImage `json:"com.linkedin.voyager.common.MediaProcessorImage,omitempty"`
}

type MediaProcessorImage struct {
	ID string `json:"id,omitempty"`
}

type Phone struct {
	Number string `json:"number,omitempty"`
}

type Permissions struct {
	LandingPageAdmin bool `json:"landingPageAdmin,omitempty"`
	Admin            bool `json:"admin,omitempty"`
	AdAccountHolder  bool `json:"adAccountHolder,omitempty"`
}

type MultiLocaleTaglines struct {
	Localized       MultiLocale `json:"localized,omitempty"`
	PreferredLocale Locale      `json:"preferredLocale,omitempty"`
}

type ViewerPermissions struct {
	CanReadPipelineBuilderAdminPage      bool   `json:"canReadPipelineBuilderAdminPage,omitempty"`
	CanCreateOrganicShare                bool   `json:"canCreateOrganicShare,omitempty"`
	CanUntagFromMention                  bool   `json:"canUntagFromMention,omitempty"`
	CanReadOrganizationVisitorAnalytics  bool   `json:"canReadOrganizationVisitorAnalytics,omitempty"`
	CanCreateComment                     bool   `json:"canCreateComment,omitempty"`
	CanDeleteShare                       bool   `json:"canDeleteShare,omitempty"`
	CanCreateReaction                    bool   `json:"canCreateReaction,omitempty"`
	CanEnableCommentsShare               bool   `json:"canEnableCommentsShare,omitempty"`
	CanDisableCommentsShare              bool   `json:"canDisableCommentsShare,omitempty"`
	CanDeleteDarkShare                   bool   `json:"canDeleteDarkShare,omitempty"`
	RecipeType                           string `json:"$recipeType,omitempty"`
	CanSeeOrganizationAdministrativePage bool   `json:"canSeeOrganizationAdministrativePage,omitempty"`
}

type CallToAction struct {
	CallToActionType    string              `json:"callToActionType,omitempty"`
	Visible             bool                `json:"visible,omitempty"`
	CallToActionMessage CallToActionMessage `json:"callToActionMessage,omitempty"`
	URL                 string              `json:"url,omitempty"`
}

type CallToActionMessage struct {
	TextDirection string `json:"textDirection,omitempty"`
	Text          string `json:"text,omitempty"`
}

type EmployeeExperienceSettings struct {
	MyCompanyVisibility                   bool   `json:"myCompanyVisibility,omitempty"`
	PymkVisibility                        string `json:"pymkVisibility,omitempty"`
	MyCompanyEmployeeVerificationRequired bool   `json:"myCompanyEmployeeVerificationRequired,omitempty"`
	HighlightsVisibility                  string `json:"highlightsVisibility,omitempty"`
	TrendingContentVisibility             string `json:"trendingContentVisibility,omitempty"`
	BroadcastsVisibility                  string `json:"broadcastsVisibility,omitempty"`
}

type CompanyType struct {
	LocalizedName string `json:"localizedName,omitempty"`
	Code          string `json:"code,omitempty"`
}

type BackgroundCoverImage struct {
	Image    Image    `json:"image,omitempty"`
	CropInfo CropInfo `json:"cropInfo,omitempty"`
}

type CropInfo struct {
	X      int `json:"x,omitempty"`
	Width  int `json:"width,omitempty"`
	Y      int `json:"y,omitempty"`
	Height int `json:"height,omitempty"`
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

func (ln *Linkedin) CompanyByName(name string) (*CompanyNode, error) {
	raw, err := ln.get("/organization/companies", url.Values{
		"decorationId":  {"com.linkedin.voyager.deco.organization.web.WebFullCompanyMain-28"},
		"q":             {"universalName"},
		"universalName": {name},
	})

	if err != nil {
		return nil, err
	}

	compNode := new(CompanyNode)
	if err := json.Unmarshal(raw, compNode); err != nil {
		return nil, err
	}

	return compNode, nil
}

func (comp *CompanyNode) SetLinkedin(ln *Linkedin) {
	comp.ln = ln
}

func (comp *CompanyNode) Next() bool {
	if comp.stopCursor {
		return false
	}

	start := strconv.Itoa(comp.Paging.Start)
	count := strconv.Itoa(comp.Paging.Count)
	raw, err := comp.ln.get("/search/blended", url.Values{
		"keywords":     {comp.Keywords},
		"origin":       {OriginSwitchSearchVertical},
		"q":            {QAll},
		"filters":      {composeFilter(DefaultSearchCompanyFilter)},
		"queryContext": {composeFilter(DefaultSearchCompanyQueryContext)},
		"start":        {start},
		"count":        {count},
	})

	if err != nil {
		comp.err = err
		return false
	}

	compNode := new(CompanyNode)
	if err := json.Unmarshal(raw, compNode); err != nil {
		comp.err = err
		return false
	}

	comp.Elements = compNode.Elements
	comp.Paging.Start = compNode.Paging.Start + compNode.Paging.Count

	if len(comp.Elements) == 0 {
		return false
	}

	if len(comp.Elements[0].Elements) < comp.Paging.Count {
		comp.stopCursor = true
	}

	return true
}

func (comp *CompanyNode) Error() error {
	return comp.err
}
