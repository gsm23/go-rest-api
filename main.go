package main

import (
	//"flag"
	"fmt"

	"github.com/gsm23/go-rest-api/config"
)
var conf string

//type Configuration struct {
//	apiServer	apiServer		`yaml:"server,omitempty"`
//	Kafka		Kafka			`yaml:"kafka,omitempty"`
//	App			Application		`yaml:"application,omitempty"`
//}

func main() {
	var f Configuration = Configuration{}
	cfgpath := "config.yaml"
	//fmt.Println("Hello")
	//flag.StringVar(&conf, "file", "config.yaml", "Config file usage")
	//flag.Parse()
	//cfgpath, err := config.Load(conf)
	config.ReadConfigYaml(cfgpath, &f)
	fmt.Println(f.App)
}
