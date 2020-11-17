package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/linkedin"
)

func givenRecommendation(profile *linkedin.ProfileNode) error {
	recomm, err := profile.GivenRecommendation()
	if err != nil {
		return err
	}

	f, err := os.Create("given_recommendation.json")
	if err != nil {
		return err
	}

	defer f.Close()
	return json.NewEncoder(f).Encode(recomm)
}

func receivedRecommendation(profile *linkedin.ProfileNode) error {
	recomm, err := profile.ReceivedRecommendation()
	if err != nil {
		return err
	}

	f, err := os.Create("received_recommendation.json")
	if err != nil {
		return err
	}

	defer f.Close()
	return json.NewEncoder(f).Encode(recomm)
}
