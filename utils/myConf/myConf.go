package myConf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lfz757077613/goLearn/utils/myLog"
	"github.com/lfz757077613/goLearn/utils/shutDownhook"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"sync/atomic"
)

var iniConf atomic.Value

const confPath = "main.ini"

func init() {
	initConf, err := ini.Load(confPath)
	if err != nil {
		myLog.Panicf("load conf error: [%s]", err)
	}
	initConf.BlockMode = false
	iniConf.Store(initConf)

	// 增加热更新能力
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		myLog.Panicf("new watcher error: [%s]", err)
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					bytes, err := ioutil.ReadFile(confPath)
					if err != nil {
						myLog.Errorf("read conf error: [%s]", err)
						break
					}
					myLog.Infof("modified conf file:%s", string(bytes))
					newConf, err := ini.Load(confPath)
					if err != nil {
						myLog.Errorf("reload conf error: [%s]", err)
						break
					}
					newConf.BlockMode = false
					iniConf.Store(newConf)
				}
			case err := <-watcher.Errors:
				myLog.Errorf("watch conf file error: [%s]", err)
			}
		}
	}()
	if err := watcher.Add(confPath); err != nil {
		myLog.Panicf("add watcher file error: [%s]", err)
	}
	shutDownhook.AddShutdownHook(func() {
		if err := watcher.Close(); err!=nil {
			myLog.Errorf("myConf watcher close error: [%s]", err)
		}
	})
}

func GetString(section, key, defaultValue string) string {
	conf, _ := iniConf.Load().(*ini.File)
	return conf.Section(section).Key(key).MustString(defaultValue)
}

func GetInt(section, key string, defaultValue int) int {
	conf, _ := iniConf.Load().(*ini.File)
	return conf.Section(section).Key(key).MustInt(defaultValue)
}

func GetInt64(section, key string, defaultValue int64) int64 {
	conf, _ := iniConf.Load().(*ini.File)
	return conf.Section(section).Key(key).MustInt64(defaultValue)
}

func GetFloat64(section, key string, defaultValue float64) float64 {
	conf, _ := iniConf.Load().(*ini.File)
	return conf.Section(section).Key(key).MustFloat64(defaultValue)
}
