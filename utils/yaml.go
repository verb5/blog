package utils

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type ServerInfoYaml struct {
	ServerInfo SystemInfo `yaml:"server_info"`
}

type SystemInfoYaml struct {
	Location string `yaml:"location"`
	Year     uint   `yaml:"year"`
	Email    string `yaml:"email"`
}

func ParseYaml(f string) *ServerInfoYaml {
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		panic("unable to read config file")
	}

	server_info := &ServerInfoYaml{}
	err = yaml.Unmarshal(yamlFile, &server_info)
	if err != nil {
		panic("not a valid yaml file")
	}
	return server_info

}
