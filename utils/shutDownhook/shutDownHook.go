package shutDownhook

import (
	"github.com/lfz757077613/goLearn/utils/myLog"
	"reflect"
	"runtime"
)

var shutdownHook []func()

func AddShutdownHook(f func()) {
	funcWrapper := func() {
		defer func() {
			if err := recover(); err != nil {
				funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
				myLog.Errorf("shutdown hook panic wrap [%s]: [%s]", funcName, err)
			}
		}()
		f()
	}
	shutdownHook = append(shutdownHook, funcWrapper)
}

func RunShutdownHook() {
	for _, f := range shutdownHook {
		f()
	}
}
