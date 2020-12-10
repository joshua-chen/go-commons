/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 09:21:47
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:55:20
 */
package config

import (
	_ "io/ioutil"

	_ "github.com/kataras/golog"
	_ "gopkg.in/yaml.v2"

)

var AppConfig AppInfo

type AppInfo struct {
	Port       string   `yaml:"port"`
	AnonymousRequest *AnonymousRequest `yaml:"anonymousRequest"`
	Static   *Static   `yaml:"static"`
	JwtTimeout int64    `yaml:"jwtTimeout"`
	LogLevel   string   `yaml:"logLevel"`
	Secret     string   `yaml:"secret"`
	APIPrefix  APIPrefix   `yaml:"apiPrefix"`
	SQLPath   []string   `yaml:"sqlPath"`
	ViewPath   []string   `yaml:"viewPath"`
	UploadPath   string   `yaml:"uploadPath"`
	Swagger Swagger   `yaml:"swagger"`
}

type Static  struct{
 	RequestPath string `yaml:"requestPath"`
	Directory  string `yaml:"directory"`
	 
}

type AnonymousRequest struct{
	Path string `yaml:"path"`
	Urls []string `yaml:"urls"`
	Prefixes []string `yaml:"prefixes"`
	Suffixes []string `yaml:"suffixes"`
}
type APIPrefix struct{
	Base string `yaml:"base"`
	Web string `yaml:"web"`
	Wap string `yaml:"wap"`
	Common string `yaml:"common"`
}

type Swagger struct{
	Docs SwaggerDocs `yaml:"docs"`
}

type SwaggerDocs struct{
	Web string `yaml:"web"`
	Wap string `yaml:"wap"`
	Common string `yaml:"common"`
}
