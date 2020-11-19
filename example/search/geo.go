package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin/v1"
)

func searchGeo(ln *golinkedin.Linkedin, keyword string) error {
	geoNode, err := ln.SearchGeo(keyword)
	if err != nil {
		return err
	}

	geos := make([]golinkedin.Geo, 0)
	for geoNode.Next() {
		geos = append(geos, geoNode.Elements...)
		if len(geos) >= 20 {
			break
		}
	}

	geoNode.Elements = geos
	f, err := os.Create("geos.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(geoNode)
}
