package golinkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type SchoolNode struct {
	Keywords string   `json:"keywords,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
	Elements []School `json:"elements,omitempty"`
	Paging   Paging   `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type School struct {
	// Element contains schools from search school result
	Elements         []School      `json:"elements,omitempty"`
	ExtendedElements []interface{} `json:"extendedElements,omitempty"`
	ObjectUrn        string        `json:"objectUrn,omitempty"`
	EntityUrn        string        `json:"entityUrn,omitempty"`
	Active           bool          `json:"active,omitempty"`
	Logo             *Logo         `json:"logo,omitempty"`
	SchoolName       string        `json:"schoolName,omitempty"`
	TrackingID       string        `json:"trackingId,omitempty"`
	Image            Image         `json:"image,omitempty"`
	Subtext          Text          `json:"subtext,omitempty"`
	TargetUrn        string        `json:"targetUrn,omitempty"`
	Text             Text          `json:"text,omitempty"`
	DashTargetUrn    string        `json:"dashTargetUrn,omitempty"`
	Type             string        `json:"type,omitempty"`
	TrackingUrn      string        `json:"trackingUrn,omitempty"`
	Title            Title         `json:"title,omitempty"`
}

// SchoolByName lookup school by universal name.
func (ln *Linkedin) SchoolByName(name string) (*CompanyNode, error) {
	raw, err := ln.get("/organization/companies", url.Values{
		"decorationId":  {"com.linkedin.voyager.deco.organization.web.WebFullCompanyMain-28"},
		"q":             {"universalName"},
		"universalName": {name},
	})

	if err != nil {
		return nil, err
	}

	schNode := new(CompanyNode)
	if err := json.Unmarshal(raw, schNode); err != nil {
		return nil, err
	}

	return schNode, nil
}

func (sch *SchoolNode) SetLinkedin(ln *Linkedin) {
	sch.ln = ln
}

func (sch *SchoolNode) Next() bool {
	if sch.stopCursor {
		return false
	}

	start := strconv.Itoa(sch.Paging.Start)
	count := strconv.Itoa(sch.Paging.Count)
	raw, err := sch.ln.get("/search/blended", url.Values{
		"keywords":     {sch.Keywords},
		"origin":       {OriginSwitchSearchVertical},
		"q":            {QAll},
		"start":        {start},
		"count":        {count},
		"filters":      {composeFilter(DefaultSearchSchoolFilter)},
		"queryContext": {composeFilter(DefaultSearchSchoolQueryContext)},
	})

	if err != nil {
		sch.err = err
		return false
	}

	schNode := new(SchoolNode)
	if err := json.Unmarshal(raw, schNode); err != nil {
		sch.err = err
		return false
	}

	sch.Elements = schNode.Elements
	sch.Paging.Start = schNode.Paging.Start + schNode.Paging.Count

	if len(sch.Elements) == 0 {
		return false
	}

	if len(sch.Elements[0].Elements) < sch.Paging.Count {
		sch.stopCursor = true
	}

	return true
}

func (sch *SchoolNode) Error() error {
	return sch.err
}
