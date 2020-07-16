package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/tamboto2000/linkedin/raw"
)

type Certification struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Authority *Company `json:"authority"`
	FromMonth int      `json:"fromMonth"`
	FromYear  int      `json:"fromYear"`
	ToMonth   int      `json:"toMonth"`
	ToYear    int      `json:"toYear"`
	URL       string   `json:"url"`
}

func (p *Profile) SyncCertification() error {
	return p.syncCertifications()
}

func (p *Profile) syncCertifications() error {
	urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/certifications?start=0&count=50"
	urlParsed, _ := url.Parse(urlStr)
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

	certificationView := new(raw.CertificationsView)
	if err = json.Unmarshal(body, certificationView); err != nil {
		return err
	}

	certifications := make([]Certification, 0)
	for _, e := range certificationView.Elements {
		certification := Certification{
			ID:        e.LicenseNumber,
			Name:      e.Name,
			FromMonth: int(e.TimePeriod.StartDate.Month),
			FromYear:  int(e.TimePeriod.StartDate.Year),
			ToMonth:   int(e.TimePeriod.EndDate.Month),
			ToYear:    int(e.TimePeriod.EndDate.Year),
			URL:       e.URL,
		}

		var companyLogoURL string
		companyLogoraw := e.Company.Logo.COMLinkedinCommonVectorImage
		if companyLogoraw.Artifacts != nil && len(companyLogoraw.Artifacts) > 0 {
			companyLogoURL = companyLogoraw.RootURL
			companyLogoURL += companyLogoraw.Artifacts[len(companyLogoraw.Artifacts)-1].FileIdentifyingURLPathSegment
		}

		companyID := strings.Replace(e.Company.EntityUrn, "urn:li:fs_miniCompany:", "", 1)
		companyURL := "https://www.linkedin.com/company/" + companyID
		certification.Authority = &Company{
			ID:            companyID,
			UniversalName: e.Company.UniversalName,
			Name:          e.Company.Name,
			LogoURL:       companyLogoURL,
			Active:        e.Company.Active,
			URL:           companyURL,
		}

		certifications = append(certifications, certification)
	}

	p.Certifications = certifications

	return nil
}
