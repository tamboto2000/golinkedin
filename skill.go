package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/tamboto2000/linkedin/raw"
)

type Skill struct {
	ID               string    `json:"id" csv:"id"`
	Name             string    `json:"name,omitempty" csv:"name"`
	EndorsementCount int       `json:"endorsmentCount,omitempty" csv:"endorsment_count"`
	Endorsers        []Profile `json:"endorsers,omitempty" csv:"-"`
}

func (p *Profile) SyncSkill() error {
	return p.syncSkill()
}

func (p *Profile) syncSkill() error {
	urlParsed, _ := url.Parse("https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/skillCategory?includeHiddenEndorsers=true")
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

	skillsView := new(raw.EndorsedSkillView)
	if err = json.Unmarshal(body, skillsView); err != nil {
		return err
	}

	skills := make([]Skill, 0)
	for _, e := range skillsView.Elements {
		for _, esk := range e.EndorsedSkills {
			skill := Skill{
				Name:             esk.Skill.Name,
				EndorsementCount: int(esk.EndorsementCount),
			}

			rawSkillID := strings.Split(esk.EntityUrn, ",")
			skill.ID = strings.Replace(rawSkillID[len(rawSkillID)-1], ")", "", 1)

			for _, prof := range esk.Endorsements {
				rawProfile := prof.Endorser.MiniProfile
				profile := Profile{
					Username: rawProfile.PublicIdentifier,
					ID:       strings.Replace(rawProfile.EntityUrn, "urn:li:fs_miniProfile:", "", 1),
					Name:     rawProfile.FirstName + " " + rawProfile.LastName,
					HeadLine: rawProfile.Occupation,
				}

				var pictURL string
				rawPict := rawProfile.Picture.COMLinkedinCommonVectorImage
				if rawPict.Artifacts != nil && len(rawPict.Artifacts) > 0 {
					pictURL = rawPict.RootURL
					pictURL += rawPict.Artifacts[len(rawPict.Artifacts)-1].FileIdentifyingURLPathSegment
				}

				profile.ProfilePict = pictURL
				skill.Endorsers = append(skill.Endorsers, profile)
			}

			skills = append(skills, skill)
		}
	}

	p.Skill = skills

	return nil
}
