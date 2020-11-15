package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/linkedin"
)

func searchSchool(ln *linkedin.Linkedin, keywords string) error {
	schNode, err := ln.SearchSchool(keywords)
	if err != nil {
		return err
	}

	schs := make([]linkedin.School, 0)
	for schNode.Next() {
		schs = append(schs, schNode.Elements...)
		if len(schs) >= 20 {
			break
		}
	}

	schNode.Elements = schs
	f, err := os.Create("schools.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(schNode)
}
