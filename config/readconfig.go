package config

type Config struct {
	Description string
	Fruits      map[string][]string
}
/*
func Readfile() {
	fmt.Println("Inside Config Module....")
	filename, err := filepath.Rel(".", "../config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(filename)
	config, err = ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &Config)
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
