package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type SchoolNode struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Elements []School `json:"elements,omitempty"`
	Paging   Paging   `json:"paging,omitempty"`
	Keywords string   `json:"keywords,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
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
	raw, err := sch.ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {sch.Keywords},
		"origin":   {OOther},
		"q":        {Type},
		"type":     {TSchool},
		"start":    {start},
		"count":    {count},
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

	if len(sch.Elements) < sch.Paging.Count {
		sch.stopCursor = true
	}

	return true
}

func (sch *SchoolNode) Error() error {
	return sch.err
}
