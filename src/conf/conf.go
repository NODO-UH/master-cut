package conf

import (
	"encoding/json"
	"os"
)

type GroupConfiguration struct {
	Name   *string `json:"name"`
	File   *string `json:"file"`
	Script *string `json:"script"`
}

type MasterCutConfiguration struct {
	Logs   *string              `json:"logs"`
	Groups []GroupConfiguration `json:"groups"`
}

var Configuration MasterCutConfiguration

func SetupConfiguration(confPath string) error {
	confFile, err := os.Open(confPath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(confFile)
	if err := decoder.Decode(&Configuration); err != nil {
		return err
	}
	return nil
}

func GetGroup(group string) *GroupConfiguration {
	for _, g := range Configuration.Groups {
		if *g.Name == group {
			return &g
		}
	}
	return nil
}
