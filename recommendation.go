package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type RecommendationNode struct {
	ProfileID string           `json:"profileId"`
	Q         string           `json:"q"`
	Elements  []Recommendation `json:"elements"`
	Paging    Paging           `json:"paging"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Recommendation struct {
	EntityUrn                      string       `json:"entityUrn"`
	Created                        int64        `json:"created"`
	Recommendee                    *MiniProfile `json:"recommendee"`
	RecommendationText             string       `json:"recommendationText"`
	RecommendeeEntity              string       `json:"recommendeeEntity"`
	VisibilityOnRecommenderProfile string       `json:"visibilityOnRecommenderProfile"`
	LastModified                   int64        `json:"lastModified"`
	Relationship                   string       `json:"relationship"`
	Recommender                    *MiniProfile `json:"recommender"`
	Status                         string       `json:"status"`
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
	rec.Paging = recNode.Paging

	if len(rec.Elements) < recNode.Paging.Count {
		rec.stopCursor = true
	}

	return true
}

func (rec *RecommendationNode) Error() error {
	return rec.err
}
