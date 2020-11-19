package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/golinkedin/v1"
)

func activity(profile *golinkedin.ProfileNode) error {
	act, err := profile.Activity(golinkedin.ActivityArticle)
	if err != nil {
		return err
	}

	acts := make([]golinkedin.Activity, 0)
	for act.Next() {
		acts = append(acts, act.Elements...)
		if len(acts) >= 20 {
			break
		}
	}

	act.Elements = acts
	f, err := os.Create("activities.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(act)
}
