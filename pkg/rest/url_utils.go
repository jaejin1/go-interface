package rest

import (
	"fmt"
	"net/url"
)


func DefaultServerURL(host, apiPath string) (*url.URL, string, error) {
	if host == "" {
		return nil, "", fmt.Errorf("host must be a URL or a host:port pair")
	}
	base := host
	hostURL, err := url.Parse(base)
	if err != nil || hostURL.Scheme == "" || hostURL.Host == "" {
		scheme := "http://"
		//TODO TLS

		hostURL, err = url.Parse(scheme + base)
		if err != nil {
			return nil, "", err
		}
		if hostURL.Path != "" && hostURL.Path != "/" {
			return nil, "", fmt.Errorf("host must be a URL or a host:port pair: %q", base)
		}
	}

	//apiPath version
	versionedAPIPath := apiPath

	return hostURL, versionedAPIPath, nil
}
func defaultServerUrlFor(config *Config) (*url.URL, string ,error) {
	host := config.Host
	if host == "" {
		host = "localhost"
	}

	return DefaultServerURL(host, config.APIPath)
}