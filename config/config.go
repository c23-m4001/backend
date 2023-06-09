package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         uint16 `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database_name"`
}

type GoogleConfig struct {
	Oauth GoogleOauthConfig `yaml:"oauth"`
}

type GoogleOauthConfig struct {
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectUri  string `yaml:"redirect_uri"`
}

type YamlConfig struct {
	timeLocation *time.Location
	BaseDir      string
	StorageDir   string

	CorsAllowedOrigins []string       `yaml:"cors_allowed_origins"`
	DatabaseConfig     DatabaseConfig `yaml:"database"`
	Environment        string         `yaml:"environment"`
	GeoIpFileName      string         `yaml:"geo_ip_file_name"`
	GoogleConfig       GoogleConfig   `yaml:"google"`
	IsDebug            bool           `yaml:"debug"`
	LogChannels        []string       `yaml:"log_channels"`
	Port               uint           `yaml:"port"`
	Timezone           string         `yaml:"timezone"`
}

var config YamlConfig
var configOnce sync.Once

func GetConfig() YamlConfig {
	configOnce.Do(func() {
		if err := LoadConfig(); err != nil {
			panic(err)
		}
	})

	return config
}

func LoadConfig() error {
	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := os.Stat(fmt.Sprintf("%s/conf.yml", baseDir)); err != nil {
		_, filename, _, _ := runtime.Caller(0)
		baseDir = path.Join(path.Dir(filename), "../")
	}

	config.BaseDir = strings.TrimRight(strings.ReplaceAll(baseDir, "\\\\", "/"), "/")
	config.StorageDir = fmt.Sprintf("%s/storage", config.BaseDir)

	yamlFilePath := fmt.Sprintf("%s/conf.yml", config.BaseDir)
	if _, err := os.Stat(yamlFilePath); err != nil {
		return fmt.Errorf("conf.yml file not found")
	}

	yamlFile, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return err
	}

	err = yaml.UnmarshalStrict(yamlFile, &config)
	if err != nil {
		return err
	}

	config.timeLocation, err = time.LoadLocation(config.Timezone)
	if err != nil {
		return err
	}

	return nil
}

func GetBaseDir() string {
	return GetConfig().BaseDir
}

func GetStorageDir() string {
	return GetConfig().StorageDir
}

func GetJwtPrivateKeyFilePath() string {
	return fmt.Sprintf("%s/jwt/private.key", GetStorageDir())
}

func GetJwtPublicKeyFilePath() string {
	return fmt.Sprintf("%s/jwt/public.key", GetStorageDir())
}

func GetJwtGitIgnoreFilePath() string {
	return fmt.Sprintf("%s/jwt/.gitignore", GetStorageDir())
}

func GetGeoIPFilePath() string {
	return fmt.Sprintf("%s/geoip2/%s", GetStorageDir(), GetConfig().GeoIpFileName)
}

func GetGoogleOauthConfig() GoogleOauthConfig {
	return GetConfig().GoogleConfig.Oauth
}

func GetPostgresConfig() DatabaseConfig {
	return GetConfig().DatabaseConfig
}

func GetLogChannels() []string {
	return GetConfig().LogChannels
}

func GetTimeLocation() *time.Location {
	return GetConfig().timeLocation
}

func IsDebug() bool {
	return GetConfig().IsDebug
}

func IsProduction() bool {
	return !GetConfig().IsDebug
}

func DisableDebug() {
	config.IsDebug = false
}

func EnableDebug() {
	config.IsDebug = true
}
