package linkedin

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:73.0) Gecko/20100101 Firefox/73.0"

var header = http.Header{
	"Accept-Language":           []string{"en-US,en;q=0.5"},
	"Accept":                    []string{"application/json"},
	"Connection":                []string{"keep-alive"},
	"Host":                      []string{"www.linkedin.com"},
	"Upgrade-Insecure-Requests": []string{"1"},
	"User-Agent":                []string{userAgent},
}

type client struct {
	cookie    string
	ajaxToken string
	proxy     *url.URL
}

func (c client) getRequest(urlParsed *url.URL) (*http.Response, error) {
	req, _ := http.NewRequest("GET", urlParsed.String(), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", c.cookie)
	req.Header.Add("Host", "www.linkedin.com")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("csrf-token", c.ajaxToken)

	cl := &http.Client{
		CheckRedirect: func() func(req *http.Request, via []*http.Request) error {
			redirects := 0
			return func(req *http.Request, via []*http.Request) error {
				if redirects > 10 {
					return errors.New("Linkedin session invalid")
				}

				redirects++
				return nil
			}
		}(),
	}
	if c.proxy != nil {
		cl.Transport = &http.Transport{Proxy: http.ProxyURL(c.proxy), TLSHandshakeTimeout: 20 * time.Second}
	}

	return cl.Do(req)
}
