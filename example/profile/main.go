package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/linkedin"
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

	// get contact info
	if err := contactInfo(profile); err != nil {
		panic(err.Error())
	}
}
