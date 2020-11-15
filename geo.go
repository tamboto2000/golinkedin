package linkedin

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type GeoNode struct {
	Metadata     Metadata      `json:"metadata,omitempty"`
	Elements     []Geo         `json:"elements,omitempty"`
	Paging       Paging        `json:"paging,omitempty"`
	QueryContext *QueryContext `json:"queryContext,omitempty"`
	Keywords     string        `json:"keywords,omitempty"`

	err error
	ln  *Linkedin
}

type GeoLocation struct {
	Geo        Geo    `json:"geo,omitempty"`
	GeoUrn     string `json:"geoUrn,omitempty"`
	RecipeType string `json:"$recipeType,omitempty"`
}

type Geo struct {
	CountryUrn                             string   `json:"countryUrn,omitempty"`
	Country                                *Country `json:"country,omitempty"`
	DefaultLocalizedNameWithoutCountryName string   `json:"defaultLocalizedNameWithoutCountryName,omitempty"`
	EntityUrn                              string   `json:"entityUrn,omitempty"`
	RecipeType                             string   `json:"$recipeType,omitempty"`
	DefaultLocalizedName                   string   `json:"defaultLocalizedName,omitempty"`
	TargetUrn                              string   `json:"targetUrn,omitempty"`
	Text                                   Text     `json:"text,omitempty"`
	DashTargetUrn                          string   `json:"dashTargetUrn,omitempty"`
	Type                                   string   `json:"type,omitempty"`
	TrackingID                             string   `json:"trackingId,omitempty"`
}

type Text struct {
	Text string `json:"text,omitempty"`
}

type Metadata struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

func (g *GeoNode) SetLinkedin(ln *Linkedin) {
	g.ln = ln
}

func (g *GeoNode) Next() bool {
	start := strconv.Itoa(g.Paging.Start)
	count := strconv.Itoa(g.Paging.Count)
	raw, err := g.ln.get("/typeahead/hitsV2", url.Values{
		"keywords":     {g.Keywords},
		"origin":       {Other},
		"q":            {Type},
		"queryContext": {composeFilter(g.QueryContext)},
		"type":         {TGeo},
		"useCase":      {GeoAbbreviated},
		"start":        {start},
		"count":        {count},
	})

	if err != nil {
		g.err = err
		return false
	}

	geoNode := new(GeoNode)
	if err := json.Unmarshal(raw, geoNode); err != nil {
		g.err = err
		return false
	}

	g.Elements = geoNode.Elements
	g.Paging.Start = geoNode.Paging.Start + geoNode.Paging.Count

	if len(g.Elements) == 0 {
		return false
	}

	return true
}

func (g *GeoNode) Error() error {
	return g.err
}
