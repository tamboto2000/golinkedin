package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// SkillNode contains list of profile skills
type SkillNode struct {
	ProfileID  string  `json:"profileId,omitempty"`
	Paging     Paging  `json:"paging,omitempty"`
	Elements   []Skill `json:"elements,omitempty"`
	RecipeType string  `json:"$recipeType,omitempty"`

	err error
	ln  *Linkedin
}

type Skill struct {
	Name            string       `json:"name,omitempty"`
	MultiLocaleName *MultiLocale `json:"multiLocaleName,omitempty"`
	RecipeType      string       `json:"$recipeType,omitempty"`
	EntityUrn       string       `json:"entityUrn,omitempty"`
}

func (s *SkillNode) SetLinkedin(ln *Linkedin) {
	s.ln = ln
}

func (s *SkillNode) Next() bool {
	start := strconv.Itoa(s.Paging.Start)
	count := strconv.Itoa(s.Paging.Count)
	raw, err := s.ln.get("/identity/profiles/"+s.ProfileID+"/skills", url.Values{
		"start": {start},
		"count": {count},
	})

	if err != nil {
		s.err = err
		return false
	}

	skillNode := new(SkillNode)
	if err := json.Unmarshal(raw, skillNode); err != nil {
		s.err = err
		return false
	}

	s.Elements = skillNode.Elements
	s.Paging.Start = skillNode.Paging.Start + skillNode.Paging.Count

	if len(s.Elements) == 0 {
		return false
	}

	return true
}

func (s *SkillNode) Error() error {
	return s.err
}
