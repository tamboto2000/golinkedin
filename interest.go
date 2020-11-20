package golinkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type InterestNode struct {
	ProfileID string            `json:"profileId,omitempty"`
	Type      string            `json:"type,omitempty"`
	Elements  []InterestElement `json:"elements,omitempty"`
	Paging    Paging            `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type InterestElement struct {
	Entity        *Entity        `json:"entity,omitempty"`
	FollowingInfo *FollowingInfo `json:"followingInfo,omitempty"`
}

type Entity struct {
	MiniCompany *MiniCompany `json:"com.linkedin.voyager.entities.shared.MiniCompany,omitempty"`
	MiniGroup   *MiniGroup   `json:"com.linkedin.voyager.entities.shared.MiniGroup,omitempty"`
	MiniSchool  *MiniSchool  `json:"com.linkedin.voyager.entities.shared.MiniSchool,omitempty"`
	MiniProfile *MiniProfile `json:"com.linkedin.voyager.entities.shared.MiniProfile,omitempty"`
}

type FollowingInfo struct {
	FollowingType string `json:"followingType,omitempty"`
	EntityUrn     string `json:"entityUrn,omitempty"`
	FollowerCount int64  `json:"followerCount,omitempty"`
	Following     bool   `json:"following,omitempty"`
	TrackingUrn   string `json:"trackingUrn,omitempty"`
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
