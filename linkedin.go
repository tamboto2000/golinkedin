package linkedin

import (
	"errors"
	"net/http"
	"net/url"
)

type Linkedin struct {
	client client
}

func New(cookie string) (*Linkedin, error) {
	ajaxToken := findAjaxToken(cookie)
	if ajaxToken == "" {
		return nil, errors.New("check your cookie. Error: ajax token not found")
	}

	ln := new(Linkedin)
	ln.client.cookie = cookie
	ln.client.ajaxToken = ajaxToken

	return ln, nil
}

func (ln *Linkedin) SetProxy(urlstr string) error {
	prox, err := url.Parse(urlstr)
	if err != nil {
		return err
	}

	ln.client.proxy = prox
	return nil
}

func findAjaxToken(r string) string {
	header := http.Header{}
	header.Add("Cookie", r)
	request := http.Request{Header: header}
	for _, c := range request.Cookies() {
		if c.Name == "JSESSIONID" {
			return c.Value
		}
	}

	return ""
}
