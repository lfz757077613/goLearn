package myLog

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/lfz757077613/goLearn/utils"
	"github.com/lfz757077613/goLearn/utils/myConf"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var logEntry *logrus.Entry

// 系统监控直接使用Info、Error等，业务监控使用GetUidTraceLog获得和uid绑定的log
func init() {
	infoPath := myConf.GetString("server", "infoLogPath", "./info.log")
	errorPath := myConf.GetString("server", "errorLogPath", "./error.log")
	localIp, err := utils.GetLocalIp()
	if err != nil {
		panic(err)
	}
	log := logrus.StandardLogger()
	logEntry = log.WithField("localIp", localIp)
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// mylog.SetNoLock()
	infoWriter, err := rotatelogs.New(
		infoPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(infoPath),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(myConf.GetInt("server", "rotationHour", 24))),
		rotatelogs.WithMaxAge(time.Hour*time.Duration(myConf.GetInt("server", "maxAgeHour", 72))),
	)
	if err != nil {
		panic(err)
	}
	errorWriter, err := rotatelogs.New(
		errorPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(errorPath),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(myConf.GetInt("server", "rotationHour", 24))),
		rotatelogs.WithMaxAge(time.Hour*time.Duration(myConf.GetInt("server", "maxAgeHour", 72))),
	)
	if err != nil {
		panic(err)
	}
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: errorWriter,
		},
		&logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logEntry.WithField(key, value)
}

func Info(args ...interface{}) {
	logEntry.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logEntry.Infof(format, args...)
}

func Error(args ...interface{}) {
	logEntry.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logEntry.Errorf(format, args...)
}

func GetUidTraceLog(c *gin.Context) *logrus.Entry {
	if traceLog, ok := c.Get("traceLog"); ok {
		return traceLog.(*logrus.Entry)
	}
	return logEntry
}
