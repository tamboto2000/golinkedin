package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type InterestNode struct {
	ProfileID string            `json:"profileId"`
	Type      string            `json:"type"`
	Elements  []InterestElement `json:"elements"`
	Paging    Paging            `json:"paging"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type InterestElement struct {
	Entity        *Entity        `json:"entity"`
	FollowingInfo *FollowingInfo `json:"followingInfo"`
}

type Entity struct {
	MiniCompany *MiniCompany `json:"com.linkedin.voyager.entities.shared.MiniCompany"`
	MiniGroup   *MiniGroup   `json:"com.linkedin.voyager.entities.shared.MiniGroup"`
	MiniSchool  *MiniSchool  `json:"com.linkedin.voyager.entities.shared.MiniSchool"`
}

type FollowingInfo struct {
	FollowingType string `json:"followingType"`
	EntityUrn     string `json:"entityUrn"`
	FollowerCount int64  `json:"followerCount"`
	Following     bool   `json:"following"`
	TrackingUrn   string `json:"trackingUrn"`
}

func (inter *InterestNode) SetLinkedin(ln *Linkedin) {
	inter.ln = ln
}

func (inter *InterestNode) Next() bool {
	if inter.stopCursor {
		return false
	}

	raw, err := inter.ln.get("/identity/profiles/"+inter.ProfileID+"/following", url.Values{
		"entityType": {inter.Type},
		"q":          {"followedEntities"},
		"start":      {strconv.Itoa(inter.Paging.Start)},
		"count":      {strconv.Itoa(inter.Paging.Count)},
	})

	if err != nil {
		inter.err = err
		return false
	}

	interNode := new(InterestNode)
	if err := json.Unmarshal(raw, interNode); err != nil {
		inter.err = err
		return false
	}

	inter.Elements = interNode.Elements
	inter.Paging.Start = interNode.Paging.Start + interNode.Paging.Count

	if len(inter.Elements) < inter.Paging.Count {
		inter.stopCursor = true
	}

	return true
}

func (inter *InterestNode) Error() error {
	return inter.err
}
