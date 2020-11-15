package main

import (
	"github.com/tamboto2000/linkedin"
)

func searchGeo(ln *linkedin.Linkedin, keyword string) (*linkedin.GeoNode, error) {
	geoNode, err := ln.SearchGeo(keyword, linkedin.DefaultGeoQueryContext)
	if err != nil {
		panic(err.Error())
	}

	geos := make([]linkedin.Geo, 0)
	for geoNode.Next() {
		geos = append(geos, geoNode.Elements...)
		if len(geos) >= 20 {
			break
		}
	}

	geoNode.Elements = geos

	return geoNode, nil
}
