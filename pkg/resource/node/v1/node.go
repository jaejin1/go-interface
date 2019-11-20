package v1

import (
	"go-interface/pkg/rest"
)

type NodeGetter interface {
	Nodes(name string) NodeInterface
}

type NodeInterface interface {
	Create() (string, error)
	Delete() (string, error)
	Update() (string, error)
	GET() (string, error)
}

type nodes struct {
	client rest.Interface
	resourceName string
}

func newNodes(c *NodeV1Client ,resourceName string) *nodes {
	return &nodes{
		client: c.RESTClient(),
		resourceName: resourceName,
	}
}


func (c *nodes) Create() (result string, err error) {
	result = ""
	err = c.client.Post().
		Name(c.resourceName).
		Do().
		Into()
	return
}

func (c *nodes) Delete() (result string, err error) {
	result = ""
	err = c.client.Delete().
		Name(c.resourceName).
		Do().
		Into()
	return
}

func (c *nodes) Update() (result string, err error) {
	result = ""
	err = c.client.Put().
		Name(c.resourceName).
		Do().
		Into()
	return
}

func (c *nodes) GET() (result string, err error) {
	result = ""
	err = c.client.Get().
		Name(c.resourceName).
		Do().
		Into()
	return
}
