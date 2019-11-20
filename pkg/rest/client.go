package rest

import (
	"net/http"
	"net/url"
	"strings"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Get() *Request
	Delete() *Request
}

type RESTClient struct {
	base   *url.URL
	Client *http.Client

	//TODO APIVersion()
	versionedAPIPath string
}


func NewRESTClient(baseURL *url.URL, versionedAPIPath string, client *http.Client) (*RESTClient, error) {
	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawPath = ""
	base.Fragment = ""

	return &RESTClient{
		base: &base,
		versionedAPIPath: versionedAPIPath,

		Client: client,
	}, nil
}

func (c *RESTClient) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

func (c *RESTClient) Post() *Request {
	return c.Verb("POST")
}

func (c *RESTClient) Put() *Request {
	return c.Verb("PUT")
}

func (c *RESTClient) Get() *Request {
	return c.Verb("GET")
}

func (c *RESTClient) Delete() *Request {
	return c.Verb("DELETE")
}