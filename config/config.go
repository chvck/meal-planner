package config

import (
	"encoding/json"
	"io/ioutil"
)

// Info contains the application config.
type Info struct {
	DbString  string `json:"dbString"`
	DbType    string `json:"dbType"`
	HttpPort  int    `json:"httpPort"`
	HttpsPort int    `json:"httpsPort"`
	Hostname  string `json:"hostname"`
}

type Connection struct {
}

// ParseJSON unmarshals bytes to structs
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

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
