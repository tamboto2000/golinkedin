package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type GroupNode struct {
	Keywords string             `json:"keywords,omitempty"`
	Metadata Metadata           `json:"metadata,omitempty"`
	Elements []GroupNodeElement `json:"elements,omitempty"`
	Paging   Paging             `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type GroupNodeElement struct {
	ExtendedElements []interface{} `json:"extendedElements,omitempty"`
	Elements         []Group       `json:"elements,omitempty"`
	Type             string        `json:"type,omitempty"`
}

type Group struct {
	Image       Image    `json:"image,omitempty"`
	TargetUrn   string   `json:"targetUrn,omitempty"`
	TrackingUrn string   `json:"trackingUrn,omitempty"`
	Title       Headline `json:"title,omitempty"`
	Type        string   `json:"type,omitempty"`
	Headline    Headline `json:"headline,omitempty"`
	Subline     Headline `json:"subline,omitempty"`
	TrackingID  string   `json:"trackingId,omitempty"`
}

func (gr *GroupNode) SetLinkedin(ln *Linkedin) {
	gr.ln = ln
}

func (gr *GroupNode) Next() bool {
	if gr.stopCursor {
		return false
	}

	start := strconv.Itoa(gr.Paging.Start)
	count := strconv.Itoa(gr.Paging.Count)
	raw, err := gr.ln.get("/search/blended", url.Values{
		"keywords":     {gr.Keywords},
		"origin":       {OriginSwitchSearchVertical},
		"q":            {QAll},
		"filters":      {composeFilter(DefaultSearchGroupFilter)},
		"queryContext": {composeFilter(DefaultSearchGroupQueryContext)},
		"start":        {start},
		"count":        {count},
	})

	if err != nil {
		gr.err = err
		return false
	}

	grNode := new(GroupNode)
	if err := json.Unmarshal(raw, grNode); err != nil {
		gr.err = err
		return false
	}

	gr.Elements = grNode.Elements
	gr.Paging.Start = grNode.Paging.Start + grNode.Paging.Count

	if len(gr.Elements) == 0 {
		return false
	}

	if len(gr.Elements[0].Elements) < gr.Paging.Count {
		gr.stopCursor = true
	}

	return true
}

func (gr *GroupNode) Error() error {
	return gr.err
}
