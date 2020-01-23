package myLog

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/lfz757077613/goLearn/utils"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

const (
	_infoPath = "info.log"
	_errorPath = "error.log"
	_rotationHour = 24
	_maxAgeHour = 72
	_timeStampFormat = "2016-01-02 15:04:05"
)
var _logEntry *logrus.Entry
// 系统监控直接使用Info、Error等，业务监控使用GetUidTraceLog获得和uid绑定的log
func init() {
	log := logrus.StandardLogger()
	_logEntry = log.WithField("localIp", utils.GetLocalIp())
	log.SetFormatter(&logrus.TextFormatter{
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
	allWriter := io.MultiWriter(infoWriter, errorWriter)
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: allWriter,
			logrus.PanicLevel: allWriter,
		},
		&logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: _timeStampFormat,
		},
	))
}

func WithField(key string, value interface{}) *logrus.Entry {
	return _logEntry.WithField(key, value)
}

func Info(args ...interface{}) {
	_logEntry.Info(args...)
}

func Infof(format string, args ...interface{}) {
	_logEntry.Infof(format, args...)
}

func Error(args ...interface{}) {
	_logEntry.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	_logEntry.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	_logEntry.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	_logEntry.Panicf(format, args...)
}

func GetUidTraceLog(c *gin.Context) *logrus.Entry {
	if traceLog, ok := c.Get("traceLog"); ok {
		return traceLog.(*logrus.Entry)
	}
	return _logEntry
}
