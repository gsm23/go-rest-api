package config

import (
	"fmt"
	//	"gopkg.in/yaml.v1"
	//"io/ioutil"
	"path/filepath"
)

type Config struct {
	Description string
	Fruits      map[string][]string
}

func Readfile() {
	fmt.Println("Inside Config Module....")
	filename, err := filepath.Rel("../")
	if err != nil {
		panic(err)
	}
	fmt.Println(filename)
	//	yamlFile, err := ioutil.ReadFile(filename)
	//	var config Config
	//	err = yaml.Unmarshal(yamlFile, &config)
	//	if err != nil {
	//		panic(err)
	//	}

	/*	if err != nil {
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
	*/
}
