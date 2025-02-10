package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var Config *config

// config总配置文件
type config struct {
	System system
	Logger logger
	Mysql  mysql
	Redis  redis
}

// 系统配置
type system struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

// 日志配置
type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

// mysql配置
type mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func init() {
	yamlFile, err := os.ReadFile(`./config.yaml`)
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(yamlFile, &Config)

}
