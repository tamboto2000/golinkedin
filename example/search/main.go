package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/linkedin"
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
