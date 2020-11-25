
- 装饰器类和原始类继承自同样的父类，这样可以对原始类嵌套多个

- 装饰器类的意图是增强功能

<h2 id='1'> 1、代理模式和装饰器模式区别 </h2>

&ensp;&ensp;&ensp;&ensp;尽管装饰和代理代码结构相同，但是它们的意图不同，代理类添加的是<br/>
和原始功能无关的逻辑；而装饰器是对原始功能的增强

- 代理模式模板

```
//
// 被代理类的类接口
//
type IA interface {
    F()
}

//
// 被代理类
//
type A struct {
}

func (a *A)F() {
    // todo
}

//
// 代理类
//
type AProxy struct {
    // 组合被代理类
    a  IA
}

func NewAProxy(a IA) IA {
    return &AProxy {
        a: a,
    }
} 

func (a *AProxy)F() {
    // 新添加的代码逻辑
    a.F()
    // 新添加的代码逻辑
}
```


- 装饰器模板

```
//
// 原始类接口
//
type IA interface {
    F()
}

//
// 原始类
//
type A struct {
}

func (a *A)F() {
    // todo
}

//
// 装饰类
//
type ADecorator struct {
    // 组合被装饰类
    a  IA
}

func NewADecorator(a IA) IA {
    return &ADecorator {
        a: a,
    }
}

func (a *ADecorator) F() {
    // 功能增强代码
    a.F()
    // 功能增强代码
}

```
