package config

import (
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"path/filepath"
  )

  type Config struct {
	Description string
	Fruits map[string][]string
  }

func readfile() {
	fmt.Println("Inside Config Module....")
	filename, _ := filepath.Abs("../config.yaml")
  	yamlFile, err := ioutil.ReadFile(filename)

  	if err != nil {
    	panic(err)
  	}

  	var config Config

  	err = yaml.Unmarshal(yamlFile, &config)
  	if err != nil {
    	panic(err)
  	}

  	fmt.Printf("Value: %#v\n", config.Description)
  	fmt.Printf("Value: %#v\n", config.Fruits)
}
}
