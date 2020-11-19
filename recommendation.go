package golinkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type RecommendationNode struct {
	ProfileID string           `json:"profileId,omitempty"`
	Q         string           `json:"q,omitempty"`
	Elements  []Recommendation `json:"elements,omitempty"`
	Paging    Paging           `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Recommendation struct {
	EntityUrn                      string       `json:"entityUrn,omitempty"`
	Created                        int64        `json:"created,omitempty"`
	Recommendee                    *MiniProfile `json:"recommendee,omitempty"`
	RecommendationText             string       `json:"recommendationText,omitempty"`
	RecommendeeEntity              string       `json:"recommendeeEntity,omitempty"`
	VisibilityOnRecommenderProfile string       `json:"visibilityOnRecommenderProfile,omitempty"`
	LastModified                   int64        `json:"lastModified,omitempty"`
	Relationship                   string       `json:"relationship,omitempty"`
	Recommender                    *MiniProfile `json:"recommender,omitempty"`
	Status                         string       `json:"status,omitempty"`
}

func (rec *RecommendationNode) SetLinkedin(ln *Linkedin) {
	rec.ln = ln
}

func (rec *RecommendationNode) Next() bool {
	if rec.stopCursor {
		return false
	}

	raw, err := rec.ln.get("/identity/profiles/"+rec.ProfileID+"/recommendations", url.Values{
		"q":     {rec.Q},
		"start": {strconv.Itoa(rec.Paging.Start)},
		"count": {strconv.Itoa(rec.Paging.Count)},
	})

	if err != nil {
		rec.err = err
		return false
	}

	recNode := new(RecommendationNode)
	if err := json.Unmarshal(raw, recNode); err != nil {
		rec.err = err
		return false
	}

	rec.Elements = recNode.Elements
	rec.Paging.Start = recNode.Paging.Start + recNode.Paging.Count

	if len(rec.Elements) < rec.Paging.Count {
		rec.stopCursor = true
	}

	return true
}

func (rec *RecommendationNode) Error() error {
	return rec.err
}
