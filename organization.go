package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// OrganizationNode contains user organizations info
type OrganizationNode struct {
	ProfileID  string         `json:"profileId,omitempty"`
	Paging     Paging         `json:"paging,omitempty"`
	RecipeType string         `json:"$recipeType,omitempty"`
	Elements   []Organization `json:"elements,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Organization struct {
	EntityUrn               string      `json:"entityUrn,omitempty"`
	DateRange               DateRange   `json:"dateRange,omitempty"`
	Name                    string      `json:"name,omitempty"`
	MultiLocaleName         MultiLocale `json:"multiLocaleName,omitempty"`
	RecipeType              string      `json:"$recipeType,omitempty"`
	PositionHeld            string      `json:"positionHeld,omitempty"`
	MultiLocalePositionHeld MultiLocale `json:"multiLocalePositionHeld,omitempty"`
	TimePeriod              *TimePeriod `json:"timePeriod,omitempty"`
	Position                string      `json:"position,omitempty"`
}

func (org *OrganizationNode) SetLinkedin(ln *Linkedin) {
	org.ln = ln
}

func (org *OrganizationNode) Next() bool {
	if org.stopCursor {
		return false
	}

	start := strconv.Itoa(org.Paging.Start)
	count := strconv.Itoa(org.Paging.Count)
	raw, err := org.ln.get("/identity/profiles/"+org.ProfileID+"/organizations", url.Values{
		"start": {start},
		"count": {count},
	})

	if err != nil {
		org.err = err
		return false
	}

	orgNode := new(OrganizationNode)
	if err := json.Unmarshal(raw, orgNode); err != nil {
		org.err = err
		return false
	}

	org.Elements = orgNode.Elements
	org.Paging.Start = orgNode.Paging.Start + orgNode.Paging.Count

	if len(org.Elements) == 0 {
		return false
	}

	if len(org.Elements) < org.Paging.Count {
		org.stopCursor = true
	}

	return true
}

func (org *OrganizationNode) Error() error {
	return org.err
}
