package model

import (
	"fmt"
	"runtime"
	"path"
	"log"
	"encoding/json"
	//"path/filepath"
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

	fmt.Println(configPath)
	configStruct := Config{}
	
	config := io.LoadFile()

	err := json.Unmarshal(config, &configStruct)

	if err != nil {
		log.Fatal("Config loading: ", err)
	}

	return configStruct
}