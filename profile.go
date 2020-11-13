package linkedin

import (
	"encoding/json"
	"net/url"
	"strings"
)

// ProfileNode contains user profile info
type ProfileNode struct {
	Elements []Profile `json:"elements,omitempty"`
	Paging   Paging    `json:"paging"`

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
	ProfileCertifications                 CertificationNode  `json:"profileCertifications,omitempty"`
	TrackingID                            string             `json:"trackingId,omitempty"`
	MultiLocaleSummary                    MultiLocale        `json:"multiLocaleSummary,omitempty"`
	VersionTag                            string             `json:"versionTag,omitempty"`
	ProfilePicture                        *ProfilePicture    `json:"profilePicture,omitempty"`
	FirstName                             string             `json:"firstName,omitempty"`
	ProfileSkills                         *SkillNode         `json:"profileSkills,omitempty"`
	RecipeType                            string             `json:"$recipeType,omitempty"`
	MultiLocaleHeadline                   *MultiLocale       `json:"multiLocaleHeadline,omitempty"`
	ProfileHonors                         *HonorNode         `json:"ProfileHonors,omitempty"`
	Memorialized                          bool               `json:"memorialized,omitempty"`
	LastName                              string             `json:"lastName,omitempty"`
	VolunteerCauses                       []string           `json:"volunteerCauses"`
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
	PrimaryLocale                         *Locale            `json:"primaryLocale"`

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

func (p *ProfileNode) Educations() *EducationNode {
	profileID := strings.Replace(p.Elements[0].EntityUrn, "urn:li:fsd_profile:", "", 1)
	educations := p.Elements[0].ProfileEducations
	educations.ln = p.ln
	educations.ProfileID = profileID

	return educations
}

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
