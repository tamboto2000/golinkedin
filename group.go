package golinkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type GroupNode struct {
	Keywords string   `json:"keywords,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
	Elements []Group  `json:"elements,omitempty"`
	Paging   Paging   `json:"paging,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Group struct {
	// Elements contains groups from search group result
	Elements                                    []Group       `json:"elements,omitempty"`
	ExtendedElements                            []interface{} `json:"extendedElements,omitempty"`
	Image                                       *Image        `json:"image,omitempty"`
	TargetUrn                                   string        `json:"targetUrn,omitempty"`
	TrackingUrn                                 string        `json:"trackingUrn,omitempty"`
	Title                                       *Text         `json:"title,omitempty"`
	Type                                        string        `json:"type,omitempty"`
	Headline                                    *Text         `json:"headline,omitempty"`
	Subline                                     *Text         `json:"subline,omitempty"`
	TrackingID                                  string        `json:"trackingId,omitempty"`
	GroupName                                   string        `json:"groupName,omitempty"`
	EntityUrn                                   string        `json:"entityUrn,omitempty"`
	MemberCount                                 int           `json:"memberCount,omitempty"`
	Logo                                        *Image        `json:"logo,omitempty"`
	RecipeType                                  string        `json:"$recipeType,omitempty"`
	URL                                         string        `json:"url,omitempty"`
	DashEntityUrn                               string        `json:"dashEntityUrn"`
	DisplayNotificationSubscriptionLevelOptions bool          `json:"displayNotificationSubscriptionLevelOptions"`
	GroupPostNotificationsEdgeSettingUrn        string        `json:"groupPostNotificationsEdgeSettingUrn"`
	PostApprovalEnabled                         bool          `json:"postApprovalEnabled"`
	Description                                 *Text         `json:"description"`
	ShowPostApprovalOption                      bool          `json:"showPostApprovalOption"`
	Rules                                       string        `json:"rules"`
	Owners                                      []Attribute   `json:"owners"`
	LogoUrn                                     string        `json:"logoUrn"`
	InvitationLevel                             string        `json:"invitationLevel"`
	MemberConnectionsCount                      int           `json:"memberConnectionsCount"`
	CreatedAt                                   int           `json:"createdAt"`
	Name                                        *Text         `json:"name"`
	GlobalNewPostNotificationSettingOn          bool          `json:"globalNewPostNotificationSettingOn"`
	RecommendByAdminAvailable                   bool          `json:"recommendByAdminAvailable"`
	GroupUrn                                    string        `json:"groupUrn"`
	Managers                                    []Attribute   `json:"managers"`
}

func (ln *Linkedin) GroupByID(id int) (*Group, error) {
	raw, err := ln.get("/groups/groups/urn:li:group:"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	group := new(Group)
	if err := json.Unmarshal(raw, group); err != nil {
		return nil, err
	}

	return group, nil
}

func (gr *Group) GroupID() int {
	split := strings.Split(gr.EntityUrn, ":")
	id, _ := strconv.Atoi(split[len(split)-1])

	return id
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
