package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin"
)

func searchSchool(ln *golinkedin.Linkedin, keywords string) error {
	schNode, err := ln.SearchSchool(keywords)
	if err != nil {
		return err
	}

	schs := make([]golinkedin.School, 0)
	for schNode.Next() {
		schs = append(schs, schNode.Elements[0].Elements...)
		if len(schs) >= 20 {
			break
		}
	}

	schNode.Elements[0].Elements = schs
	f, err := os.Create("schools.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(schNode)
}
