package myLog

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/lfz757077613/goLearn/utils"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	_infoPath = "info.log"
	_errorPath = "error.log"
	_rotationHour = 24
	_maxAgeHour = 72
	_timeStampFormat = "2016-01-02 15:04:05"
)

// 系统监控直接使用Info、Error等，业务监控使用GetUidTraceLog获得和uid绑定的log
func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: _timeStampFormat,
	})
	// 不向控制台打印
	//log.Out = ioutil.Discard
	// 写文件不加锁，具体参考logrus文档确认可不可以不锁
	// log.SetNoLock()
	infoWriter, err := rotatelogs.New(
		_infoPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(_infoPath),
		rotatelogs.WithRotationTime(time.Hour*_rotationHour),
		rotatelogs.WithMaxAge(time.Hour*_maxAgeHour),
	)
	if err != nil {
		panic(err)
	}
	errorWriter, err := rotatelogs.New(
		_errorPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(_errorPath),
		rotatelogs.WithRotationTime(time.Hour*_rotationHour),
		rotatelogs.WithMaxAge(time.Hour*_maxAgeHour),
	)
	if err != nil {
		panic(err)
	}
	logrus.AddHook(&DefaultFieldsHook{})
	logrus.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: errorWriter,
			logrus.PanicLevel: errorWriter,
		},
		&logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: _timeStampFormat,
		},
	))
}

type DefaultFieldsHook struct {
}

func (df *DefaultFieldsHook) Fire(entry *logrus.Entry) error {
	entry.Data["localIp"] = utils.GetLocalIp()
	return nil
}

func (df *DefaultFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func GetUidTraceLog(c *gin.Context) *logrus.Entry {
	if traceLog, ok := c.Get("traceLog"); ok {
		return traceLog.(*logrus.Entry)
	}
	return logrus.WithField("localIp", utils.GetLocalIp())
}
