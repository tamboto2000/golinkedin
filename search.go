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
	People = "PEOPLE"
)

// Search Origin
const (
	// FacetedSearch could be used for people search
	FacetedSearch = "FACETED_SEARCH"
	// SwitchSearchVertical could be used for people search
	SwitchSearchVertical = "SWITCH_SEARCH_VERTICAL"
	// otther could be used for geo search
	Other = "OTHER"
)

// Contact Interest values
const (
	BoardMember = "boardMember"
	ProBono     = "proBono"
)

// Geo Sub Type Filters
const (
	MarketArea     = "MARKET_AREA"
	CountryRegion  = "COUNTRY_REGION"
	AdminDivision1 = "ADMIN_DIVISION_1"
	City           = "CITY"
)

// DefaultGeoQueryContext is default query context for geo search
var DefaultGeoQueryContext = &QueryContext{
	GeoVersion:            3,
	BingGeoSubTypeFilters: []string{MarketArea, CountryRegion, AdminDivision1, City},
}

// Values of param useCase
const (
	GeoAbbreviated = "GEO_ABBREVIATED"
)

// Values of param type
const (
	// I name it TGeo because there is a struct named Geo. God help me...
	TGeo = "GEO"
)

// Values of param q, not to be confused with tag `q` on param struct or QueryContext
const (
	Type = "type"
	All  = "all"
)

// PeopleSearchFilter is filter for people search.
// Query string representation:
//  List(
// 	 currentCompany->1344581|2135950|225166,
// 	 pastCompany->2145443|225166,
// 	 geoUrn->102478259|90010101|90010103,
// 	 industry->41|96,
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

// SearchGeo search geolocation by keyword. You can add custom QueryContext
func (ln *Linkedin) SearchGeo(keyword string, qctx *QueryContext) (*GeoNode, error) {
	if qctx == nil {
		qctx = DefaultGeoQueryContext
	}
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords":     {keyword},
		"origin":       {Other},
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
	geoNode.Keywords = keyword

	return geoNode, nil
}

// SearchCompany search companies by keyword
func (ln *Linkedin) SearchCompany(keyword string) (*CompanyNode, error) {
	return nil, nil
}
