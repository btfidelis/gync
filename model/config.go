package model

import (
	"runtime"
	"path"
	"log"
	"encoding/json"
	"github.com/btfidelis/gync/core"
)

const CONFIG_FILE = "../config.json"

type Config struct {
	BackupPath	string
	CheckInterval int
}

func GetConfig() Config {
	_, file, _, _ := runtime.Caller(1)
	configPath := path.Join(path.Dir(file), CONFIG_FILE)
	io := core.NewIOManager(configPath)
	configStruct := Config{}
	
	config := io.LoadFile()	

	err := json.Unmarshal(config, &configStruct)

	if err != nil {
		log.Fatal("Config loading: ", err)
	}

	return configStruct
}