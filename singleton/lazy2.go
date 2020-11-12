package singleton

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var lazy2Instance *lazyLog2

//
// 小写隐藏，避免外部调用创建实例
//
type lazyLog2 struct {
}

func (l *lazyLog2) Log(msg ...interface{}) {
	fmt.Println(msg...)
}

//
// 用到才初始化
//
func LazyInstance() *lazyLog {

	// 加锁 + 双重级别的判断
	if lazy2Instance == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if lazy2Instance == nil {
			lazy2Instance = &lazyLog2{}
		}
	}

	return lazyInstance
}
