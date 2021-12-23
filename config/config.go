package config

var (
	Config *ProjectConfig
)

func init() {
	Config = &ProjectConfig{}
}

type ProjectConfig struct {
	HttpServer  HttpServerConfig  `json:"http_server" yaml:"http_server" mapstructure:"http_server"`
	Application ApplicationConfig `json:"application" yaml:"application" mapstructure:"application"`
}

type ApplicationConfig struct {
	Mode          string `json:"mod" yaml:"mod" mapstructure:"mode"`
	NodeId        string `json:"node_id" yaml:"node_id" mapstructure:"node_id"`
	ClockSequence int    `json:"clock_sequence" yaml:"clock_sequence" mapstructure:"clock_sequence"`
}

type HttpServerConfig struct {
	Port string `json:"port" yaml:"port" mapstructure:"port"`
}
