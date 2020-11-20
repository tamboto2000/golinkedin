package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin"
)

func searchCompany(ln *golinkedin.Linkedin, keyword string) error {
	compNode, err := ln.SearchCompany(keyword)
	if err != nil {
		return err
	}

	comps := make([]golinkedin.Company, 0)
	for compNode.Next() {
		comps = append(comps, compNode.Elements[0].Elements...)
		if len(comps) >= 20 {
			break
		}
	}

	compNode.Elements[0].Elements = comps
	f, err := os.Create("companies.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(compNode)
}
