package builder

type ResourcePoolConfig struct {
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
// 问题 1：若是资源池的属性变多，外部要设置，构造函数的入参就变的很多
//
// 问题 2：参数过多，可读性差，还可能传错参数而不自知，比如maxIdle 和 minIdle 弄反
//
func NewResourcePoolConfig(name string, maxTotal, maxIdle, minIdle int) *ResourcePoolConfig {
	return &ResourcePoolConfig{
		name:     name,
		maxTotal: maxTotal,
		maxIdle:  maxIdle,
		minIdle:  minIdle,
	}
}

//
// 问题 1 和 问题 2 首先想到的是用 set 方法解决, 必填项目放到构造函数中，可选项放到set方法中
//
//    rcsPool := NewResourcePoolConfig("ThreadPool")
//    rcsPool.SetMaxTotal(10).SetMaxIdle(5).SetMinIdle(3)
//
// 引入问题：
//          问题3：必填项再增多，仍然使得构造函数变长，假如必填项放到 set 方法中，那如何校验必填项是否设置(用户不掉)
//
//          问题4：配置项假设有依赖关系，设置一个就必须设置另外的，比如 maxIdle 和 minIdle 必须小于 maxTotal，这些逻辑放在何处
//
//          问题5：若ResourceePoolConfig 是创建后不可改对象，那么久不能set 方法向外暴露改属性的接口，此时如何设置属性
//
//
func NewResourcePoolConfig2(name string) *ResourcePoolConfig {
	return &ResourcePoolConfig{
		name:     name, // name 是必填项
		maxTotal: 5,    // 可选项
		maxIdle:  3,    // 可选项
		minIdle:  1,    // 可选项
	}
}

//
// set 方法设置可选项
//
func (rp *ResourcePoolConfig) SetMaxTotal(maxTotal int) *ResourcePoolConfig {
	rp.maxTotal = maxTotal
	return rp
}

func (rp *ResourcePoolConfig) SetMaxIdle(maxIdle int) *ResourcePoolConfig {
	rp.maxIdle = maxIdle
	return rp
}

func (rp *ResourcePoolConfig) SetMinIdle(minIdle int) *ResourcePoolConfig {
	rp.minIdle = minIdle
	return rp
}

//
// 创建者模式此时派上用场了：把校验逻辑放到 Builder 类中，先创建建造者，并且通过 Set()
// 方法设置建造者的变量值，然后在使用 build() 方法真真创建对象之前，做集中的校验，校验
// 通过后才会创建对象，另外将类的构造函数改为私有权限，这样我们就只能通过建造者来创建
// 该类的对象，并且不提供 Set 方法，这样类成了就不可变对象了
//
type Builder struct {
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
// 隐藏构造函数
//
func newResoucePoolConfig(b *Builder) *ResourcePoolConfig {
	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}
}

func (b *Builder) Build() {
	if len(b.name) <= 0 {
		panic("resource pool name blank")
	}

	if b.maxIdle > b.maxTotal {
		panic("idle > total")
	}

	if b.minIdle > b.maxTotal || b.minIdle > b.maxIdle {
		panic("mindIdle > maxTotal, minIdle > maxIdle")
	}

	return
}

func (b *Builder) SetName(name string) *Builder {
	if len(name) <= 0 {
		panic("resource pool name blank")
	}
	b.name = name
	return b
}

func (b *Builder) SetMaxIdle(maxIdle int) *Builder {
	if maxIdle < 0 {
		panic("maxIdle less than 0")
	}
	b.maxIdle = maxIdle
	return b
}

func (b *Builder) SetMaxTotal(maxTotal int) *Builder {
	if maxTotal < 0 {
		panic("maxTotal less than 0")
	}
	b.maxTotal = maxTotal
	return b
}

func (b *Builder) SetMinIdle(minIdle int) *Builder {
	if minIdle < 0 {
		panic("minIdle less than 0")
	}
	b.minIdle = minIdle
	return b
}
