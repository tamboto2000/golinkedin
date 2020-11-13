// Package linkedin is a library for scraping Linkedin.
// Unfortunately, auto login is impossible (probably...), so you need to retrieve Linkedin session cookies manually.
// As mentioned above, the purpose of this package is only for scraping, so there is no method for create, update, or delete data.
// Not all object is documented or present because Franklin Collin Tamboto, the original author, does not fully understand the purpose
// of some object returned by Linkedin internal API, and because the nature of Linkedin internal API that treat almost every object as
// optional, empty field or object will not be presented by Linkedin internal API.
// Feel free to fork and contribute!
package linkedin

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// ErrSessionInvalid is returned when Linkedin session can not be used
	// for requesting Linkedin API
	ErrSessionInvalid = "Linkedin session invalid"

	basePath  = "https://www.linkedin.com"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:82.0) Gecko/20100101 Firefox/82.0"
)

var apiBase string = basePath + "/voyager/api"

// Linkedin hold all (covered) features and data resources
type Linkedin struct {
	client  *http.Client
	cookies []*http.Cookie
}

// New initiate new Linkedin object
func New() *Linkedin {
	ln := &Linkedin{client: new(http.Client)}
	return ln
}

// SetCookieStr set Linkedin session cookie string
func (ln *Linkedin) SetCookieStr(c string) {
	header := http.Header{}
	header.Add("Cookie", c)
	request := http.Request{Header: header}

	ln.SetCookies(request.Cookies())
}

// SetCookies set Linkedin session cookies from parsed cookie string
func (ln *Linkedin) SetCookies(c []*http.Cookie) {
	ln.cookies = c
}

// SetProxy set proxy to client
func (ln *Linkedin) SetProxy(p string) error {
	uri, err := url.Parse(p)
	if err != nil {
		return err
	}

	ln.client.Transport = &http.Transport{Proxy: http.ProxyURL(uri), TLSHandshakeTimeout: 20 * time.Second}
	return nil
}

func (ln *Linkedin) get(path string, q url.Values) ([]byte, error) {
	uri, err := url.Parse(apiBase + path)
	if err != nil {
		return nil, err
	}

	if q != nil {
		uri.RawQuery = q.Encode()
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Connection", "keep-alive")
	// req.Header.Add("Cookie", c.cookie)
	req.Header.Add("Host", "www.linkedin.com")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("csrf-token", ln.jsessionid())

	for _, cookie := range ln.cookies {
		req.AddCookie(cookie)
	}

	resp, err := ln.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 200 {
		return nil, errors.New(string(raw))
	}

	return raw, nil
}

func (ln *Linkedin) jsessionid() string {
	for _, c := range ln.cookies {
		if c.Name == "JSESSIONID" {
			return c.Value
		}
	}

	return ""
}
