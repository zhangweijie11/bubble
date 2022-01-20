package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type DbConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

func (dbConf *DbConf) GetDbConf() *DbConf {
	getWd, _ := os.Getwd()
	// 拼接路径
	confPath := getWd + "/config.yml"
	// 读取配置文件
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 解析文件
	err = yaml.Unmarshal(yamlFile, dbConf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return dbConf
}
