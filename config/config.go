package config

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         uint16 `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database_name"`
}

type YamlConfig struct {
	BaseDir        string
	Environment    string         `yaml:"environment"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
}

var config YamlConfig

func init() {

}

func GetConfig() YamlConfig {
	return config
}
