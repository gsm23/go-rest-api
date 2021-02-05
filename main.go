package main

import (
	"flag"
	"fmt"
	"github.com/gsm23/go-rest-api/config"
)

var conf string

type (
	C = config.Configuration
)

func main() {
	var f C = C{}
	flag.StringVar(&conf, "file", "config.yaml", "Config file usage")
	flag.Parse()
	config.ReadConfigYaml(conf, &f)

	fmt.Print(f.ApiServer)
	//api.MyApp()
}
