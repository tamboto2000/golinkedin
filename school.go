package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"

	"github.com/tamboto2000/linkedin/raw"
)

type School struct {
	URL          string `json:"url,omitempty" csv:"url"`
	Logo         string `json:"logoImg,omitempty" csv:"logo"`
	Name         string `json:"name,omitempty" csv:"name"`
	Activities   string `json:"activities"`
	Grade        string `json:"grade"`
	Description  string `json:"description"`
	DateRange    string `json:"dateRange,omitempty" csv:"date_range"`
	Degree       string `json:"degree,omitempty" csv:"degree"`
	FieldOfStudy string `json:"fieldOfStudy" csv:"field_of_study"`
}

func (p *Profile) SyncSchool() error {
	return p.syncSchool()
}

func (p *Profile) syncSchool() error {
	urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.ID + "/educations?start=0&count=20"
	urlParsed, _ := url.Parse(urlStr)
	resp, err := p.ln.client.getRequest(urlParsed)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 200 {
		return errors.New("linkedin API error: " + string(body))
	}

	educationView := new(raw.EducationViewClass)
	if err = json.Unmarshal(body, educationView); err != nil {
		return err
	}

	p.School = composeSchool(educationView)

	return nil
}

func composeSchool(schoolView *raw.EducationViewClass) []School {
	schools := make([]School, 0)
	for _, s := range schoolView.Elements {
		school := School{}
		schoolID := strings.Replace(s.SchoolUrn, "urn:li:fs_miniSchool:", "", 1)
		school.URL = "https://www.linkedin.com/school/" + schoolID + "/?legacySchoolId=" + schoolID
		var logoURL string
		logoRaw := s.School.Logo.COMLinkedinCommonVectorImage
		if logoRaw.Artifacts != nil && len(logoRaw.Artifacts) > 0 {
			logoURL = logoRaw.RootURL
			logoURL += logoRaw.Artifacts[len(logoRaw.Artifacts)-1].FileIdentifyingURLPathSegment
		}

		school.Logo = logoURL
		school.Name = s.SchoolName
		school.DateRange = strconv.FormatInt(s.TimePeriod.StartDate.Year, 10) + " - " + strconv.FormatInt(s.TimePeriod.EndDate.Year, 10)
		school.Degree = s.DegreeName
		school.FieldOfStudy = s.FieldOfStudy
		school.Activities = s.Activities
		school.Grade = s.Grade
		school.Description = s.Description

		schools = append(schools, school)
	}

	return schools
}
