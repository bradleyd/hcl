package config

import (
  "io/ioutil"
  "strings"
  "os"
  "path"
)

func Load(filename ...string) (string, error) {
  fn := configPath() 

  // allow a filename/path to be passed in
  if len(filename) > 0 {
    fn = filename[0]
  }
 
  res, err := ioutil.ReadFile(fn)
  buf := strings.Trim(string(res), "\n")
  return buf, err

}

func homeDir() string {
  return os.Getenv("HOME")
}

func configPath() string {
  return path.Join(homeDir(), ".hcl")
}
