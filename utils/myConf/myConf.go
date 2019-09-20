package myConf

import (
	"gopkg.in/ini.v1"
)

var iniConf *ini.File

const confPath = "main.ini"

func init() {
	var err error
	iniConf, err = ini.Load(confPath)
	if err != nil {
		panic(err)
	}
	// 配置只读，可以提升性能
	iniConf.BlockMode = false
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
