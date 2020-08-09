package versions

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Client accesses a target page
type Client interface {
	List() ([]string, error)
}

type client struct {
	url        url.URL
	httpClient *http.Client
}

// NewClient initializes the client
func NewClient() Client {
	c := defaultClient()
	return c
}

func defaultClient() *client {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "golang.org"
	u.Path = "dl"
	return &client{
		url:        u,
		httpClient: http.DefaultClient,
	}
}

func (c *client) List() ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, c.url.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var versions []string
	doc.Find(".container a[class='download']").Each(func(i int, s *goquery.Selection) {
		u, _ := s.Attr("href")
		// e.g. u = /dl/go1.14.4.linux-amd64.tar.gz
		versions = append(versions, fmt.Sprintf("%s://%s%s", c.url.Scheme, c.url.Hostname(), u))
	})
	return versions, nil
}
