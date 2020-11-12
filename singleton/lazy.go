package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once
var lazyInstance *lazyLog

//
// 小写隐藏，避免外部调用创建实例
//
type lazyLog struct {
}

func (l *lazyLog) Log(msg ...interface{}) {
	fmt.Println(msg...)
}

//
// 用到才初始化
//
func Instance() *lazyLog {
	once.Do(func() {
		lazyInstance = &lazyLog{}
	})

	return lazyInstance
}
