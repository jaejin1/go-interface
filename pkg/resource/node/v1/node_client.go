package v1

import "go-interface/pkg/rest"

type NodeV1Client struct {
	restClient rest.Interface
}

func (c *NodeV1Client) Nodes(name string) NodeInterface {
	return newNodes(c, name)
}

func (c *NodeV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func NewForConfig(c *rest.Config) (*NodeV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NodeV1Client{client}, nil
}

func setConfigDefaults(config *rest.Config) error {
	config.Host = "localhost"
	config.APIPath = "/api"

	return nil
}