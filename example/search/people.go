package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/linkedin"
)

func searchPeople(ln *linkedin.Linkedin, keywords string) error {
	peopleNode, err := ln.SearchPeople(keywords)
	if err != nil {
		return err
	}

	comps := make([]linkedin.People, 0)
	for peopleNode.Next() {
		comps = append(comps, peopleNode.Elements...)
		if len(comps) >= 20 {
			break
		}
	}

	peopleNode.Elements = comps
	f, err := os.Create("peoples.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(peopleNode)
}
