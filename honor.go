package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type HonorNode struct {
	ProfileID  string  `json:"profileId,omitempty"`
	Paging     Paging  `json:"paging,omitempty"`
	RecipeType string  `json:"$recipeType,omitempty"`
	Elements   []Honor `json:"elements,omitempty"`

	err error
	ln  *Linkedin
}

type Honor struct {
	MultiLocaleTitle       *MultiLocale     `json:"multiLocaleTitle,omitempty"`
	Description            string           `json:"description,omitempty"`
	OccupationUnion        *OccupationUnion `json:"occupationUnion,omitempty"`
	Title                  string           `json:"title,omitempty"`
	Issuer                 string           `json:"issuer,omitempty"`
	IssuedOn               *IssuedOn        `json:"issuedOn,omitempty"`
	EntityUrn              string           `json:"entityUrn,omitempty"`
	RecipeType             string           `json:"$recipeType,omitempty"`
	MultiLocaleIssuer      *MultiLocale     `json:"multiLocaleIssuer,omitempty"`
	MultiLocaleDescription *MultiLocale     `json:"multiLocaleDescription,omitempty"`
	Occupation             string           `json:"occupation,omitempty"`
	IssueDate              *Date            `json:"issueDate,omitempty"`
}

type IssuedOn struct {
	Month int64 `json:"month,omitempty"`
	Year  int64 `json:"year,omitempty"`
}

type OccupationUnion struct {
	ProfilePosition string `json:"profilePosition,omitempty"`
}

func (hn *HonorNode) SetLinkedin(ln *Linkedin) {
	hn.ln = ln
}

func (hn *HonorNode) Next() bool {
	start := strconv.Itoa(hn.Paging.Start)
	count := strconv.Itoa(hn.Paging.Count)
	raw, err := hn.ln.get("/identity/profiles/"+hn.ProfileID+"/honors", url.Values{
		"start": {start},
		"count": {count},
	})

	if err != nil {
		hn.err = err
		return false
	}

	hnNode := new(HonorNode)
	if err := json.Unmarshal(raw, hnNode); err != nil {
		hn.err = err
		return false
	}

	hn.Elements = hnNode.Elements
	hn.Paging.Start = hnNode.Paging.Start + hnNode.Paging.Count

	if len(hn.Elements) == 0 {
		return false
	}

	return true
}

func (hn *HonorNode) Error() error {
	return hn.err
}
