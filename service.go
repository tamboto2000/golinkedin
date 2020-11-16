package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ServiceNode struct {
	Metadata Metadata  `json:"metadata,omitempty"`
	Elements []Service `json:"elements,omitempty"`
	Paging   Paging    `json:"paging,omitempty"`
	Keywords string    `json:"keywords,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Service struct {
	TargetUrn     string `json:"targetUrn,omitempty"`
	ObjectUrn     string `json:"objectUrn,omitempty"`
	Text          Text   `json:"text,omitempty"`
	DashTargetUrn string `json:"dashTargetUrn,omitempty"`
	Type          string `json:"type,omitempty"`
	TrackingID    string `json:"trackingId,omitempty"`
}

func (svc *ServiceNode) SetLinkedin(ln *Linkedin) {
	svc.ln = ln
}

func (svc *ServiceNode) Next() bool {
	if svc.stopCursor {
		return false
	}

	start := strconv.Itoa(svc.Paging.Start)
	count := strconv.Itoa(svc.Paging.Count)
	raw, err := svc.ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {svc.Keywords},
		"origin":   {OriginOther},
		"q":        {QType},
		"type":     {TypeSkill},
		"useCase":  {MarketplaceSkills},
		"start":    {start},
		"count":    {count},
	})

	if err != nil {
		svc.err = err
		return false
	}

	svcNode := new(ServiceNode)
	if err := json.Unmarshal(raw, svcNode); err != nil {
		svc.err = err
		return false
	}

	svc.Elements = svcNode.Elements
	svc.Paging.Start = svcNode.Paging.Start + svcNode.Paging.Count

	if len(svc.Elements) == 0 {
		return false
	}

	if len(svc.Elements) < svc.Paging.Count {
		svc.stopCursor = true
	}

	return true
}

func (svc *ServiceNode) Error() error {
	return svc.err
}
