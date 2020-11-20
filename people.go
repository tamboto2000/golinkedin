package golinkedin

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type PeopleNode struct {
	Metadata     Metadata            `json:"metadata,omitempty"`
	Elements     []PeopleNodeElement `json:"elements,omitempty"`
	Paging       Paging              `json:"paging,omitempty"`
	Keywords     string              `json:"keywords,omitempty"`
	Filters      *PeopleSearchFilter `json:"peopleSearchFilter,omitempty"`
	QueryContext *QueryContext       `json:"queryContext,omitempty"`

	err        error
	ln         *Linkedin
	stopCursor bool
}

type PeopleNodeElement struct {
	ExtendedElements []ExtendedElement `json:"extendedElements,omitempty"`
	Elements         []People          `json:"elements,omitempty"`
	Type             string            `json:"type,omitempty"`
}

type ExtendedElement struct {
	RelatedSearches []RelatedSearch `json:"relatedSearches,omitempty"`
	Type            string          `json:"type,omitempty"`
}

type RelatedSearch struct {
	Query      *Query `json:"query,omitempty"`
	TrackingID string `json:"trackingId,omitempty"`
}

type Query struct {
	Attributes []QueryAttribute `json:"attributes,omitempty"`
	Text       string           `json:"text,omitempty"`
}

type QueryAttribute struct {
	Start  int64      `json:"start,omitempty"`
	Length int64      `json:"length,omitempty"`
	Type   *TypeClass `json:"type,omitempty"`
}

type TypeClass struct {
	COMLinkedinPemberlyTextBold COMLinkedinPemberlyTextBold `json:"com.linkedin.pemberly.text.Bold,omitempty"`
}

type COMLinkedinPemberlyTextBold struct {
}

type People struct {
	Image                Image           `json:"image,omitempty"`
	Subtext              Text            `json:"subtext,omitempty"`
	TargetUrn            string          `json:"targetUrn,omitempty"`
	ObjectUrn            string          `json:"objectUrn,omitempty"`
	Text                 Text            `json:"text,omitempty"`
	DashTargetUrn        string          `json:"dashTargetUrn,omitempty"`
	Type                 string          `json:"type,omitempty"`
	TrackingID           string          `json:"trackingId,omitempty"`
	MemberDistance       *MemberDistance `json:"memberDistance,omitempty"`
	SocialProofImagePile []Image         `json:"socialProofImagePile,omitempty"`
	TrackingUrn          string          `json:"trackingUrn,omitempty"`
	NavigationURL        string          `json:"navigationUrl,omitempty"`
	Title                *Title          `json:"title,omitempty"`
	Headless             bool            `json:"headless,omitempty"`
	Badges               *Badges         `json:"badges,omitempty"`
	SocialProofText      string          `json:"socialProofText,omitempty"`
	SnippetText          *SnippetText    `json:"snippetText,omitempty"`
	SecondaryTitle       *Title          `json:"secondaryTitle,omitempty"`
	PublicIdentifier     string          `json:"publicIdentifier,omitempty"`
	Headline             *Title          `json:"headline,omitempty"`
	NameMatch            bool            `json:"nameMatch,omitempty"`
	Subline              *Title          `json:"subline,omitempty"`
	Insights             []Insight       `json:"insights,omitempty"`
}

type Insight struct {
	Type        string      `json:"type,omitempty"`
	InsightText InsightText `json:"insightText,omitempty"`
}

type InsightText struct {
	TextDirection string                 `json:"textDirection,omitempty"`
	Attributes    []InsightTextAttribute `json:"attributes,omitempty"`
	Text          string                 `json:"text,omitempty"`
}

type SnippetText struct {
	Attributes []InsightTextAttribute `json:"attributes,omitempty"`
	Text       string                 `json:"text,omitempty"`
}

type InsightTextAttribute struct {
	Start  int64  `json:"start,omitempty"`
	Length int64  `json:"length,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Badges struct {
	Premium    bool   `json:"premium,omitempty"`
	Influencer bool   `json:"influencer,omitempty"`
	OpenLink   bool   `json:"openLink,omitempty"`
	EntityUrn  string `json:"entityUrn,omitempty"`
	JobSeeker  bool   `json:"jobSeeker,omitempty"`
}

type MemberDistance struct {
	Value string `json:"value,omitempty"`
}

type MiniProfile struct {
	FirstName        string           `json:"firstName,omitempty"`
	LastName         string           `json:"lastName,omitempty"`
	Occupation       string           `json:"occupation,omitempty"`
	ObjectUrn        string           `json:"objectUrn,omitempty"`
	EntityUrn        string           `json:"entityUrn,omitempty"`
	PublicIdentifier string           `json:"publicIdentifier,omitempty"`
	Picture          *Picture         `json:"picture,omitempty"`
	TrackingID       string           `json:"trackingId,omitempty"`
	BackgroundImage  *BackgroundImage `json:"backgroundImage,omitempty"`
}

type BackgroundImage struct {
	COMLinkedinCommonVectorImage VectorImage `json:"com.linkedin.common.VectorImage,omitempty"`
}

func (ppl *PeopleNode) SetLinkedin(ln *Linkedin) {
	ppl.ln = ln
}

func (ppl *PeopleNode) Next() bool {
	if ppl.stopCursor {
		return false
	}

	start := strconv.Itoa(ppl.Paging.Start)
	count := strconv.Itoa(ppl.Paging.Count)
	raw, err := ppl.ln.get("/search/blended", url.Values{
		"keywords":     {ppl.Keywords},
		"origin":       {OriginFacetedSearch},
		"q":            {QAll},
		"start":        {start},
		"count":        {count},
		"filters":      {composeFilter(ppl.Filters)},
		"queryContext": {composeFilter(ppl.QueryContext)},
	})

	if err != nil {
		ppl.err = err
		return false
	}

	pplNode := new(PeopleNode)
	if err := json.Unmarshal(raw, pplNode); err != nil {
		ppl.err = err
		return false
	}

	ppl.Elements = pplNode.Elements
	ppl.Paging.Start = pplNode.Paging.Start + pplNode.Paging.Count

	if len(ppl.Elements) == 0 {
		return false
	}

	if len(ppl.Elements[0].Elements) < ppl.Paging.Count {
		ppl.stopCursor = true
	}

	fmt.Println("cursor")

	return true
}

func (ppl *PeopleNode) Error() error {
	return ppl.err
}
