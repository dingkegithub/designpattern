package decorator

import "fmt"

type IA interface {
	F()
}

//
// A
//
type A struct {
}

func (a *A) F() {
	fmt.Println("IA Strong")
}

//
// ADecorator
//
type ADecorator struct {
	a IA
}

func NewADecorator(a IA) IA {
	return &ADecorator{
		a: a,
	}
}

func (a *ADecorator) F() {
	a.a.F()
	fmt.Println("IA Stronger")
}

//
// ADecorator
//
type BDecorator struct {
	a IA
}

func NewBDecorator(a IA) IA {
	return &BDecorator{
		a: a,
	}
}

func (b *BDecorator) F() {
	b.a.F()
	fmt.Println("IA Strongest")
}

func usage() {
	a := new(A)
	// 装饰
	aStrong := NewADecorator(a)

	// 再装饰
	aStronger := NewBDecorator(aStrong)
	aStronger.F()
}

/*
输出：

IA Strong
IA Stronger
IA Strongest
*/
