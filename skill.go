package golinkedin

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

	err        error
	ln         *Linkedin
	stopCursor bool
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
	if s.stopCursor {
		return false
	}

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

	if len(s.Elements) < s.Paging.Count {
		s.stopCursor = true
	}

	return true
}

func (s *SkillNode) Error() error {
	return s.err
}
