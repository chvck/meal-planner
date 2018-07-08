package config

import (
	"encoding/json"
	"io/ioutil"
)

// Info contains the application config.
type Info struct {
	DbServer   string `json:"dbServer"`
	DbPort     uint   `json:"dbPort"`
	DbName     string `json:"dbName"`
	HTTPPort   uint   `json:"httpPort"`
	HTTPSPort  uint   `json:"httpsPort"`
	Hostname   string `json:"hostname"`
	AuthKey    string `json:"authKey"`
	DbUsername string `json:"dbUsername"`
	DbPassword string `json:"dbPassword"`
}

// ParseJSON unmarshals bytes to structs
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// Load opens the config json file at the specified path and returns the parsed struct
func Load(path string) (*Info, error) {
	i := Info{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return nil, err
	}

	return &i, nil
}
