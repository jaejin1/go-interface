package config

import (
	"go-interface/pkg/resource/node/v1"
	"go-interface/pkg/rest"
)

func GetConfig(configPath string) (*v1.NodeV1Client, error) {
	//TODO read config file
	config := &rest.Config{}

	client, err := v1.NewForConfig(config)

	return client, err
}