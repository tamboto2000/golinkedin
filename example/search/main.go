package main

import (
	"github.com/tamboto2000/linkedin"
)

func main() {
	ln := linkedin.New()
	ln.SetCookieStr(`your_linkedin_cookies`)

	// search geos
	if err := searchGeo(ln, "America"); err != nil {
		panic(err.Error())
	}

	// search companies
	if err := searchCompany(ln, "Telco"); err != nil {
		panic(err.Error())
	}

	// search people
	if err := searchPeople(ln, "Ikhwan"); err != nil {
		panic(err.Error())
	}
}
