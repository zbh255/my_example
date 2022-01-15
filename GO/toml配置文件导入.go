package main

import (
	"github.com/BurntSushi/toml"
	"os"
)

type AutoGenerated struct {
	Run struct {
		Ipaddr string `toml:"ipaddr"`
		Port   string `toml:"port"`
		Mode   string `toml:"mode"`
	} `toml:"run"`
	Operation struct {
		Database     string `toml:"database"`
		Oauth        string `toml:"oauth"`
		JwtStateSave bool   `toml:"jwt_state_save"`
	} `toml:"operation"`
	Mysql struct {
		Maxidle  int    `toml:"maxidle"`
		Maxopen  int    `toml:"maxopen"`
		Debug    bool   `toml:"debug"`
		Username string `toml:"username"`
		Dbname   string `toml:"dbname"`
		Password string `toml:"password"`
		Ipaddr   string `toml:"ipaddr"`
		Port     string `toml:"port"`
	} `toml:"mysql"`
	Sqllite struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
		Filepath string `toml:"filepath"`
	} `toml:"sqllite"`
	Model struct {
		Imgsave         string `toml:"imgsave"`
		Localurl        string `toml:"localurl"`
		UploadMax       int    `toml:"uploadMax"`
		Contentfilename string `toml:"contentfilename"`
	} `toml:"model"`
	Tencentcloud struct {
		Bucketurl string `toml:"bucketurl"`
		Secretid  string `toml:"secretid"`
		Secretkey string `toml:"secretkey"`
	} `toml:"tencentcloud"`
	Jwt struct {
		EncodeMethod     string `toml:"encodeMethod"`
		Key              string `toml:"key"`
		MaxEffectiveTime string `toml:"maxEffectiveTime"`
	} `toml:"jwt"`
	Redis struct {
		Ipaddr   string `toml:"ipaddr"`
		Port     string `toml:"port"`
		Password string `toml:"password"`
	} `toml:"redis"`
}

type ConFig interface {
	InitConfig() *Config
}
type Config struct {
	ConfData *AutoGenerated
}

func (c *Config) InitConfig() *Config {
	var conf AutoGenerated
	apath, _ := os.Getwd()
	// bpath := "/conf/pro/app.conf.toml"
	cpath := "/conf/dev/app.conf.toml"
	// capath := "/Users/harder/github.com-codes/go-gobang_game/conf/dev/app.conf.toml"
	if _, err := toml.DecodeFile(apath+cpath, &conf); err != nil {
		println("DecodeFile Toml error / ", err)
	}
	return &Config{ConfData: &conf}
}
