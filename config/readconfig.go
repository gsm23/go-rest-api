package config

import (
	//"debug/gosym"
	//"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	apiServer	apiServer		`yaml:"server,omitempty"`
	Kafka		Kafka			`yaml:"kafka,omitempty"`
	App			Application		`yaml:"application,omitempty"`
}

type apiServer struct {
	port	int	`yaml:"port,omitempty"`
	debug	bool	`yaml:"debug,omitempty"`
	bindIp	string	`yaml:"bindIp,omitempty"`
	readTimeOutSec	int	`yaml:"read_timeout_seconds,omitempty"`
	writeTimeOutSec	int	`yaml:"write_timeout_seconds,omitempty"`
	methods	[]string	`yaml:"methods,omitempty"`
	http_type	string	`yaml:"http_type,omitempty"`
}

type Kafka struct {
	BootstrapServer string `yaml:"bootstrap_server,omitempty"`
	Topic          string `yaml:"topic,omitempty"`
	Serializer     string `yaml:"serializer,omitempty"`
	BufferMemory   int64  `yaml:"buffer_memory,omitempty"`
	Compression    string `yaml:"compression,omitempty"`
	BatchSize      int64  `yaml:"batch_size,omitempty"`
	ClientID       string `yaml:"client_id,omitempty"` //Provides a uniq name to the application sending the payload
	MaxRequestSize int64  `yaml:"max_request_size,omitempty"`
}

// Application holds the data related to Application
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
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
