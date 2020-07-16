package linkedin

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"

	"github.com/tamboto2000/linkedin/raw"
)

type Company struct {
	ID            string `json:"id"`
	UniversalName string `json:"universalName"`
	Name          string `json:"name"`
	LogoURL       string `json:"logoURL"`
	Active        bool   `json:"active"`
	URL           string `json:"url"`
}

type Group struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LogoURL     string `json:"logoURL"`
}

type Interests struct {
	Influencers []Profile `json:"influencers"`
	Companies   []Company `json:"companies"`
	Groups      []Group   `json:"groups"`
	Schools     []School  `json:"schools"`
}

func (p *Profile) SyncInterests() error {
	p.Interests = new(Interests)
	tp := newThreadPool()
	tp.add(4)
	go tp.run(p.syncInterestInfluencer)
	go tp.run(p.syncInterestsCompany)
	go tp.run(p.syncInterestsSchool)
	go tp.run(p.syncInterestsGroup)
	tp.wait()

	return tp.getError()
}

func (p *Profile) syncInterestInfluencer() error {
	interests := make([]Profile, 0)
	count := 10
	start := 0
	hasNextPage := true

	for hasNextPage {
		urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/following?"
		urlStr += "count=" + strconv.Itoa(count) + "&"
		urlStr += "entityType=INFLUENCER&"
		urlStr += "q=followedEntities&"
		urlStr += "start=" + strconv.Itoa(start)
		urlParsed, _ := url.Parse(urlStr)

		resp, err := p.ln.client.getRequest(urlParsed)

		if err != nil {
			return err
		}

		defer resp.Body.Close()

		resp_body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		payload := new(raw.InterestInfluencerView)
		if err = json.Unmarshal(resp_body, payload); err != nil {
			return err
		}

		for _, e := range payload.Elements {
			rawProfile := e.Entity.ComLinkedinVoyagerIdentitySharedMiniProfile
			var profilePict string
			rawProfPict := rawProfile.Picture.ComLinkedinCommonVectorImage
			if rawProfPict.Artifacts != nil && len(rawProfPict.Artifacts) > 0 {
				profilePict = rawProfPict.RootURL
				profilePict += rawProfPict.Artifacts[len(rawProfPict.Artifacts)-1].FileIdentifyingURLPathSegment
			}

			interests = append(interests, Profile{
				ID:          strings.Replace(rawProfile.EntityUrn, "urn:li:fs_miniProfile:", "", 1),
				Username:    rawProfile.PublicIdentifier,
				Name:        rawProfile.FirstName + " " + rawProfile.LastName,
				Occupation:  rawProfile.Occupation,
				ProfilePict: profilePict,
			})
		}

		if len(payload.Elements) < count-1 {
			hasNextPage = false
		}

		start += count
	}

	p.Interests.Influencers = interests

	return nil
}

func (p *Profile) syncInterestsCompany() error {
	companies := make([]Company, 0)
	count := 10
	start := 0
	hasNextPage := true
	for hasNextPage {
		urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/following?"
		urlStr += "count=" + strconv.Itoa(count) + "&"
		urlStr += "entityType=COMPANY&"
		urlStr += "q=followedEntities&"
		urlStr += "start=" + strconv.Itoa(start)
		urlParsed, _ := url.Parse(urlStr)

		resp, err := p.ln.client.getRequest(urlParsed)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		resp_body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		payload := new(raw.InterestCompanyLoad)
		if err = json.Unmarshal(resp_body, payload); err != nil {
			return err
		}

		for _, e := range payload.Elements {
			var logoURL string
			logoRaw := e.Entity.ComLinkedinVoyagerEntitiesSharedMiniCompany.Logo.ComLinkedinCommonVectorImage
			if logoRaw.Artifacts != nil && len(logoRaw.Artifacts) > 0 {
				logoURL = logoRaw.RootURL
				logoURL += logoRaw.Artifacts[len(logoRaw.Artifacts)-1].FileIdentifyingURLPathSegment
			}

			rawCompany := e.Entity.ComLinkedinVoyagerEntitiesSharedMiniCompany
			var companyURL string
			if rawCompany.UniversalName != "" {
				companyURL = "https://www.linkedin.com/company/" + rawCompany.UniversalName
			}

			companies = append(companies, Company{
				UniversalName: rawCompany.UniversalName,
				Name:          rawCompany.Name,
				Active:        rawCompany.Active,
				LogoURL:       logoURL,
				URL:           companyURL,
			})
		}

		if len(payload.Elements) < count-1 {
			hasNextPage = false
		}

		start += count
	}

	p.Interests.Companies = companies

	return nil
}

func (p *Profile) syncInterestsSchool() error {
	schools := make([]School, 0)
	count := 10
	start := 0
	hasNextPage := true
	for hasNextPage {
		urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/following?"
		urlStr += "count=" + strconv.Itoa(count) + "&"
		urlStr += "entityType=SCHOOL&"
		urlStr += "q=followedEntities&"
		urlStr += "start=" + strconv.Itoa(start)
		urlParsed, _ := url.Parse(urlStr)

		resp, err := p.ln.client.getRequest(urlParsed)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		resp_body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		payload := new(raw.InterestSchoolLoad)
		if err = json.Unmarshal(resp_body, payload); err != nil {
			return err
		}

		for _, e := range payload.Elements {
			rawSchool := e.Entity.ComLinkedinVoyagerEntitiesSharedMiniSchool
			var logoURL string
			rawLogo := rawSchool.Logo.ComLinkedinCommonVectorImage
			if rawLogo.Artifacts != nil && len(rawLogo.Artifacts) > 0 {
				logoURL = rawLogo.RootURL
				logoURL += rawLogo.Artifacts[len(rawLogo.Artifacts)-1].FileIdentifyingURLPathSegment
			}

			schoolID := strings.Replace(rawSchool.EntityUrn, "urn:li:fs_miniSchool:", "", 1)
			schools = append(schools, School{
				URL:  "https://www.linkedin.com/school/" + schoolID + "/?legacySchoolId=" + schoolID,
				Logo: logoURL,
				Name: rawSchool.SchoolName,
			})
		}

		if len(payload.Elements) < count-1 {
			hasNextPage = false
		}

		start += count
	}

	p.Interests.Schools = schools

	return nil
}

func (p *Profile) syncInterestsGroup() error {
	groups := make([]Group, 0)
	count := 10
	start := 0
	hasNextPage := true
	for hasNextPage {
		urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/following?"
		urlStr += "count=" + strconv.Itoa(count) + "&"
		urlStr += "entityType=GROUP&"
		urlStr += "q=followedEntities&"
		urlStr += "start=" + strconv.Itoa(start)
		urlParsed, _ := url.Parse(urlStr)

		resp, err := p.ln.client.getRequest(urlParsed)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		resp_body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		payload := new(raw.InterestGroupLoad)
		if err = json.Unmarshal(resp_body, payload); err != nil {
			return err
		}

		for _, e := range payload.Elements {
			rawGroup := e.Entity.ComLinkedinVoyagerEntitiesSharedMiniGroup
			var logoURL string
			rawLogo := rawGroup.Logo.ComLinkedinCommonVectorImage
			if rawLogo.Artifacts != nil && len(rawLogo.Artifacts) > 0 {
				logoURL = rawLogo.RootURL
				logoURL += rawLogo.Artifacts[len(rawLogo.Artifacts)-1].FileIdentifyingURLPathSegment
			}

			groupID := strings.Replace(rawGroup.ObjectUrn, "urn:li:group:", "", 1)
			groupIDInt, _ := strconv.Atoi(groupID)
			groups = append(groups, Group{
				ID:          groupIDInt,
				URL:         "https://www.linkedin.com/groups/" + groupID,
				Name:        rawGroup.GroupName,
				Description: rawGroup.GroupDescription,
				LogoURL:     logoURL,
			})
		}

		if len(payload.Elements) < count-1 {
			hasNextPage = false
		}

		start += count
	}

	p.Interests.Groups = groups

	return nil
}
