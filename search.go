package golinkedin

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
	Rank3 = "O"
)

// Result Type
const (
	ResultPeople    = "PEOPLE"
	ResultSchools   = "SCHOOLS"
	ResultCompanies = "COMPANIES"
	ResultGroups    = "GROUPS"
)

// Search Origin
const (
	// OriginFacetedSearch could be used for people search
	OriginFacetedSearch = "FACETED_SEARCH"
	// OriginSwitchSearchVertical could be used for people search
	OriginSwitchSearchVertical = "SWITCH_SEARCH_VERTICAL"
	// OriginOtther could be used for geo search
	OriginOther = "OTHER"
	// OriginMemberProfileCannedSearch can be used for people search
	OriginMemberProfileCannedSearch = "MEMBER_PROFILE_CANNED_SEARCH"
)

// Contact Interest values
const (
	ContactInterestBoardMember = "boardMember"
	ContactInterestProBono     = "proBono"
)

// Geo Sub Type Filters
const (
	GeoSubTypeMarketArea     = "MARKET_AREA"
	GeoSubTypeCountryRegion  = "COUNTRY_REGION"
	GeoSubTypeAdminDivision1 = "ADMIN_DIVISION_1"
	GeoSubTypeCity           = "CITY"
)

// Values of param useCase
const (
	GeoAbbreviated    = "GEO_ABBREVIATED"
	MarketplaceSkills = "MARKETPLACE_SKILLS"
)

// Values of param type
const (
	TypeGeo         = "GEO"
	TypeCompany     = "COMPANY"
	TypeConnections = "CONNECTIONS"
	TypePeople      = "PEOPLE"
	TypeIndustry    = "INDUSTRY"
	TypeSchool      = "SCHOOL"
	TypeSkill       = "SKILL"
)

// Values of param q, not to be confused with tag `q` on param struct or QueryContext
const (
	QType = "type"
	QAll  = "all"
)

// Languages
const (
	LangEnglish    = "en"
	LangIndonesian = "in"
	LangChinese    = "zh"
	LangJapanese   = "ja"
	LangGerman     = "de"
	LangFrench     = "fr"
	LangSpanish    = "es"
	LangPortuguese = "pt"
	LangOther      = "_o"
)

// DefaultGeoSearchQueryContext is default query context for geo search
var DefaultGeoSearchQueryContext = &QueryContext{
	GeoVersion: 3,
	BingGeoSubTypeFilters: []string{
		GeoSubTypeMarketArea,
		GeoSubTypeCountryRegion,
		GeoSubTypeAdminDivision1,
		GeoSubTypeCity,
	},
}

// DefaultSearchSchoolFilter used for search school filters param
var DefaultSearchSchoolFilter = &Filters{
	ResultType: ResultSchools,
}

// DefaultSearchSchoolQueryContext used for search school queryContext param
var DefaultSearchSchoolQueryContext = &QueryContext{
	SpellCorrectionEnabled: true,
}

// DefaultSearchCompanyFilter used for search companies filters param
var DefaultSearchCompanyFilter = &Filters{
	ResultType: ResultCompanies,
}

// DefaultSearchCompanyQueryContext used for search companies queryContext param
var DefaultSearchCompanyQueryContext = &QueryContext{
	FlagshipSearchIntent: "SEARCH_SRP",
}

// DefaultSearchGroupFilter used for search groups filters param
var DefaultSearchGroupFilter = &Filters{
	ResultType: ResultGroups,
}

// DefaultSearchGroupQueryContext used for search groups queryContext param
var DefaultSearchGroupQueryContext = &QueryContext{
	SpellCorrectionEnabled: true,
}

// DefaultSearchPeopleFilter used for search peoples filters param
var DefaultSearchPeopleFilter = &PeopleSearchFilter{
	ResultType: ResultPeople,
}

// DefaultSearchPeopleQueryContext used for search peoples queryContext param
var DefaultSearchPeopleQueryContext = &QueryContext{
	SpellCorrectionEnabled: true,
	RelatedSearchesEnabled: true,
}

// PeopleSearchFilter is filter for people search.
type PeopleSearchFilter struct {
	CurrentCompany  []int    `q:"currentCompany" json:"currentCompany,omitempty"`
	PastCompany     []int    `q:"pastCompany" json:"pastCompany,omitempty"`
	GeoURN          []int    `q:"geoUrn" json:"geoUrn,omitempty"`
	Industry        []int    `q:"industry" json:"industry,omitempty"`
	Network         []string `q:"network" json:"network,omitempty"`
	ProfileLanguage []string `q:"profileLanguage" json:"profileLanguage,omitempty"`
	School          []int    `q:"school" json:"school,omitempty"`
	ServiceCategory []string `q:"serviceCategory" json:"serviceCategory,omitempty"`
	// Profile ID
	ConnectionOf    string   `q:"connectionOf" json:"connectionOf,omitempty"`
	ContactInterest []string `q:"contactInterest" json:"contactInterest,omitempty"`
	FirstName       string   `q:"firstName" json:"firstName,omitempty"`
	LastName        string   `q:"lastName" json:"lastName,omitempty"`
	Title           string   `q:"title" json:"title,omitempty"`
	Company         string   `q:"company" json:"company,omitempty"`
	// It will be still 'school', but with string value, i don't know why Linkedin have such ambiguous parameter.
	// So you will have 'school' param with array of int, and 'school' param with string
	SchoolStr string `q:"schoolStr" json:"schoolStr,omitempty"`

	// will automaticaly set
	ResultType string `q:"resultType" json:"resultType,omitempty"`
}

// Filters is generic filter used by no filter search, such as school search or company search
type Filters struct {
	ResultType string `q:"resultType" json:"resultType,omitempty"`
}

type QueryContext struct {
	SpellCorrectionEnabled bool     `q:"spellCorrectionEnabled" json:"spellCorrectionEnabled,omitempty"`
	RelatedSearchesEnabled bool     `q:"relatedSearchesEnabled" json:"relatedSearchesEnabled,omitempty"`
	FlagshipSearchIntent   string   `q:"flagshipSearchIntent" json:"flagshipSearchIntent,omitempty"`
	GeoVersion             int      `q:"geoVersion" json:"geoVersion,omitempty"`
	BingGeoSubTypeFilters  []string `q:"bingGeoSubTypeFilters" json:"bingGeoSubTypeFilters,omitempty"`
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

// SearchPeople search people based on filter.
// If filter is nil, default value will be used, and so with ctx and origin
func (ln *Linkedin) SearchPeople(keywords string, filter *PeopleSearchFilter, ctx *QueryContext, origin string) (*PeopleNode, error) {
	if filter == nil {
		filter = DefaultSearchPeopleFilter
	}

	if ctx == nil {
		ctx = DefaultSearchPeopleQueryContext
	}

	if origin == "" {
		origin = OriginFacetedSearch
	}

	filter.ResultType = ResultPeople

	raw, err := ln.get("/search/blended", url.Values{
		"keywords":     {keywords},
		"origin":       {origin},
		"q":            {QAll},
		"filters":      {composeFilter(filter)},
		"queryContext": {composeFilter(ctx)},
	})

	if err != nil {
		return nil, err
	}

	peopleNode := new(PeopleNode)
	if err := json.Unmarshal(raw, peopleNode); err != nil {
		return nil, err
	}

	peopleNode.Keywords = keywords
	peopleNode.ln = ln
	peopleNode.Filters = filter
	peopleNode.QueryContext = DefaultSearchPeopleQueryContext
	peopleNode.Paging.Start += peopleNode.Paging.Count

	return peopleNode, nil
}

// SearchGeo search geolocation by keywords
func (ln *Linkedin) SearchGeo(keywords string) (*GeoNode, error) {
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords":     {keywords},
		"origin":       {OriginOther},
		"q":            {QType},
		"queryContext": {composeFilter(DefaultGeoSearchQueryContext)},
		"type":         {TypeGeo},
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
	geoNode.Keywords = keywords
	geoNode.Paging.Start += geoNode.Paging.Count

	return geoNode, nil
}

// SearchCompany search companies by keywords
func (ln *Linkedin) SearchCompany(keywords string) (*CompanyNode, error) {
	raw, err := ln.get("/search/blended", url.Values{
		"keywords":     {keywords},
		"origin":       {OriginSwitchSearchVertical},
		"q":            {QAll},
		"filters":      {composeFilter(DefaultSearchCompanyFilter)},
		"queryContext": {composeFilter(DefaultSearchCompanyQueryContext)},
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
	compNode.Paging.Start += compNode.Paging.Count

	return compNode, nil
}

// SearchIndustry search industries by keywords
func (ln *Linkedin) SearchIndustry(keywords string) (*IndustryNode, error) {
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {keywords},
		"origin":   {OriginOther},
		"q":        {QType},
		"type":     {TypeIndustry},
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
	indNode.Paging.Start += indNode.Paging.Count

	return indNode, nil
}

// SearchSchool search school by keywords
func (ln *Linkedin) SearchSchool(keywords string) (*SchoolNode, error) {
	raw, err := ln.get("/search/blended", url.Values{
		"keywords":     {keywords},
		"origin":       {OriginSwitchSearchVertical},
		"q":            {QAll},
		"filters":      {composeFilter(DefaultSearchSchoolFilter)},
		"queryContext": {composeFilter(DefaultSearchSchoolQueryContext)},
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
	schoolNode.Paging.Start += schoolNode.Paging.Count

	return schoolNode, nil
}

// SearchService search service by keywords
func (ln *Linkedin) SearchService(keywords string) (*ServiceNode, error) {
	raw, err := ln.get("/typeahead/hitsV2", url.Values{
		"keywords": {keywords},
		"origin":   {OriginOther},
		"q":        {QType},
		"type":     {TypeSkill},
		"useCase":  {MarketplaceSkills},
	})

	if err != nil {
		return nil, err
	}

	svcNode := new(ServiceNode)
	if err := json.Unmarshal(raw, svcNode); err != nil {
		return nil, err
	}

	svcNode.ln = ln
	svcNode.Keywords = keywords
	svcNode.Paging.Start += svcNode.Paging.Count

	return svcNode, nil
}

// SearchGroup search groups by keywords
func (ln *Linkedin) SearchGroup(keywords string) (*GroupNode, error) {
	raw, err := ln.get("/search/blended", url.Values{
		"keywords":     {keywords},
		"origin":       {OriginSwitchSearchVertical},
		"q":            {QAll},
		"filters":      {composeFilter(DefaultSearchGroupFilter)},
		"queryContext": {composeFilter(DefaultSearchGroupQueryContext)},
	})

	if err != nil {
		return nil, err
	}

	groupNode := new(GroupNode)
	if err := json.Unmarshal(raw, groupNode); err != nil {
		return nil, err
	}

	groupNode.ln = ln
	groupNode.Keywords = keywords
	groupNode.Paging.Start += groupNode.Paging.Count

	return groupNode, nil
}
