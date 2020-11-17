# Linkedin
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tamboto2000/golinkedin)](https://pkg.go.dev/github.com/tamboto2000/golinkedin) [![GitHub](https://img.shields.io/github/license/tamboto2000/linkedin)](https://github.com/tamboto2000/golinkedin/blob/v1/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tamboto2000/linkedin) [![GitHub tag (latest by tag)](https://img.shields.io/badge/tag-v1.1.0-informational)](https://github.com/tamboto2000/golinkedin/tree/v1.1.0) [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/tamboto2000/linkedin/v1-build/v1)](https://github.com/tamboto2000/golinkedin/actions/runs/368476376)

Linkedin is a library for scraping Linkedin.
Unfortunately, auto login is impossible (probably...), so you need to retrieve Linkedin session cookies manually.
As mentioned above, the purpose of this package is only for scraping, so there is no method for create, update, or delete data.
Not all object is documented or present because Franklin Collin Tamboto, the original author, does not fully understand the purpose of some object returned by Linkedin internal API, and because the nature of Linkedin internal API that treat almost every object as optional, empty field or object will not be returned by Linkedin internal API, so some object or fields might be missing.
Feel free to fork and contribute!

# Current Features

  - Lookup Full Profile by Username
  - Profile Organizations Lookup
  - Profile Educations Lookup
  - Profile Certifications Lookup
  - Profile Honors Lookup
  - Profile Positions Lookup
  - Profile Interest Lookup
  - Profile Acitivity Lookup
  - Profile Recommendation Lookup
  - Profile Skill Lookup
  - Geolocation Search  
  - Company Search
  - Group Search
  - People/Profile Search
  - Service Search
  - School Search

# Upcoming Features

  - Profile Treasury Media Data
  - Profile Publications Lookup
  - Profile Volunteer Exeperiences Lookup
  - Profile Projects Lookup
  - Profile Patents Lookup
  - Profile Languages Data
  - Profile Courses Lookup

### Installation

This package require go version 1.14 or above.
Make sure you have go modules activated.
```sh
$ GO111MODULE=on go get github.com/tamboto2000/golinkedin/v1
```

### Full Profile Lookup Example
```go
package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin/v1"
)

func main() {
	ln := linkedin.New()
	ln.SetCookieStr(`your_linkedin_cookies`)

	profile, err := ln.ProfileByUsername("linkedin_username")
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create("profile.json")
	if err != nil {
		panic(err.Error())
	}

	if err := json.NewEncoder(f).Encode(profile); err != nil {
		panic(err.Error())
	}
}

```

### Search Geo Example
```go
package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin/v1"
)

func main() {
	ln := linkedin.New()
	ln.SetCookieStr(`your_linkedin_cookies`)

	// search geo
	geoNode, err := searchGeo(ln, "USA")
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create("geo.json")
	if err != nil {
		panic(err.Error())
	}

	if err := json.NewEncoder(f).Encode(geoNode); err != nil {
		panic(err.Error())
	}
}

func searchGeo(ln *linkedin.Linkedin, keyword string) (*linkedin.GeoNode, error) {
	geoNode, err := ln.SearchGeo(keyword)
	if err != nil {
		panic(err.Error())
	}

	geos := make([]linkedin.Geo, 0)
	for geoNode.Next() {
		geos = append(geos, geoNode.Elements...)
		if len(geos) >= 20 {
			break
		}
	}

	geoNode.Elements = geos

	return geoNode, nil
}
```

For now, every Node have SetLinkedin(), Error() error, and Next() bool method, except for ProfileNode

### Todos

 - Write Tests
 - Add More Data and Features
 - Add CodeCove

License
----

MIT
