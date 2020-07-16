package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/tamboto2000/linkedin/raw"
)

type DOB struct {
	Year  int `json:"year,omitempty" csv:"year"`
	Month int `json:"month" csv:"month"`
	Day   int `json:"day" csv:"day"`
}

type Website struct {
	URL  string `json:"url,omitempty" csv:"url"`
	Type string `json:"type,omitempty" csv:"type"`
}

type Phone struct {
	Number string `json:"number,omitempty" csv:"number"`
	Type   string `json:"type,omitempty" csv:"type"`
}

type ContactInfo struct {
	DOB         *DOB      `json:"dob,omitempty" csv:"-"`
	Email       string    `json:"email,omitempty" csv:"email"`
	Address     string    `json:"address,omitempty" csv:"address"`
	Websites    []Website `json:"websites,omitempty" csv:"-"`
	Twitter     []string  `json:"twitter,omitempty" csv:"-"`
	PhoneNumber []Phone   `json:"phoneNumber,omitempty" csv:"-"`
	ConnectedAt string    `json:"connectedAt"`
}

func (p *Profile) SyncContact() error {
	return p.syncContact()
}

func (p *Profile) syncContact() error {
	urlParsed, _ := url.Parse("https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/profileContactInfo")
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

	cr := new(raw.ContactView)
	if err = json.Unmarshal(body, cr); err != nil {
		return err
	}

	contactInfo := ContactInfo{}

	contactInfo.DOB = &DOB{
		Year:  cr.BirthDateOn.Year,
		Month: cr.BirthDateOn.Month,
		Day:   cr.BirthDateOn.Day,
	}

	contactInfo.Email = cr.EmailAddress
	contactInfo.Address = cr.Address
	contactInfo.ConnectedAt = time.Unix(cr.ConnectedAt/1000, 0).Format("2006-01-02 03:04:05")

	websites := []Website{}
	for _, website := range cr.Websites {
		websites = append(websites, Website{
			URL:  website.URL,
			Type: website.Type.MetaData.Category,
		})
	}

	if len(websites) != 0 {
		contactInfo.Websites = websites
	}

	for _, twitter := range cr.TwitterHandles {
		contactInfo.Twitter = append(contactInfo.Twitter, twitter.Name)
	}

	phones := []Phone{}

	for _, phone := range cr.PhoneNumbers {
		phones = append(phones, Phone{
			Number: phone.Number,
			Type:   phone.Type,
		})
	}

	if len(phones) != 0 {
		contactInfo.PhoneNumber = phones
	}

	p.Contact = &contactInfo

	return nil
}
