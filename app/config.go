package app

import (
	"database/sql"
	"github.com/cagox/config"
	"github.com/gorilla/mux"
	"log"
	"os"
)

var Config *ConfigStruct

type ConfigStruct struct {
	config.ConfigurationStruct

	SiteName string
	LogPath  string

	DatabaseUserName string
	DatabasePassword string
	DatabaseName     string
	DatabaseOptions  string

	//The items below are not in the JSON file.
	Router  *mux.Router
	Logger  *log.Logger
	LogFile *os.File

	//Database Related
	Database *sql.DB
}

func init() {
	Config = &ConfigStruct{}
	loadConfigs()
	Config.Router = mux.NewRouter()
}

func loadConfigs() {
	config.LoadConfigs(Config, "FLUXCONF")
}
