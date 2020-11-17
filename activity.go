package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ActivityNode struct {
	Metadata   *Metadata  `json:"metadata,omitempty"`
	Elements   []Activity `json:"elements,omitempty"`
	Paging     Paging     `json:"paging,omitempty"`
	Type       string     `json:"type,omitempty"`
	ProfileUrn string     `json:"profileUrn,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type Activity struct {
	Image              *BackgroundImage `json:"image,omitempty"`
	Featured           bool             `json:"featured,omitempty"`
	PermaLink          string           `json:"permaLink,omitempty"`
	LinkedinArticleUrn string           `json:"linkedinArticleUrn,omitempty"`
	ContentText        *Text            `json:"contentText,omitempty"`
	Title              string           `json:"title,omitempty"`
	NumLikes           int64            `json:"numLikes,omitempty"`
	PostedDate         *Date            `json:"postedDate,omitempty"`
	SocialDetail       *SocialDetail    `json:"socialDetail,omitempty"`
	TrackingData       *TrackingData    `json:"trackingData,omitempty"`
	NumComments        int64            `json:"numComments,omitempty"`
	CreatedDate        *Date            `json:"createdDate,omitempty"`
	PostedAt           int64            `json:"postedAt,omitempty"`
	EntityUrn          string           `json:"entityUrn,omitempty"`
	AuthorComponent    *AuthorComponent `json:"authorComponent,omitempty"`
	Actor              Actor            `json:"actor,omitempty"`
	UpdateMetadata     Metadata         `json:"updateMetadata,omitempty"`
	Content            Content          `json:"content,omitempty"`
	Commentary         Commentary       `json:"commentary,omitempty"`
}

type Commentary struct {
	TemplateType string      `json:"templateType,omitempty"`
	Text         Description `json:"text,omitempty"`
}

type Content struct {
	COMLinkedinVoyagerFeedRenderArticleComponent COMLinkedinVoyagerFeedRenderArticleComponent `json:"com.linkedin.voyager.feed.render.ArticleComponent,omitempty"`
}

type COMLinkedinVoyagerFeedRenderArticleComponent struct {
	TemplateType            string                                                    `json:"templateType,omitempty"`
	Urn                     string                                                    `json:"urn,omitempty"`
	LargeImage              *Image                                                    `json:"largeImage,omitempty"`
	Subtitle                *Description                                              `json:"subtitle,omitempty"`
	NavigationContext       *NavigationContext                                        `json:"navigationContext,omitempty"`
	Type                    string                                                    `json:"type,omitempty"`
	Title                   *Description                                              `json:"title,omitempty"`
	AuthorNavigationContext *NavigationContext                                        `json:"authorNavigationContext,omitempty"`
	Author                  *Description                                              `json:"author,omitempty"`
	Description             *Description                                              `json:"description,omitempty"`
	FollowAction            *COMLinkedinVoyagerFeedRenderArticleComponentFollowAction `json:"followAction,omitempty"`
	SubtitleImage           *Image                                                    `json:"subtitleImage,omitempty"`
	SubscribeAction         *SubscribeAction                                          `json:"subscribeAction,omitempty"`
}

type SubscribeAction struct {
	EntityUrn       string `json:"entityUrn,omitempty"`
	Subscribed      bool   `json:"subscribed,omitempty"`
	SubscriberCount int64  `json:"subscriberCount,omitempty"`
}

type COMLinkedinVoyagerFeedRenderArticleComponentFollowAction struct {
	FollowTrackingActionType   string         `json:"followTrackingActionType,omitempty"`
	FollowingInfo              *FollowingInfo `json:"followingInfo,omitempty"`
	UnfollowTrackingActionType string         `json:"unfollowTrackingActionType,omitempty"`
	Type                       string         `json:"type,omitempty"`
	TrackingActionType         string         `json:"trackingActionType,omitempty"`
}

type NavigationContext struct {
	TrackingActionType string `json:"trackingActionType,omitempty"`
	AccessibilityText  string `json:"accessibilityText,omitempty"`
	ActionTarget       string `json:"actionTarget,omitempty"`
}

type Actor struct {
	Urn                    string             `json:"urn,omitempty"`
	Image                  *Image             `json:"image,omitempty"`
	SupplementaryActorInfo *SubDescription    `json:"supplementaryActorInfo,omitempty"`
	Name                   *Name              `json:"name,omitempty"`
	SubDescription         *SubDescription    `json:"subDescription,omitempty"`
	NavigationContext      *NavigationContext `json:"navigationContext,omitempty"`
	Description            *Description       `json:"description,omitempty"`
	ShowInfluencerBadge    bool               `json:"showInfluencerBadge,omitempty"`
	FollowAction           *ActorFollowAction `json:"followAction,omitempty"`
}

type ActorFollowAction struct {
	FollowTrackingActionType string         `json:"followTrackingActionType,omitempty"`
	FollowingInfo            *FollowingInfo `json:"followingInfo,omitempty"`
	Type                     string         `json:"type,omitempty"`
	TrackingActionType       string         `json:"trackingActionType,omitempty"`
}

type Name struct {
	TextDirection string      `json:"textDirection,omitempty"`
	Attributes    []Attribute `json:"attributes,omitempty"`
	Text          string      `json:"text,omitempty"`
}

type SubDescription struct {
	TextDirection     string      `json:"textDirection,omitempty"`
	Attributes        []Attribute `json:"attributes,omitempty"`
	Text              string      `json:"text,omitempty"`
	AccessibilityText string      `json:"accessibilityText,omitempty"`
}

type AuthorComponent struct {
	Name        *Description          `json:"name,omitempty"`
	Image       *AuthorComponentImage `json:"image,omitempty"`
	Description *Description          `json:"description,omitempty"`
}

type AuthorComponentImage struct {
	Attributes                  []Attribute   `json:"attributes,omitempty"`
	AccessibilityTextAttributes []interface{} `json:"accessibilityTextAttributes,omitempty"`
	AccessibilityText           string        `json:"accessibilityText,omitempty"`
}

type Description struct {
	TextDirection string        `json:"textDirection,omitempty"`
	Attributes    []interface{} `json:"attributes,omitempty"`
	Text          string        `json:"text,omitempty"`
}

type TrackingData struct {
	TrackingID string `json:"trackingId,omitempty"`
}

type SocialDetail struct {
	ReactionElements          []interface{}              `json:"reactionElements,omitempty"`
	Comments                  *Comments                  `json:"comments,omitempty"`
	SocialPermissions         *SocialPermissions         `json:"socialPermissions,omitempty"`
	Liked                     bool                       `json:"liked,omitempty"`
	ShowShareButton           bool                       `json:"showShareButton,omitempty"`
	TotalShares               int64                      `json:"totalShares,omitempty"`
	Urn                       string                     `json:"urn,omitempty"`
	ThreadID                  string                     `json:"threadId,omitempty"`
	AllowedCommentersScope    string                     `json:"allowedCommentersScope,omitempty"`
	TotalSocialActivityCounts *TotalSocialActivityCounts `json:"totalSocialActivityCounts,omitempty"`
	EntityUrn                 string                     `json:"entityUrn,omitempty"`
	CommentingDisabled        bool                       `json:"commentingDisabled,omitempty"`
	SocialUpdateType          string                     `json:"socialUpdateType,omitempty"`
	Likes                     *ActivityNode              `json:"likes,omitempty"`
}

type Comments struct {
	Metadata *Metadata     `json:"metadata,omitempty"`
	Paging   Paging        `json:"paging,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
}

type SocialPermissions struct {
	CanPostComments bool   `json:"canPostComments,omitempty"`
	EntityUrn       string `json:"entityUrn,omitempty"`
	CanShare        bool   `json:"canShare,omitempty"`
	CanReact        bool   `json:"canReact,omitempty"`
}

type TotalSocialActivityCounts struct {
	SocialDetailEntityUrn string              `json:"socialDetailEntityUrn,omitempty"`
	Urn                   string              `json:"urn,omitempty"`
	NumComments           int64               `json:"numComments,omitempty"`
	ReactionTypeCounts    []ReactionTypeCount `json:"reactionTypeCounts,omitempty"`
	EntityUrn             string              `json:"entityUrn,omitempty"`
	NumShares             int64               `json:"numShares,omitempty"`
	NumLikes              int64               `json:"numLikes,omitempty"`
	Liked                 bool                `json:"liked,omitempty"`
}

type ReactionTypeCount struct {
	Count        int64  `json:"count,omitempty"`
	ReactionType string `json:"reactionType,omitempty"`
}

func (act *ActivityNode) SetLinkedin(ln *Linkedin) {
	act.ln = ln
}

func (act *ActivityNode) Next() bool {
	if act.stopCursor {
		return false
	}

	var raw []byte
	var err error

	if act.Type == ActivityArticle {
		raw, err = act.ln.get("/identity/profiles/williamhgates/posts", url.Values{
			"start": {strconv.Itoa(act.Paging.Start)},
			"count": {strconv.Itoa(act.Paging.Count)},
		})
	}

	if act.Type == ActivityPost {
		raw, err = act.ln.get("/identity/profileUpdatesV2", url.Values{
			"includeLongTermHistory": {"true"},
			"moduleKey":              {"member-shares:phone"},
			"numComments":            {"0"},
			"numLikes":               {"0"},
			"profileUrn":             {act.ProfileUrn},
			"q":                      {"memberShareFeed"},
			"start":                  {strconv.Itoa(act.Paging.Start)},
			"count":                  {strconv.Itoa(act.Paging.Count)},
		})
	}

	if err != nil {
		act.err = err
		return false
	}

	actNode := new(ActivityNode)
	if err := json.Unmarshal(raw, actNode); err != nil {
		act.err = err
		return false
	}

	act.Elements = actNode.Elements
	act.Paging.Start = actNode.Paging.Start + actNode.Paging.Count

	if len(act.Elements) < actNode.Paging.Count {
		act.stopCursor = true
	}

	return true
}

func (act *ActivityNode) Error() error {
	return act.err
}
