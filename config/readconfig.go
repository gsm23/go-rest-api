
package config

import (

	"github.com/gsm23/go-rest-api/common"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type (
	ApiServer = common.WebServer
	Kafka = common.Kafka
	Application = common.Application
)

type Configuration struct {
	ApiServer	ApiServer		`yaml:"server,omitempty"`
	Kafka		Kafka			`yaml:"kafka,omitempty"`
	App			Application		`yaml:"application,omitempty"`
}

func ReadConfigYaml(configFile string, obj *Configuration ) {
	configsrc, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
	err = yaml.Unmarshal(configsrc, obj)
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
}