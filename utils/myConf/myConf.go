package myConf

import (
	"flag"
	"gopkg.in/ini.v1"
)

var iniConf *ini.File

func init() {
	confFilePath := flag.String("c", "main.ini", "main conf path")
	// 注释掉才能跑测试
	//flag.Parse()
	var err error
	iniConf, err = ini.Load(*confFilePath)
	if err != nil {
		panic(err)
	}
}

func GetString(section, key, defaultValue string) string {
	return iniConf.Section(section).Key(key).MustString(defaultValue)
}

func GetInt(section, key string, defaultValue int) int {
	return iniConf.Section(section).Key(key).MustInt(defaultValue)
}

func GetInt64(section, key string, defaultValue int64) int64 {
	return iniConf.Section(section).Key(key).MustInt64(defaultValue)
}

func GetFloat64(section, key string, defaultValue float64) float64 {
	return iniConf.Section(section).Key(key).MustFloat64(defaultValue)
}
