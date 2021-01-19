package main

import (
	"flag"
	"fmt"

	"github.com/gsm23/go-rest-api/config"
)
var conf string

func main() {
	fmt.Println("Hello")
	flag.StringVar(&conf, "file", "config.yaml", "Config file usage")
	flag.Parse()
	cfgpath, err := config.Load(conf)
	config.Readfile()
}
