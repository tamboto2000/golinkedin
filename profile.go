package golinkedin

import (
	"encoding/json"
	"net/url"
	"strings"
)

// ProfileNode contains user profile info
type ProfileNode struct {
	Elements []Profile `json:"elements,omitempty"`
	Paging   Paging    `json:"paging,omitempty"`

	ln *Linkedin
}

// Profile represent user profile
type Profile struct {
	BirthDateOn                           *Date              `json:"birthDateOn,omitempty"`
	ObjectUrn                             string             `json:"objectUrn,omitempty"`
	MultiLocaleLastName                   *MultiLocale       `json:"multiLocaleLastName,omitempty"`
	MultiLocaleFirstNamePronunciationHint *MultiLocale       `json:"multiLocaleFirstNamePronunciationHint,omitempty"`
	ProfileOrganizations                  *OrganizationNode  `json:"profileOrganizations,omitempty"`
	MultiLocaleLastNamePronunciationHint  *MultiLocale       `json:"multiLocaleLastNamePronunciationHint,omitempty"`
	ProfileEducations                     *EducationNode     `json:"profileEducations,omitempty"`
	MultiLocaleFirstName                  *MultiLocale       `json:"multiLocaleFirstName,omitempty"`
	PublicIdentifier                      string             `json:"publicIdentifier,omitempty"`
	ProfileCertifications                 *CertificationNode `json:"profileCertifications,omitempty"`
	TrackingID                            string             `json:"trackingId,omitempty"`
	MultiLocaleSummary                    MultiLocale        `json:"multiLocaleSummary,omitempty"`
	VersionTag                            string             `json:"versionTag,omitempty"`
	ProfilePicture                        *ProfilePicture    `json:"profilePicture,omitempty"`
	FirstName                             string             `json:"firstName,omitempty"`
	ProfileSkills                         *SkillNode         `json:"profileSkills,omitempty"`
	RecipeType                            string             `json:"$recipeType,omitempty"`
	MultiLocaleHeadline                   *MultiLocale       `json:"multiLocaleHeadline,omitempty"`
	ProfileHonors                         *HonorNode         `json:"profileHonors,omitempty"`
	Memorialized                          bool               `json:"memorialized,omitempty"`
	LastName                              string             `json:"lastName,omitempty"`
	VolunteerCauses                       []string           `json:"volunteerCauses,omitempty"`
	ShowPremiumSubscriberBadge            bool               `json:"showPremiumSubscriberBadge,omitempty"`
	Industry                              *Industry          `json:"industry,omitempty"`
	GeoLocationBackfilled                 bool               `json:"geoLocationBackfilled,omitempty"`
	MultiLocaleFullNamePronunciationAudio *MultiLocale       `json:"multiLocaleFullNamePronunciationAudio,omitempty"`
	Premium                               bool               `json:"premium,omitempty"`
	Influencer                            bool               `json:"influencer,omitempty"`
	EntityUrn                             string             `json:"entityUrn,omitempty"`
	Headline                              string             `json:"headline,omitempty"`
	Summary                               string             `json:"summary,omitempty"`
	SupportedLocales                      []Locale           `json:"supportedLocales,omitempty"`
	EducationOnProfileTopCardShown        bool               `json:"educationOnProfileTopCardShown,omitempty"`
	IndustryUrn                           string             `json:"industryUrn,omitempty"`
	ProfilePositionGroups                 *PositionGroupNode `json:"profilePositionGroups,omitempty"`
	GeoLocation                           *GeoLocation       `json:"geoLocation,omitempty"`
	Location                              *Location          `json:"location,omitempty"`
	BackgroundPicture                     *BackgroundPicture `json:"backgroundPicture,omitempty"`
	PrimaryLocale                         *Locale            `json:"primaryLocale,omitempty"`

	// Unsolved:
	// ProfileTreasuryMediaProfile
	// ProfilePublications
	// ProfileVolunteerExperiences
	// ProfileProjects
	// ProfilePatents
	// ProfileLanguages
	// ProfileCourses
}

// ProfilePicture contains multiple qualities of profile picture
type ProfilePicture struct {
	RecipeType            string                 `json:"$recipeType,omitempty"`
	DisplayImageReference *DisplayImageReference `json:"displayImageReference,omitempty"`
	DisplayImageUrn       string                 `json:"displayImageUrn,omitempty"`
	PhotoFilterEditInfo   *PhotoFilterEditInfo   `json:"photoFilterEditInfo,omitempty"`
}

// BackgroundPicture contains multiple qualities of background banner
type BackgroundPicture struct {
	RecipeType            string                 `json:"$recipeType,omitempty"`
	DisplayImageReference *DisplayImageReference `json:"displayImageReference,omitempty"`
	DisplayImageUrn       string                 `json:"displayImageUrn,omitempty"`
}

type ContactInfo struct {
	BirthDateOn               BirthDateOn   `json:"birthDateOn,omitempty"`
	EmailAddress              string        `json:"emailAddress,omitempty"`
	BirthdayVisibilitySetting string        `json:"birthdayVisibilitySetting,omitempty"`
	Address                   string        `json:"address,omitempty"`
	EntityUrn                 string        `json:"entityUrn,omitempty"`
	Websites                  []Website     `json:"websites,omitempty"`
	TwitterHandles            []interface{} `json:"twitterHandles,omitempty"`
	PhoneNumbers              []PhoneNumber `json:"phoneNumbers,omitempty"`
}

type BirthDateOn struct {
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

type PhoneNumber struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type Website struct {
	Type Type   `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Type struct {
	COMLinkedinVoyagerIdentityProfileCustomWebsite COMLinkedinVoyagerIdentityProfileCustomWebsite `json:"com.linkedin.voyager.identity.profile.CustomWebsite,omitempty"`
}

type COMLinkedinVoyagerIdentityProfileCustomWebsite struct {
	Label string `json:"label,omitempty"`
}

// Interest types
const (
	InterestCompany    = "COMPANY"
	InterestGroup      = "GROUP"
	InterestSchool     = "SCHOOL"
	InterestInfluencer = "INFLUENCER"
)

// Activity types
const (
	ActivityPost    = "POST"
	ActivityArticle = "ARTICLE"
)

// ProfileByUsername lookup profile with basic information by public identifier (username)
func (ln *Linkedin) ProfileByUsername(username string) (*ProfileNode, error) {
	q := make(url.Values)
	q.Add("q", "memberIdentity")
	q.Add("memberIdentity", username)
	q.Add("decorationId", "com.linkedin.voyager.dash.deco.identity.profile.FullProfileWithEntities-57")

	raw, err := ln.get("/identity/dash/profiles", q)
	if err != nil {
		return nil, err
	}

	profile := new(ProfileNode)
	if err := json.Unmarshal(raw, profile); err != nil {
		return nil, err
	}

	profile.ln = ln

	return profile, nil
}

// ProfileID return profile's ID, parsed from EntityURN
func (p *ProfileNode) ProfileID() string {
	return parseProfileID(p.Elements[0].EntityUrn)
}

// Connections return profile connections.
// You can perform this action by calling Linkedin.SearchPeople
func (p *ProfileNode) Connections() (*PeopleNode, error) {
	return p.ln.SearchPeople(
		"",
		&PeopleSearchFilter{
			Network:      []string{Rank1, Rank2, Rank3},
			ConnectionOf: p.ProfileID(),
			ResultType:   ResultPeople,
		},
		&QueryContext{
			SpellCorrectionEnabled: true,
		},
		OriginMemberProfileCannedSearch,
	)
}

// Organizations prepare OrganizarionNode for cursoring
func (p *ProfileNode) Organizations() *OrganizationNode {
	profileID := parseProfileID(p.Elements[0].EntityUrn)
	org := p.Elements[0].ProfileOrganizations
	org.ln = p.ln
	org.ProfileID = profileID

	return org
}

// Educations prepare EducationNode for cursoring
func (p *ProfileNode) Educations() *EducationNode {
	profileID := parseProfileID(p.Elements[0].EntityUrn)
	educations := p.Elements[0].ProfileEducations
	educations.ln = p.ln
	educations.ProfileID = profileID

	return educations
}

// Certifications prepare CertificationNode for cursoring
func (p *ProfileNode) Certifications() *CertificationNode {
	profileID := parseProfileID(p.Elements[0].EntityUrn)
	cert := p.Elements[0].ProfileCertifications
	cert.ln = p.ln
	cert.ProfileID = profileID

	return cert
}

// Skills prepare SkillNode for cursoring
func (p *ProfileNode) Skills() *SkillNode {
	profileID := parseProfileID(p.Elements[0].EntityUrn)
	s := p.Elements[0].ProfileSkills
	s.ln = p.ln
	s.ProfileID = profileID

	return s
}

// Honors prepare HonorNode for cursoring
func (p *ProfileNode) Honors() *HonorNode {
	profileID := parseProfileID(p.Elements[0].EntityUrn)
	honor := p.Elements[0].ProfileHonors
	honor.ln = p.ln
	honor.ProfileID = profileID

	return honor
}

// PositionGroups prepare PositionGroups for cursoring
func (p *ProfileNode) PositionGroups() *PositionGroupNode {
	profileID := parseProfileID(p.Elements[0].EntityUrn)
	post := p.Elements[0].ProfilePositionGroups
	post.ln = p.ln
	post.ProfileID = profileID

	return post
}

// ContactInfo get profile contact info
func (p *ProfileNode) ContactInfo() (*ContactInfo, error) {
	raw, err := p.ln.get("/identity/profiles/"+p.ProfileID()+"/profileContactInfo", nil)
	if err != nil {
		return nil, err
	}

	contact := new(ContactInfo)
	if err := json.Unmarshal(raw, contact); err != nil {
		return nil, err
	}

	return contact, nil
}

// GivenRecommendation get profile given recommendations
func (p *ProfileNode) GivenRecommendation() (*RecommendationNode, error) {
	raw, err := p.ln.get("/identity/profiles/"+p.ProfileID()+"/recommendations", url.Values{"q": {"given"}})
	if err != nil {
		return nil, err
	}

	recNode := new(RecommendationNode)
	if err := json.Unmarshal(raw, recNode); err != nil {
		return nil, err
	}

	recNode.ln = p.ln
	recNode.ProfileID = p.ProfileID()
	recNode.Q = "given"

	return recNode, nil
}

// ReceivedRecommendation get profile received recommendations
func (p *ProfileNode) ReceivedRecommendation() (*RecommendationNode, error) {
	raw, err := p.ln.get("/identity/profiles/"+p.ProfileID()+"/recommendations", url.Values{
		"q":                      {"received"},
		"recommendationStatuses": {"List(VISIBLE)"},
	})
	if err != nil {
		return nil, err
	}

	recNode := new(RecommendationNode)
	if err := json.Unmarshal(raw, recNode); err != nil {
		return nil, err
	}

	recNode.ln = p.ln
	recNode.ProfileID = p.ProfileID()
	recNode.Q = "received"

	return recNode, nil
}

// Interest get profile interests
func (p *ProfileNode) Interest(in string) (*InterestNode, error) {
	raw, err := p.ln.get("/identity/profiles/"+p.ProfileID()+"/following", url.Values{
		"count":      {"5"},
		"entityType": {in},
		"q":          {"followedEntities"},
	})

	if err != nil {
		return nil, err
	}

	intNode := new(InterestNode)
	if err := json.Unmarshal(raw, intNode); err != nil {
		return nil, err
	}

	intNode.ln = p.ln
	intNode.ProfileID = p.ProfileID()
	intNode.Type = in

	return intNode, nil
}

func (p *ProfileNode) Activity(in string) (*ActivityNode, error) {
	var raw []byte
	var err error
	if in == ActivityArticle {
		raw, err = p.ln.get("/identity/profiles/"+p.ProfileID()+"/posts", nil)
	}

	if in == ActivityPost {
		raw, err = p.ln.get("/identity/profileUpdatesV2", url.Values{
			"includeLongTermHistory": {"true"},
			"moduleKey":              {"member-shares:phone"},
			"numComments":            {"0"},
			"numLikes":               {"0"},
			"profileUrn":             {p.Elements[0].EntityUrn},
			"q":                      {"memberShareFeed"},
		})
	}

	if err != nil {
		return nil, err
	}

	actNode := new(ActivityNode)
	if err := json.Unmarshal(raw, actNode); err != nil {
		return nil, err
	}

	actNode.ProfileUrn = p.Elements[0].EntityUrn
	actNode.Type = in
	actNode.ln = p.ln

	return actNode, nil
}

func parseProfileID(entityUrn string) string {
	return strings.Replace(entityUrn, "urn:li:fsd_profile:", "", 1)
}
