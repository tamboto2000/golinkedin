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

	// get contact info
	if err := contactInfo(profile); err != nil {
		panic(err.Error())
	}

	// get given recommendation
	if err := givenRecommendation(profile); err != nil {
		panic(err.Error())
	}

	// get received recommendation
	if err := receivedRecommendation(profile); err != nil {
		panic(err.Error())
	}

	// get activities
	if err := activity(profile); err != nil {
		panic(err.Error())
	}
}
