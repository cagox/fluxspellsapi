package app

import (
	"github.com/cagox/config"
	"github.com/gorilla/mux"
)

var Config *ConfigStruct

type ConfigStruct struct {
	config.ConfigurationStruct

	SiteName string

	//The items below are not in the JSON file.
	Router *mux.Router
}

func init() {
	Config = &ConfigStruct{}
	loadConfigs()
	Config.Router = mux.NewRouter()
}

func loadConfigs() {
	config.LoadConfigs(Config, "FLUXCONF")
}
