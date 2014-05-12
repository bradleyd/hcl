package config

import (
  "os"
  "path"
  "testing"
)

func TestConfigLoadsFile(t *testing.T) {
  config, err := Load("./.hcl")
  if err !=nil {
    t.Error(err)
  }

  if config.ApiToken != "12345" {
    t.Error("expected 12345, but got", config)
  }
}

func TestConfigPath(t *testing.T) {
  home := os.Getenv("HOME")
  cp := path.Join(home, ".hcl")
  if configPath() != cp {
    t.Errorf("expected %s, but got %v", cp, configPath())
  }
}
