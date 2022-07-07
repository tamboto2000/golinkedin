# v2 is on the way!
At the beginning, this library is created because of my previous job which involve scraping social media platforms, the intended purpose of this library is to get the Linkedin data as much as possible but still follow the raw structure that Linkedin has, which is to be frank, a ball of tangled mess of JSON. Now that a lot of people starting to use this library too I think it is time for me to be a bit more serious to maintain this. So with that in mind, I announce that there will be V2 of this library, which will introduce much simpler data structure and will be much more easy to use, and of course, much more clearer documentation. If anybody wants to help me, feel free to do so. Thanks!

\- Franklin

# Golinkedin
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tamboto2000/golinkedin)](https://pkg.go.dev/github.com/tamboto2000/golinkedin) [![GitHub](https://img.shields.io/github/license/tamboto2000/golinkedin)](https://github.com/tamboto2000/golinkedin/blob/v1/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tamboto2000/golinkedin) [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/tamboto2000/linkedin/v1-build/v1)](https://github.com/tamboto2000/golinkedin/runs/1433927028)

Golinkedin is a library for scraping Linkedin.
Unfortunately, auto login is impossible (probably...), so you need to retrieve Linkedin session cookies manually.
As mentioned above, the purpose of this package is only for scraping, so there is no method for create, update, or delete data.
Not all object is documented or present because the original author does not fully understand the purpose of some object returned by Linkedin internal API, and because the nature of Linkedin internal API that treat almost every object as optional, empty field or object will not be returned by Linkedin internal API, so some object or fields might be missing.
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
  - Profile Accomplishments

### Installation

This package require go version 1.14 or above.
Make sure you have go modules activated.
```sh
$ GO111MODULE=on go get github.com/tamboto2000/golinkedin
```

### Full Profile Lookup Example
```go
package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin"
)

func main() {
	ln := golinkedin.New()
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

	"github.com/tamboto2000/golinkedin"
)

func main() {
	ln := golinkedin.New()
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

func searchGeo(ln *golinkedin.Linkedin, keyword string) (*golinkedin.GeoNode, error) {
	geoNode, err := ln.SearchGeo(keyword)
	if err != nil {
		panic(err.Error())
	}

	geos := make([]golinkedin.Geo, 0)
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
