package common

//Webserver configurations
type WebServer struct {
	Port	int	`yaml:"port,omitempty"`
	Debug	bool	`yaml:"debug,omitempty"`
	BindIp	string	`yaml:"bindIp,omitempty"`
	ReadTimeOutSec	int	`yaml:"read_timeout_seconds,omitempty"`
	WriteTimeOutSec	int	`yaml:"write_timeout_seconds,omitempty"`
	//Methods	[]string	`yaml:"methods,omitempty"`
	contentType	string	`yaml:"contentType,omitempty"`
}

//KAfka specific configurations
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
	MinPasswordStr int	`yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
	BootstrapServers	string	`yaml:"bootstrap_server,omitempty"`
}

