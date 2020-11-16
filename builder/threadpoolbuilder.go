package builder

import (
	"fmt"
)

//
// 实际要被构建的对象
//
type ThreadPool struct {
	// 资源池名称
	name string

	// 资源总量
	maxTotal int

	// 最大空闲数目
	maxIdle int

	// 最小空闲数目
	minIdle int
}

func (tp *ThreadPool) Status() {
	fmt.Printf("name: %s; maxTotal: %d; maxIdle: %d; minIdle: %d\n", tp.name, tp.maxTotal, tp.maxIdle, tp.minIdle)
}

//
// 隐藏构造函数
// 一旦创建，其不可更改,
// 创建参数的验证全部交给 Builder
//
//  用法：
//
//     build := ThreadPoolBuilder{}
//     tpPool := build.SetName("threadpool").SetMaxTotal(10).SetMaxIdle(6).SetMinIdle(3).Build()
//
func newThreadPool(tb *ThreadPoolBuilder) *ThreadPool {
	return &ThreadPool{
		name:     tb.name,
		maxIdle:  tb.maxIdle,
		maxTotal: tb.maxTotal,
		minIdle:  tb.minIdle,
	}
}

// Builder
//
// 创建者模式此时派上用场了：把校验逻辑放到 Builder 类中，先创建建造者，并且通过 Set()
// 方法设置建造者的变量值，然后在使用 build() 方法真真创建对象之前，做集中的校验，校验
// 通过后才会创建对象，另外将类的构造函数改为私有权限，这样我们就只能通过建造者来创建
// 该类的对象，并且不提供 Set 方法，这样类成了就不可变对象了
//
type ThreadPoolBuilder struct {
	// 资源池名称
	name string

	// 资源总量
	maxTotal int

	// 最大空闲数目
	maxIdle int

	// 最小空闲数目
	minIdle int
}

//
// Build 方法返回构建的时机对象，这里是线程池
//       校验逻辑全部在Build 中
//
func (b *ThreadPoolBuilder) Build() *ThreadPool {
	if len(b.name) <= 0 {
		panic("resource pool name blank")
	}

	if b.maxIdle > b.maxTotal {
		panic("idle > total")
	}

	if b.minIdle > b.maxTotal || b.minIdle > b.maxIdle {
		panic("mindIdle > maxTotal, minIdle > maxIdle")
	}

	return newThreadPool(b)
}

//
// Build 开放set 方法
//
func (b *ThreadPoolBuilder) SetName(name string) *ThreadPoolBuilder {
	if len(name) <= 0 {
		panic("resource pool name blank")
	}
	b.name = name
	return b
}

func (b *ThreadPoolBuilder) SetMaxIdle(maxIdle int) *ThreadPoolBuilder {
	if maxIdle < 0 {
		panic("maxIdle less than 0")
	}
	b.maxIdle = maxIdle
	return b
}

func (b *ThreadPoolBuilder) SetMaxTotal(maxTotal int) *ThreadPoolBuilder {
	if maxTotal < 0 {
		panic("maxTotal less than 0")
	}
	b.maxTotal = maxTotal
	return b
}

func (b *ThreadPoolBuilder) SetMinIdle(minIdle int) *ThreadPoolBuilder {
	if minIdle < 0 {
		panic("minIdle less than 0")
	}
	b.minIdle = minIdle
	return b
}
