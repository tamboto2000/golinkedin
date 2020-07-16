package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"

	"github.com/tamboto2000/linkedin/raw"
)

type Location struct {
	FullAddress        string `json:"fullAddress,omitempty" csv:"full_address"`
	LocationName       string `json:"addressName,omitempty" csv:"address_name"`
	CountryCode        string `json:"countryCode,omitempty" csv:"country_code"`
	GeoLocation        string `json:"geoLocation" csv:"geo_location"`
	GeoLocationCountry string `json:"geoLocationCountry" csv:"geo_location_country"`
}

type Profile struct {
	Username                string           `json:"username" csv:"username"`
	ID                      string           `json:"id" csv:"id"`
	Name                    string           `json:"name,omitempty" csv:"name"`
	ProfilePict             string           `json:"profilePict,omitempty" csv:"profile_pict"`
	BackgroundImg           string           `json:"backgroundImg,omitempty"`
	Summary                 string           `json:"summary,omitempty" csv:"summary"`
	HeadLine                string           `json:"headLine,omitempty" csv:"headline"`
	Occupation              string           `json:"occupation,omitempty"`
	School                  []School         `json:"school,omitempty" csv:"-"`
	Location                *Location        `json:"location,omitempty" csv:"-"`
	Connection              *Connection      `json:"connection,omitempty" csv:"-"`
	Job                     []Job            `json:"job,omitempty" csv:"-"`
	Skill                   []Skill          `json:"skill,omitempty" csv:"-"`
	Contact                 *ContactInfo     `json:"contact,omitempty" csv:"-"`
	ReceivedRecommendations []Recommendation `json:"receivedRecommendations,omitempty"`
	GivenRecommendations    []Recommendation `json:"givenRecommendations,omitempty"`
	Interests               *Interests       `json:"interests,omitempty"`
	Certifications          []Certification  `json:"certifications,omitempty"`

	ln *Linkedin
}

func (ln *Linkedin) NewProfile(username string) (*Profile, error) {
	profile, err := reqProfileView(username, ln.client)
	if err != nil {
		return nil, err
	}

	profile.ln = ln
	return profile, nil
}

func reqProfileView(username string, cl client) (*Profile, error) {
	urlP, _ := url.Parse("https://www.linkedin.com/voyager/api/identity/profiles/" + username + "/profileView")
	resp, err := cl.getRequest(urlP)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		if resp.StatusCode == 404 {
			return nil, errors.New("user not found")
		}

		return nil, errors.New(string(body))
	}

	viewProfile := new(raw.ViewProfile)
	if err = json.Unmarshal(body, viewProfile); err != nil {
		return nil, err
	}

	return composeProfile(viewProfile), nil
}

func composeProfile(viewLoad *raw.ViewProfile) *Profile {
	profile := new(Profile)
	profile.Username = viewLoad.Profile.MiniProfile.PublicIdentifier
	profile.ID = viewLoad.PositionView.ProfileID
	profile.Name = viewLoad.Profile.MiniProfile.FirstName + " " + viewLoad.Profile.MiniProfile.LastName

	profPictRaw := viewLoad.Profile.MiniProfile.Picture.COMLinkedinCommonVectorImage
	if profPictRaw.Artifacts != nil && len(profPictRaw.Artifacts) > 0 {
		profPict := profPictRaw.RootURL
		profPict += profPictRaw.Artifacts[len(profPictRaw.Artifacts)-1].FileIdentifyingURLPathSegment
	}

	profile.Summary = viewLoad.Profile.Summary
	profile.HeadLine = viewLoad.Profile.Headline
	profile.Occupation = viewLoad.Profile.MiniProfile.Occupation
	profile.Location = &Location{
		FullAddress:        viewLoad.Profile.GeoLocationName + ", " + viewLoad.Profile.GeoCountryName,
		GeoLocation:        viewLoad.Profile.GeoLocationName + ", " + viewLoad.Profile.GeoCountryName,
		GeoLocationCountry: viewLoad.Profile.GeoCountryName,
		LocationName:       viewLoad.Profile.GeoLocationName,
		CountryCode:        viewLoad.Profile.Location.BasicLocation.CountryCode,
	}

	miniProfile := viewLoad.Profile.MiniProfile
	vector := miniProfile.Picture.COMLinkedinCommonVectorImage
	if vector.Artifacts != nil && len(vector.Artifacts) > 0 {
		profPict := vector.RootURL
		profPict += vector.Artifacts[len(vector.Artifacts)-1].FileIdentifyingURLPathSegment
		profile.ProfilePict = profPict
	}

	vector = miniProfile.BackgroundImage.COMLinkedinCommonVectorImage
	if vector.Artifacts != nil && len(vector.Artifacts) > 0 {
		backPict := vector.RootURL
		backPict += vector.Artifacts[len(vector.Artifacts)-1].FileIdentifyingURLPathSegment
		profile.BackgroundImg = backPict
	}

	return profile
}
