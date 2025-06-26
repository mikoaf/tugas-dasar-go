package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type SensorConfig struct {
	DiskPath string `yaml:"disk_path"`
	UpdateInterval int `yaml:"update_interval"`
}

type Config struct {
	Sensors SensorConfig `yaml:"sensors"`
}

func NewConfig(path string) (*Config, error){
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}