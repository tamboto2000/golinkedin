package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/tamboto2000/linkedin/raw"
)

type Job struct {
	Location       *Location `json:"location,omitempty" csv:"-"`
	CompanyName    string    `json:"companyName,omitempty" csv:"company_name"`
	CompanyURL     string    `json:"url,omitempty" csv:"company_url"`
	CompanyLogoURL string    `json:"companyLogoUrl"`
	JobTitle       string    `json:"jobTitle,omitempty" csv:"job_title"`
	FromYear       int       `json:"fromYear" csv:"from_year"`
	FromMonth      int       `json:"fromMonth" csv:"from_month"`
	ToYear         int       `json:"toYear" csv:"to_year"`
	ToMonth        int       `json:"toMonth" csv:"to_month"`
	Description    string    `json:"description,omitempty" csv:"description"`
}

func (p *Profile) SyncJob() error {
	return p.syncJob()
}

func (p *Profile) syncJob() error {
	urlParsed, _ := url.Parse("https://www.linkedin.com/voyager/api/identity/profiles/" + p.ID + "/positionGroups?start=0&count=50")
	resp, err := p.ln.client.getRequest(urlParsed)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 200 {
		return errors.New("linkedin API error: " + string(body))
	}

	positionView := new(raw.PositionView)
	if err = json.Unmarshal(body, positionView); err != nil {
		return err
	}

	jobs := make([]Job, 0)
	for _, pst := range positionView.Elements {
		var logoURL string
		if pst.MiniCompany != nil {
			logoRaw := pst.MiniCompany.Logo.COMLinkedinCommonVectorImage
			if logoRaw.Artifacts != nil && len(logoRaw.Artifacts) > 0 {
				logoURL = logoRaw.RootURL
				logoURL += logoRaw.Artifacts[len(logoRaw.Artifacts)-1].FileIdentifyingURLPathSegment
			}
		}

		for _, subPst := range pst.Positions {
			job := Job{}
			if subPst.LocationName != "" || subPst.GeoLocationName != "" {
				job.Location = new(Location)
				if subPst.LocationName != "" {
					job.Location.LocationName = subPst.LocationName
				}

				if subPst.GeoLocationName != "" {
					job.Location.FullAddress = subPst.GeoLocationName
				}
			}

			job.CompanyName = subPst.CompanyName
			if subPst.CompanyUrn != "" {
				companyID := strings.Replace(subPst.CompanyUrn, "urn:li:fs_miniCompany:", "", 1)
				job.CompanyURL = "https://www.linkedin.com/company/" + companyID
			}

			job.CompanyLogoURL = logoURL
			job.JobTitle = subPst.Title
			job.FromYear = int(subPst.TimePeriod.StartDate.Year)
			job.FromMonth = int(subPst.TimePeriod.StartDate.Month)
			job.ToYear = int(subPst.TimePeriod.EndDate.Year)
			job.ToMonth = int(subPst.TimePeriod.EndDate.Month)
			job.Description = subPst.Description

			jobs = append(jobs, job)
		}
	}

	p.Job = jobs

	return nil
}
