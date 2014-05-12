package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type ConfigOptions struct {
	ApiToken string `json:"api_token"`
	Color    string `json:"color"`
}

func Load(filename ...string) (ConfigOptions, error) {
	var cfg ConfigOptions
	fn := configPath()
	// allow a filename/path to be passed in
	if len(filename) > 0 {
		fn = filename[0]
	}

	res, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(res, &cfg)
	if err != nil {
		fmt.Println("json error: ", err)
		os.Exit(1)
	}

	return cfg, err
}

func homeDir() string {
	return os.Getenv("HOME")
}

func configPath() string {
	return path.Join(homeDir(), ".hcl")
}
