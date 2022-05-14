package global

import openapiclient "github.com/KangDroid-Cloud/CoreNetworkCommunication"

var DefaultServerConfiguration = openapiclient.Configuration{
	DefaultHeader: make(map[string]string),
	UserAgent:     "OpenAPI-Generator/1.0.0/go",
	Debug:         false,
	Servers: openapiclient.ServerConfigurations{
		{
			URL:         "https://localhost:7121",
			Description: "Local Server API",
		},
	},
}
