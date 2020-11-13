package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type CertificationNode struct {
	ProfileID  string          `json:"profileId,omitempty"`
	Paging     Paging          `json:"paging,omitempty"`
	RecipeType string          `json:"$recipeType,omitempty"`
	Elements   []Certification `json:"elements,omitempty"`

	err error
	ln  *Linkedin
}

type Certification struct {
	DateRange                DateRange   `json:"dateRange,omitempty"`
	MultiLocaleLicenseNumber MultiLocale `json:"multiLocaleLicenseNumber,omitempty"`
	CompanyUrn               string      `json:"companyUrn,omitempty"`
	URL                      string      `json:"url,omitempty"`
	MultiLocaleAuthority     MultiLocale `json:"multiLocaleAuthority,omitempty"`
	EntityUrn                string      `json:"entityUrn,omitempty"`
	Authority                string      `json:"authority,omitempty"`
	Name                     string      `json:"name,omitempty"`
	MultiLocaleName          MultiLocale `json:"multiLocaleName,omitempty"`
	LicenseNumber            string      `json:"licenseNumber,omitempty"`
	Company                  Company     `json:"company,omitempty"`
	RecipeType               string      `json:"$recipeType,omitempty"`
	DisplaySource            string      `json:"displaySource,omitempty"`
}

// SetLinkedin set Linkedin client
func (c *CertificationNode) SetLinkedin(ln *Linkedin) {
	c.ln = ln
}

// Next cursoring educations.
// New certifications stored in CertificationNode.Elements
func (c *CertificationNode) Next() bool {
	start := strconv.Itoa(c.Paging.Start)
	count := strconv.Itoa(c.Paging.Count)
	raw, err := c.ln.get("/identity/profiles/"+c.ProfileID+"/certifications", url.Values{
		"start": {start},
		"count": {count},
	})

	if err != nil {
		c.err = err
		return false
	}

	certNode := new(CertificationNode)
	if err := json.Unmarshal(raw, certNode); err != nil {
		c.err = err
		return false
	}

	c.Elements = certNode.Elements
	c.Paging.Start = certNode.Paging.Start + certNode.Paging.Count

	if len(c.Elements) == 0 {
		return false
	}

	return true
}

func (c *CertificationNode) Error() error {
	return c.err
}
