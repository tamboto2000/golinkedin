package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin"
)

func searchPeople(ln *linkedin.Linkedin, keywords string) error {
	pplNode, err := ln.SearchPeople(keywords, nil)
	if err != nil {
		return err
	}

	schs := make([]linkedin.People, 0)
	for pplNode.Next() {
		schs = append(schs, pplNode.Elements[0].Elements...)
		if len(schs) >= 20 {
			break
		}
	}

	pplNode.Elements[0].Elements = schs
	f, err := os.Create("peoples.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(pplNode)
}
