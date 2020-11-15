package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type PeopleNode struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Elements []People `json:"elements,omitempty"`
	Paging   Paging   `json:"paging,omitempty"`
	Keywords string   `json:"keywords,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type People struct {
	Image         Image  `json:"image,omitempty"`
	Subtext       Text   `json:"subtext,omitempty"`
	TargetUrn     string `json:"targetUrn,omitempty"`
	ObjectUrn     string `json:"objectUrn,omitempty"`
	Text          Text   `json:"text,omitempty"`
	DashTargetUrn string `json:"dashTargetUrn,omitempty"`
	Type          string `json:"type,omitempty"`
	TrackingID    string `json:"trackingId,omitempty"`
}

type MiniProfile struct {
	FirstName        string  `json:"firstName,omitempty"`
	LastName         string  `json:"lastName,omitempty"`
	Occupation       string  `json:"occupation,omitempty"`
	ObjectUrn        string  `json:"objectUrn,omitempty"`
	EntityUrn        string  `json:"entityUrn,omitempty"`
	PublicIdentifier string  `json:"publicIdentifier,omitempty"`
	Picture          Picture `json:"picture,omitempty"`
	TrackingID       string  `json:"trackingId,omitempty"`
}

func (p *PeopleNode) SetLinkedin(ln *Linkedin) {
	p.ln = ln
}

func (p *PeopleNode) Next() bool {
	if p.stopCursor {
		return false
	}

	start := strconv.Itoa(p.Paging.Start)
	count := strconv.Itoa(p.Paging.Count)
	raw, err := p.ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {p.Keywords},
		"origin":   {OOther},
		"q":        {Type},
		"type":     {TConnections},
		"start":    {start},
		"count":    {count},
	})

	if err != nil {
		p.err = err
		return false
	}

	peopleNode := new(PeopleNode)
	if err := json.Unmarshal(raw, peopleNode); err != nil {
		p.err = err
		return false
	}

	p.Elements = peopleNode.Elements
	p.Paging.Start = peopleNode.Paging.Start + peopleNode.Paging.Count

	if len(p.Elements) == 0 {
		return false
	}

	if len(p.Elements) < p.Paging.Count {
		p.stopCursor = true
	}

	return true
}

func (p *PeopleNode) Error() error {
	return p.err
}
