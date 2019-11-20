package rest

import "net/http"

type Config struct {
	Host string
	APIPath string
}

func RESTClientFor(config *Config) (*RESTClient, error) {
	baseURL, versionedAPIPath, err := defaultServerUrlFor(config)
	if err != nil {
		return nil, err
	}

	//TODO transport
	var httpClient *http.Client

	return NewRESTClient(baseURL, versionedAPIPath, httpClient)
}