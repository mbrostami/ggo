package main

import (
	"github.com/tkanos/gonfig"
)

func Configuration() ConfigStruct {
	var filename = "config/confsig.json"
	configuration := ConfigStruct{}
	gonfig.GetConf(filename, &configuration)
	return configuration
}
