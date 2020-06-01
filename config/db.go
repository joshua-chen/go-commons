package config

import (
	_ "github.com/kataras/golog"
	_ "gopkg.in/yaml.v2"
)

var (
	DBConfig DB
)

type DB struct {
	Master DBInfo `yaml:"master"`
	Slave  DBInfo `yaml:"slave"`

	TablePrefixes []TablePrefix `yaml:"tablePrefix"`
}

type TablePrefix struct {
	PrefixName string `yaml:"prefixName"`
	FeatureName string `yaml:"featureName"`
}
type DBInfo struct {
	Dialect      string `yaml:"dialect"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	Charset      string `yaml:"charset"`
	ShowSql      bool   `yaml:"showSql"`
	LogLevel     string `yaml:"logLevel"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`

	//ParseTime       bool   `yaml:"parseTime"`
	//ConnMaxLifetime int64  `yaml:"connMaxLifetime: 10"`
	//Sslmode         string `yaml:"sslmode"`
}
