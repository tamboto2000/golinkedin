package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type SchoolNode struct {
	Keywords string              `json:"keywords,omitempty"`
	Metadata Metadata            `json:"metadata,omitempty"`
	Elements []SchoolNodeElement `json:"elements,omitempty"`
	Paging   Paging              `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type SchoolNodeElement struct {
	ExtendedElements []interface{} `json:"extendedElements,omitempty"`
	Elements         []School      `json:"elements,omitempty"`
	Type             string        `json:"type,omitempty"`
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
