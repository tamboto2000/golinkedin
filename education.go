package golinkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// EducationNode contains user educations info
type EducationNode struct {
	// will set after calling ProfileNode.Educations()
	ProfileID string      `json:"profileId,omitempty"`
	Elements  []Education `json:"elements,omitempty"`
	Paging    Paging      `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Education struct {
	EntityUrn                     string                         `json:"entityUrn,omitempty"`
	Activities                    string                         `json:"activities,omitempty"`
	School                        *School                        `json:"school,omitempty"`
	TimePeriod                    *TimePeriod                    `json:"timePeriod,omitempty"`
	Grade                         string                         `json:"grade,omitempty"`
	Description                   string                         `json:"description,omitempty"`
	DegreeName                    string                         `json:"degreeName,omitempty"`
	SchoolName                    string                         `json:"schoolName,omitempty"`
	FieldOfStudy                  string                         `json:"fieldOfStudy,omitempty"`
	Recommendations               []interface{}                  `json:"recommendations,omitempty"`
	SchoolUrn                     string                         `json:"schoolUrn,omitempty"`
	DateRange                     *DateRange                     `json:"dateRange,omitempty"`
	ProfileTreasuryMediaEducation *ProfileTreasuryMediaEducation `json:"profileTreasuryMediaEducation,omitempty"`
	MultiLocaleSchoolName         *MultiLocale                   `json:"multiLocaleSchoolName,omitempty"`
	MultiLocaleFieldOfStudy       *MultiLocale                   `json:"multiLocaleFieldOfStudy,omitempty"`
	RecipeType                    string                         `json:"$recipeType,omitempty"`
	MultiLocaleDescription        *MultiLocale                   `json:"multiLocaleDescription,omitempty"`
	MultiLocaleActivities         *MultiLocale                   `json:"multiLocaleActivities,omitempty"`
	MultiLocaleDegreeName         *MultiLocale                   `json:"multiLocaleDegreeName,omitempty"`
}

type School struct {
	ObjectUrn     string `json:"objectUrn,omitempty"`
	EntityUrn     string `json:"entityUrn,omitempty"`
	Active        bool   `json:"active,omitempty"`
	Logo          *Logo  `json:"logo,omitempty"`
	SchoolName    string `json:"schoolName,omitempty"`
	TrackingID    string `json:"trackingId,omitempty"`
	Image         Image  `json:"image,omitempty"`
	Subtext       Text   `json:"subtext,omitempty"`
	TargetUrn     string `json:"targetUrn,omitempty"`
	Text          Text   `json:"text,omitempty"`
	DashTargetUrn string `json:"dashTargetUrn,omitempty"`
	Type          string `json:"type,omitempty"`
	TrackingUrn   string `json:"trackingUrn,omitempty"`
	Title         Title  `json:"title,omitempty"`
}

type Title struct {
	TextDirection string `json:"textDirection,omitempty"`
	Text          string `json:"text,omitempty"`
}

type ProfileTreasuryMediaEducation struct {
	Paging     Paging        `json:"paging,omitempty"`
	RecipeType string        `json:"$recipeType,omitempty"`
	Elements   []interface{} `json:"elements,omitempty"`
}

func (edu *EducationNode) SetLinkedin(ln *Linkedin) {
	edu.ln = ln
}

func (edu *EducationNode) Next() bool {
	if edu.stopCursor {
		return false
	}

	start := strconv.Itoa(edu.Paging.Start)
	count := strconv.Itoa(edu.Paging.Count)
	raw, err := edu.ln.get("/identity/profiles/"+edu.ProfileID+"/educations", url.Values{
		"start": {start},
		"count": {count},
	})

	if err != nil {
		edu.err = err
		return false
	}

	eduNode := new(EducationNode)
	if err := json.Unmarshal(raw, eduNode); err != nil {
		edu.err = err
		return false
	}

	edu.Elements = eduNode.Elements
	edu.Paging.Start = eduNode.Paging.Start + eduNode.Paging.Count

	if len(edu.Elements) == 0 {
		return false
	}

	if len(edu.Elements) < edu.Paging.Count {
		edu.stopCursor = true
	}

	return true
}

func (edu *EducationNode) Error() error {
	return edu.err
}
