package config

import (
	"fmt"

	arkconfig "github.com/syariatifaris/arkeus/core/config"
)

//Database configuration
type Database struct {
	Host     string
	Type     string
	Name     string
	User     string
	Password string
	Port     int64
}

type DbConfig struct {
	Database Database
}

type ConfigurationData struct {
	Database Database
}

var (
	directories = []string{
		"files/conf/",
		"../src/github.com/syariatifaris/kumparan/files/conf/",
	}

	extension = "toml"
)

type Config interface {
	DecodeConfig(c interface{}) error
}

//NewConfiguration reads all configuration
func NewConfiguration() *ConfigurationData {
	var dbCfg DbConfig

	err := arkconfig.ReadConfigurationModule("db", extension, directories, &dbCfg)
	if err != nil {
		panic(fmt.Sprintf("unable to resolve db configuration with err %s", err.Error()))
	}

	return &ConfigurationData{
		Database: dbCfg.Database,
	}
}
