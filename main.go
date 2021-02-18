package main

import (
	"flag"
	"fmt"
	"github.com/gsm23/go-rest-api/api"
	"github.com/gsm23/go-rest-api/config"
	"log"
	"github.com/gsm23/goreadconfig"
)
var conf string

type (
	C = config.Configuration
)

func main() {
	var cfg C
	err := goreadconfig.ReadConfig("config.yaml", &cfg)
	if err != nil {
		log.Printf("Error reading configuration from file:%v", "resource/config.yaml" )
	}
	log.Println(*cfg)
	//flag.StringVar(&conf, "file", "config.yaml", "Config file usage")
	//flag.Parse()
	//config.ReadConfigYaml(conf, &f)

	//fmt.Print(reflect.ValueOf(f.ApiServer).Kind())
	//fmt.Print(f.ApiServer)
	//api.TestApi(f.ApiServer.Port)

}
