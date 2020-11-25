package user_srv

import "fmt"

//
// 被代理类接口
//
type User interface {
	Login(name, pwd string)
}

//
// 用户服务
//
type UserSrv struct {
}

func (u *UserSrv) Login(name, pwd string) {
	fmt.Println(name, " :", pwd)
}

//
// 代理类
//
type UserSrvProxy struct {
	qps int
	usr User
}

//
// 代理类
//
// 实现被代理接口
//
func (u *UserSrvProxy) Login(name, pwd string) {
	// 代理额外功能
	u.log()
	u.auth()
	u.limit()
	u.metric()
	u.statistic()
	u.idempotent()
	u.transication()

	// 不改变原始接口
	u.usr.Login(name, pwd)
}

//
// 监控
//
func (u *UserSrvProxy) metric() {
	fmt.Println("metric server")
}

//
// 统计
//
func (u *UserSrvProxy) statistic() {
	u.qps++
	fmt.Println("statistic server: ", u.qps)
}

//
// 鉴权
//
func (u *UserSrvProxy) auth() {
	fmt.Println("auth request")
}

//
// 限流
//
func (u *UserSrvProxy) limit() {
	fmt.Println("limit request")
}

//
// 事务
//
func (u *UserSrvProxy) transication() {
	fmt.Println("transication request")
}

//
// 幂等
//
func (u *UserSrvProxy) idempotent() {
	fmt.Println("idempotent request")
}

//
// 日志
//
func (u *UserSrvProxy) log() {
	fmt.Println("log request")
}
