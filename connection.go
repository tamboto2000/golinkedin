package linkedin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"

	"github.com/tamboto2000/linkedin/raw"
)

type Connection struct {
	Connection int `json:"connection,omitempty" csv:"connection"`
	Follower   int `json:"follower,omitempty" csv:"follower"`
}

func (p *Profile) SyncConnection() error {
	return p.syncConnection()
}

func (p *Profile) syncConnection() error {
	urlParsed, _ := url.Parse("https://www.linkedin.com/voyager/api/identity/profiles/" + p.Username + "/networkinfo")
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

	connectionView := new(raw.ConnectionView)
	if err = json.Unmarshal(body, connectionView); err != nil {
		return err
	}

	p.Connection = &Connection{
		Connection: int(connectionView.ConnectionsCount),
		Follower:   int(connectionView.FollowersCount),
	}

	return nil
}
