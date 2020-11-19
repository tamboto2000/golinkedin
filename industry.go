package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type IndustryNode struct {
	Metadata Metadata   `json:"metadata,omitempty"`
	Elements []Industry `json:"elements,omitempty"`
	Paging   Paging     `json:"paging,omitempty"`
	Keywords string     `json:"keywords,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Industry struct {
	LocalizedName       string `json:"localizedName"`
	Name                string `json:"name,omitempty"`
	RecipeType          string `json:"$recipeType,omitempty"`
	EntityUrn           string `json:"entityUrn,omitempty"`
	CompanyNameRequired bool   `json:"companyNameRequired,omitempty"`
	TargetUrn           string `json:"targetUrn,omitempty"`
	ObjectUrn           string `json:"objectUrn,omitempty"`
	Text                Text   `json:"text,omitempty"`
	DashTargetUrn       string `json:"dashTargetUrn,omitempty"`
	Type                string `json:"type,omitempty"`
	TrackingID          string `json:"trackingId,omitempty"`
}

func (ind *IndustryNode) SetLinkedin(ln *Linkedin) {
	ind.ln = ln
}

func (ind *IndustryNode) Next() bool {
	if ind.stopCursor {
		return false
	}

	start := strconv.Itoa(ind.Paging.Start)
	count := strconv.Itoa(ind.Paging.Count)
	raw, err := ind.ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {ind.Keywords},
		"origin":   {OriginOther},
		"q":        {QType},
		"type":     {TypeIndustry},
		"start":    {start},
		"count":    {count},
	})

	if err != nil {
		ind.err = err
		return false
	}

	indNode := new(IndustryNode)
	if err := json.Unmarshal(raw, indNode); err != nil {
		ind.err = err
		return false
	}

	ind.Elements = indNode.Elements
	ind.Paging.Start = indNode.Paging.Start + indNode.Paging.Count

	if len(ind.Elements) == 0 {
		return false
	}

	if len(ind.Elements) < ind.Paging.Count {
		ind.stopCursor = true
	}

	return true
}

func (ind *IndustryNode) Error() error {
	return ind.err
}
