package singleton

import "sync"

var dbonce sync.Once
var instances map[string]*dbConnPool

type dbConnPool struct {
}

// 懒汉式
func DbInstance(name string) *dbConnPool {
	dbonce.Do(func() {
		// 初始化三个对象，慢查询，快查询，插入操作
		instances = make(map[string]*dbConnPool)
		instances["slowquery"] = &dbConnPool{}
		instances["quickquery"] = &dbConnPool{}
		instances["insertquery"] = &dbConnPool{}
	})

	return instances[name]
}
