package rest

import (
	"errors"
	"fmt"
	"path"
)

type Request struct {
	c *RESTClient

	// generic components accessible via method setters
	verb string
	pathPrefix string
	resourceName string


	err error
}

type Result struct {
	InterfaceTest string
	err         error
}

func NewRequest(c *RESTClient) *Request {
	var pathPrefix string
	if c.base != nil {
		pathPrefix = path.Join("/", c.base.Path, c.versionedAPIPath)
	} else {
		pathPrefix = path.Join("/", c.versionedAPIPath)
	}

	r := &Request{
		c: c,
		pathPrefix: pathPrefix,
	}

	return r
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

func (r *Request) Name(resourceName string) *Request {
	if r.err != nil {
		return r
	}
	if len(resourceName) == 0 {
		r.err = fmt.Errorf("resource name may not be empty")
		return r
	}
	//TODO IsValidName check

	r.resourceName = resourceName
	return r
}

func (r *Request) Do() Result {
	//TODO
	// err := r.request(func(req *http.Request, resp *http.Response) {
	// result = r.~~(resp, req)
	// })

	//Test
	err := errors.New("test error !")

	if err != nil {
		return Result{err: err}
	}

	//Test
	return Result{InterfaceTest: "test"}
}


func (r Result) Into() error {
	// Into stores the result into obj
	if r.err != nil {
		return r.Error()
	}

	return nil
}

func (r Result) Error() error {
	if r.err != nil {
		return r.err
	}

	//TODO

	return r.err
}