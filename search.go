package linkedin

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Connection Rank
const (
	Rank1 = "F"
	Rank2 = "S"
	Rank3 = "T"
)

// Result Type
const (
	RPeople = "PEOPLE"
)

// Search Origin
const (
	// OFacetedSearch could be used for people search
	OFacetedSearch = "FACETED_SEARCH"
	// OSwitchSearchVertical could be used for people search
	OSwitchSearchVertical = "SWITCH_SEARCH_VERTICAL"
	// OOtther could be used for geo search
	OOther = "OTHER"
)

// Contact Interest values
const (
	CIBoardMember = "boardMember"
	CiProBono     = "proBono"
)

// Geo Sub Type Filters
const (
	GMarketArea     = "MARKET_AREA"
	GCountryRegion  = "COUNTRY_REGION"
	GAdminDivision1 = "ADMIN_DIVISION_1"
	GCity           = "CITY"
)

// DefaultGeoQueryContext is default query context for geo search
var DefaultGeoQueryContext = &QueryContext{
	GeoVersion:            3,
	BingGeoSubTypeFilters: []string{GMarketArea, GCountryRegion, GAdminDivision1, GCity},
}

// Values of param useCase
const (
	GeoAbbreviated = "GEO_ABBREVIATED"
)

// Values of param type
const (
	TGeo         = "GEO"
	TCompany     = "COMPANY"
	TConnections = "CONNECTIONS"
	TIndustry    = "INDUSTRY"
	TSchool      = "SCHOOL"
)

// Values of param q, not to be confused with tag `q` on param struct or QueryContext
const (
	Type = "type"
	All  = "all"
)

// Languages
const (
	LangEnglish    = "en"
	LangIndonesian = "in"
	LangChinese    = "zh"
	LangOther      = "_o"
	LangJapanese   = "ja"
)

// PeopleSearchFilter is filter for people search.
// Query string representation:
//  List(
// 	 currentCompany->1344581|2135950|225166,
// 	 pastCompany->2145443|225166,
// 	 geoUrn->102478259|90010101|90010103,
// 	 industry->41|96,
//   profileLanguage->en|in|zh|_o|ja,
// 	 network->F|S,
// 	 profileLanguage->fr,
// 	 school->12953|456070,
// 	 connectionOf->ACoAABjYrYABlGzIXhNI0L2VSJH-hYQs_41qaQ8,
// 	 contactInterest->boardMember|proBono,
// 	 resultType->PEOPLE,
// 	 firstName->Franklin,
// 	 lastName->Tamboto,
// 	 title->Sr,
// 	 company->AAAA,
// 	 school->AAAA
//  )
// I don't know why its structured like that, maybe this is somekind of NoSQL?
type PeopleSearchFilter struct {
	CurrentCompany  []int    `q:"currentCompany" json:"currentCompany,omitempty"`
	PastCompany     []int    `q:"pastCompany" json:"pastCompany,omitempty"`
	GeoURN          []int    `q:"geoUrn" json:"geoUrn,omitempty"`
	Industry        []int    `q:"industry" json:"industry,omitempty"`
	Network         []string `q:"network" json:"network,omitempty"`
	ProfileLanguage []string `q:"profileLanguage" json:"profileLanguage,omitempty"`
	School          []int    `q:"school" json:"school,omitempty"`
	// Profile ID
	ConnectionOf    string   `q:"connectionOf" json:"connectionOf,omitempty"`
	ContactInterest []string `q:"contactInterest" json:"contactInterest,omitempty"`
	ResultType      string   `q:"resultType" json:"resultType,omitempty"`
	FirstName       string   `q:"firstName" json:"firstName,omitempty"`
	LastName        string   `q:"lastName" json:"lastName,omitempty"`
	Title           string   `q:"title" json:"title,omitempty"`
	Company         string   `q:"company" json:"company,omitempty"`
	// It will be still 'school', but with string value, i don't know why Linkedin have such ambiguous parameter.
	// So you will have 'school' param with array of int, and 'school' param with string
	SchoolStr string `q:"schoolStr" json:"schoolStr,omitempty"`
}

// QueryContext query string representation:
//  List(
// 	 geoVersion->3,
// 	 bingGeoSubTypeFilters->MARKET_AREA|COUNTRY_REGION|ADMIN_DIVISION_1|CITY,
//   spellCorrectionEnabled->true,
//   relatedSearchesEnabled->true
//  )
type QueryContext struct {
	SpellCorrectionEnabled bool     `q:"spellCorrectionEnabled" json:"spellCorrectionEnabled,omitempty"`
	RelatedSearchesEnabled bool     `q:"relatedSearchesEnabled" json:"relatedSearchesEnabled,omitempty"`
	GeoVersion             int      `q:"geoVersion" json:"geoVersion,omitempty"`
	BingGeoSubTypeFilters  []string `q:"bingGeoSubTypeFilters" json:"bingGeoSubTypeFilters,omitempty"`
}

// ComposeFilter parse param or query context to, well, whatever format Linkedin using...
func ComposeFilter(obj interface{}) string {
	return composeFilter(obj)
}

type filter struct {
	keyVal map[string][]string
}

func newFilter() filter {
	return filter{keyVal: make(map[string][]string)}
}

func (f filter) add(key string, val []string) {
	f.keyVal[key] = val
}

func (f filter) str() string {
	keyVal := make([]string, 0)
	for key, val := range f.keyVal {
		if key == "schoolStr" {
			key = "school"
		}

		valStr := strings.Join(val, "|")
		keyVal = append(keyVal, key+"->"+valStr)
	}

	return "List(" + strings.Join(keyVal, ",") + ")"
}

func composeFilter(obj interface{}) string {
	val := reflect.ValueOf(obj)
	filter := newFilter()
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	ty := val.Type()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)

		vals := make([]string, 0)
		if f.IsValid() && !f.IsZero() {
			if f.Type().String() == "[]int" {
				intVals := f.Interface().([]int)
				if len(intVals) > 0 {
					for _, intVal := range intVals {
						intStr := strconv.Itoa(intVal)
						vals = append(vals, intStr)
					}
				}
			}

			if f.Type().String() == "[]string" {
				strVals := f.Interface().([]string)
				if len(strVals) > 0 {
					for _, strVal := range strVals {
						vals = append(vals, strVal)
					}
				}
			}

			if f.Type().String() == "string" {
				str := f.Interface().(string)
				if str != "" {
					vals = append(vals, str)
				}
			}

			if f.Type().String() == "int" {
				str := strconv.Itoa(f.Interface().(int))
				// I assume that there is no entity with id 0
				if str != "0" {
					vals = append(vals, str)
				}
			}

			if f.Type().String() == "bool" {
				val := f.Interface().(bool)
				var str string
				if val {
					str = "true"
				} else {
					str = "false"
				}

				vals = append(vals, str)
			}

			if len(vals) > 0 {
				paramName := ty.Field(i).Tag.Get("q")
				filter.add(paramName, vals)
			}
		}
	}

	return filter.str()
}

// SearchGeo search geolocation by keywords. You can add custom QueryContext
func (ln *Linkedin) SearchGeo(keywords string, qctx *QueryContext) (*GeoNode, error) {
	if qctx == nil {
		qctx = DefaultGeoQueryContext
	}
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords":     {keywords},
		"origin":       {OOther},
		"q":            {Type},
		"queryContext": {composeFilter(qctx)},
		"type":         {TGeo},
		"useCase":      {GeoAbbreviated},
	})

	if err != nil {
		return nil, err
	}

	geoNode := new(GeoNode)
	if err := json.Unmarshal(raw, geoNode); err != nil {
		return nil, err
	}

	geoNode.ln = ln
	geoNode.QueryContext = qctx
	geoNode.Keywords = keywords

	return geoNode, nil
}

// SearchCompany search companies by keywords
func (ln *Linkedin) SearchCompany(keywords string) (*CompanyNode, error) {
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {keywords},
		"origin":   {OOther},
		"q":        {Type},
		"type":     {TCompany},
	})

	if err != nil {
		return nil, err
	}

	compNode := new(CompanyNode)
	if err := json.Unmarshal(raw, compNode); err != nil {
		return nil, err
	}

	compNode.ln = ln
	compNode.Keywords = keywords

	return compNode, nil
}

// SearchPeople search people by keywords.
// It's similiar to Profile, but with simpler and compact data.
// This API is actually for people search people filter, on section "Connections of".
// I call this as SearchPeople because it is similiar to Profile, but with simpler and compact data.
// It's kinda ambiguous if I call it SearchConnections.
func (ln *Linkedin) SearchPeople(keywords string) (*PeopleNode, error) {
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {keywords},
		"origin":   {OOther},
		"q":        {Type},
		"type":     {TConnections},
	})

	if err != nil {
		return nil, err
	}

	peopleNode := new(PeopleNode)
	if err := json.Unmarshal(raw, peopleNode); err != nil {
		return nil, err
	}

	peopleNode.ln = ln
	peopleNode.Keywords = keywords

	return peopleNode, nil
}

// SearchIndustry search industries by keywords
func (ln *Linkedin) SearchIndustry(keywords string) (*IndustryNode, error) {
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {keywords},
		"origin":   {OOther},
		"q":        {Type},
		"type":     {TIndustry},
	})

	if err != nil {
		return nil, err
	}

	indNode := new(IndustryNode)
	if err := json.Unmarshal(raw, indNode); err != nil {
		return nil, err
	}

	indNode.ln = ln
	indNode.Keywords = keywords

	return indNode, nil
}

// SearchSchool search school by keywords
func (ln *Linkedin) SearchSchool(keywords string) (*SchoolNode, error) {
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {keywords},
		"origin":   {OOther},
		"q":        {Type},
		"type":     {TSchool},
	})

	if err != nil {
		return nil, err
	}

	schoolNode := new(SchoolNode)
	if err := json.Unmarshal(raw, schoolNode); err != nil {
		return nil, err
	}

	schoolNode.ln = ln
	schoolNode.Keywords = keywords

	return schoolNode, nil
}
