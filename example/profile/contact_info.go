package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin/v1"
)

func contactInfo(profile *linkedin.ProfileNode) error {
	contact, err := profile.ContactInfo()
	if err != nil {
		return err
	}

	f, err := os.Create("contact_info.json")
	if err != nil {
		return err
	}

	defer f.Close()
	return json.NewEncoder(f).Encode(contact)
}
