package myConf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lfz757077613/goLearn/utils/shutDownhook"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"io/ioutil"
)

var iniConf *ini.File

const confPath = "main.ini"

func init() {
	var err error
	iniConf, err = ini.Load(confPath)
	if err != nil {
		logrus.Panicf("load conf error: [%s]", err)
	}
	iniConf.BlockMode = false

	// 增加热更新能力
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.Panicf("new watcher error: [%s]", err)
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					bytes, err := ioutil.ReadFile(confPath)
					if err != nil {
						logrus.Errorf("read conf error: [%s]", err)
						break
					}
					logrus.Infof("modified conf file:%s", string(bytes))
					newConf, err := ini.Load(confPath)
					if err != nil {
						logrus.Errorf("reload conf error: [%s]", err)
						break
					}
					newConf.BlockMode = false
					iniConf = newConf
				}
			case err := <-watcher.Errors:
				logrus.Errorf("watch conf file error: [%s]", err)
			}
		}
	}()
	if err := watcher.Add(confPath); err != nil {
		logrus.Panicf("add watcher file error: [%s]", err)
	}
	shutDownhook.AddShutdownHook(func() {
		if err := watcher.Close(); err!=nil {
			logrus.Errorf("myConf watcher close error: [%s]", err)
		}
	})
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
