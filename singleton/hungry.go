package singleton

import "fmt"

// 加载时已经创建好
var instance *hangryLog = &hangryLog{}

type hangryLog struct {
}

func (hl *hangryLog) Log(msg ...interface{}) {
	fmt.Println(msg...)
}

func GetInstance() *hangryLog {
	return instance
}
