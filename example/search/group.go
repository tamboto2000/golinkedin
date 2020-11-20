package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin"
)

func searchGroup(ln *golinkedin.Linkedin, keywords string) error {
	grNode, err := ln.SearchGroup(keywords)
	if err != nil {
		return err
	}

	schs := make([]golinkedin.Group, 0)
	for grNode.Next() {
		schs = append(schs, grNode.Elements[0].Elements...)
		if len(schs) >= 20 {
			break
		}
	}

	grNode.Elements[0].Elements = schs
	f, err := os.Create("groups.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(grNode)
}
