package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/tamboto2000/linkedin/raw"
)

type Recommendation struct {
	Recommender  *Profile `json:"recommender,omitempty"`
	Recommendee  *Profile `json:"recommendee,omitempty"`
	CreatedAt    string   `json:"createdAt"`
	Text         string   `json:"text"`
	LastModified string   `json:"lastModified"`
	Status       string   `json:"status"`
}

func (p *Profile) SyncRecommendation() error {
	tp := newThreadPool()
	tp.add(2)
	go tp.run(p.syncReceivedRecommendation)
	go tp.run(p.syncGivenRecommendation)
	tp.wait()

	return tp.getError()
}

func (p *Profile) syncReceivedRecommendation() error {
	urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/recommendations?q=received&recommendationStatuses=List(VISIBLE)?"
	urlStr += "start=0&count=50"
	urlParsed, _ := url.Parse("https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/recommendations?q=received&recommendationStatuses=List(VISIBLE)")
	resp, err := p.ln.client.getRequest(urlParsed)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		return errors.New("linkedin API error: " + string(body))
	}

	recommView := new(raw.RecommendationView)
	if err = json.Unmarshal(body, recommView); err != nil {
		return err
	}

	recomms := make([]Recommendation, 0)
	for _, data := range recommView.Elements {
		lastModified := time.Unix(data.LastModified/1000.0, 0).Format("2006-01-02 15:04:05")
		createdAt := time.Unix(data.Created/1000.0, 0).Format("2006-01-02 15:04:05")
		recomm := Recommendation{
			Text:         data.RecommendationText,
			CreatedAt:    createdAt,
			LastModified: lastModified,
			Status:       data.Status,
		}

		var profPictURL string
		if len(data.Recommender.Picture.ComLinkedinCommonVectorImage.Artifacts) > 0 {
			profPictURL = data.Recommender.Picture.ComLinkedinCommonVectorImage.RootURL
			profPictURL += data.Recommender.Picture.ComLinkedinCommonVectorImage.Artifacts[len(data.Recommender.Picture.ComLinkedinCommonVectorImage.Artifacts)-1].FileIdentifyingURLPathSegment
		}

		recomm.Recommender = &Profile{
			Username:    data.Recommender.PublicIdentifier,
			ID:          strings.Replace(data.Recommender.EntityUrn, "urn:li:fs_miniProfile:", "", 1),
			Occupation:  data.Recommender.Occupation,
			ProfilePict: profPictURL,
		}

		recomms = append(recomms, recomm)
	}

	p.ReceivedRecommendations = recomms

	return nil
}

func (p *Profile) syncGivenRecommendation() error {
	urlStr := "https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/recommendations?q=given"
	urlStr += "start=0&count=50"
	urlParsed, _ := url.Parse("https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/recommendations?q=received&recommendationStatuses=List(VISIBLE)")
	resp, err := p.ln.client.getRequest(urlParsed)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		return errors.New("linkedin API error: " + string(body))
	}

	recommView := new(raw.RecommendationView)
	if err = json.Unmarshal(body, recommView); err != nil {
		return err
	}

	recomms := make([]Recommendation, 0)
	for _, data := range recommView.Elements {
		lastModified := time.Unix(data.LastModified/1000.0, 0).Format("2006-01-02 15:04:05")
		createdAt := time.Unix(data.Created/1000.0, 0).Format("2006-01-02 15:04:05")
		recomm := Recommendation{
			Text:         data.RecommendationText,
			CreatedAt:    createdAt,
			LastModified: lastModified,
			Status:       data.Status,
		}

		var profPictURL string
		if len(data.Recommendee.Picture.ComLinkedinCommonVectorImage.Artifacts) > 0 {
			profPictURL = data.Recommendee.Picture.ComLinkedinCommonVectorImage.RootURL
			profPictURL += data.Recommendee.Picture.ComLinkedinCommonVectorImage.Artifacts[len(data.Recommendee.Picture.ComLinkedinCommonVectorImage.Artifacts)-1].FileIdentifyingURLPathSegment
		}

		recomm.Recommendee = &Profile{
			Username:    data.Recommendee.PublicIdentifier,
			ID:          strings.Replace(data.Recommendee.EntityUrn, "urn:li:fs_miniProfile:", "", 1),
			Occupation:  data.Recommendee.Occupation,
			ProfilePict: profPictURL,
		}

		recomms = append(recomms, recomm)
	}

	p.GivenRecommendations = recomms

	return nil
}
