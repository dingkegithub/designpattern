>
>
>[1. 创建者模式场景](#1)
>
>[2. 创建者模式模型](#2)
>
>[3. 与工厂模式的区别](#3)
>

<h4 id='1'> 1. 创建者模式场景</h4>

如果一个类中有很多属性，为了避免构造函数的参数列表过长，影响代码的可读性和易用性，
我们可以通过构造函数配合 set() 方法来解决。但是，如果存在下面情况中的任意一种，我
们就要考虑使用建造者模式了

- 我们把类的必填属性放到构造函数中，强制创建对象的时候就设置。如果必填的属性有很多，把这些必填属性都放到构造函数中设置，那构造函数就又会出现参数列表很长的问题。如果我们把必填属性通过 set() 方法设置，那校验这些必填属性是否已经填写的逻辑就无处安放了

- 如果类的属性之间有一定的依赖关系或者约束条件，我们继续使用构造函数配合 set() 方法的设计思路，那这些依赖关系或约束条件的校验逻辑就无处安放了

- 如果我们希望创建不可变对象，也就是说，对象在创建好之后，就不能再修改内部的属性值，要实现这个功能，我们就不能在类中暴露 set() 方法。构造函数配合 set() 方法来设置属性值的方式就不适用了。

总而言之，构造函数必填参数过多、属性之间依赖关系需要判断、对象创建后不可更改需隐藏set方法

<h4 id='2'> 2. 创建者模式模型</h4>


- 被创建者对象

```
type xxx struct {
    prop1 int
    prop2 string
    ...
}

func newXxx(b *XxxBuilder) *xxx {
    return &xxx {
        prop1: b.prop1,
        prop2: b.prop2,
    } 
}
```

- 创建者

```
type XxxBuilder struct {
    //
    // 属性需要重新定义一遍
    //
    prop1 int
    prop2 string
    ...
}

//
// builder 构造函数对外开放
//
func NewXxxBuilder() *XxxBuilder {
    return &XxxBuilder{}
}

//
// Build 方法返回被构造的对象
//
// Build 里面对参数进行校验
//
func (xb *XxxBuilder) Build() *xxx {
    //
    // 校验逻辑
    //
    if xb.prop1 < 0 {
        paninc("xxx prop1 invalid)
    }

    //
    // 校验逻辑
    //
    if xb.prop2 == "" {
        paninc("xxx prop2 invalid)
    }

    //
    // 返回和逻辑的，合法的被构造对象
    //
    return newXxx(xb)
}

//
// 参数设置返回builder本身
//
func (xb *XxxBuilder) SetProp1(a int) *XxxBuilder {
    xb.prop1 = a
    return xb
}

func (xb *XxxBuilder) SetProp2(b string) *XxxBuilder {
    xb.prop2 = b
    return xb
}

```

<h4 id='3'> 3. 与工厂模式的区别</h4>

工厂模式时通过工厂来创建对象，建造者模式让建造者负责创建对象，它们之间有什么区别呢？

- 工厂模式是用来创建不同但是相关类型的对象（继承父类或者实现接口的一组子类），给定参数决定创建那种类型对象

- 建造者模式是用来创建一种类型的复杂对象，通过设置不同的可选参数，定制化的创建不同的参数。本质来讲是构建属性不同的对象
