package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin"
)

func searchIndustry(ln *golinkedin.Linkedin, keywords string) error {
	indNode, err := ln.SearchIndustry(keywords)
	if err != nil {
		return err
	}

	inds := make([]golinkedin.Industry, 0)
	for indNode.Next() {
		inds = append(inds, indNode.Elements...)
		if len(inds) >= 20 {
			break
		}
	}

	indNode.Elements = inds
	f, err := os.Create("industries.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(indNode)
}
