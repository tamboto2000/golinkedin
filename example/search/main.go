package main

import (
	"github.com/tamboto2000/golinkedin/v1"
)

func main() {
	ln := golinkedin.New()
	ln.SetCookieStr(`your_linkedin_cookies`)

	// search geos
	if err := searchGeo(ln, "America"); err != nil {
		panic(err.Error())
	}

	// search companies
	if err := searchCompany(ln, "Telco"); err != nil {
		panic(err.Error())
	}

	// search industry
	if err := searchIndustry(ln, "Bank"); err != nil {
		panic(err.Error())
	}

	// search school
	if err := searchSchool(ln, "University"); err != nil {
		panic(err.Error())
	}

	// search service
	if err := searchService(ln, "Manager"); err != nil {
		panic(err.Error())
	}

	// search groups
	if err := searchGroup(ln, "ig"); err != nil {
		panic(err.Error())
	}

	// search people
	if err := searchPeople(ln, "ikh"); err != nil {
		panic(err.Error())
	}
}
