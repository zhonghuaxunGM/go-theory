package main

import (
	"fmt"
	"reflect"
)

func pointer2() {
	i := 10
	ip := &i
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	fmt.Printf("原始指针的内存地址是：%p\n", &i)

	pointer1(ip)
	fmt.Println("int值被修改了，新值为:", i)
}

func pointer1(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", ip)

	*ip = 1
}

// ====================================================

type Pter3 struct {
	X int
}

func (t Pter3) M1() int {
	return t.X
}

func (t *Pter3) M2() int {
	return t.X
}

func pointer3() {
	var t1 = new(Pter3)
	var f1 = t1.M1 // <=> (*t1).M1
	t1.X = 9
	println(f1()) // 0

	var t2 Pter3
	var f2 = t2.M2 // <=> (&t2).M2
	t2.X = 9
	println(f2()) // 9
}

// ====================================================

type Pter4 struct {
	X int
}

func (t Pter4) M() int {
	return t.X
}

type S struct {
	*Pter4
}

func pointer41() {
	var s = S{Pter4: new(Pter4)}
	var f = s.M // <=> (*s.Pter4).M
	s.Pter4 = nil
	f()
}

func pointer42() {
	var s S
	var f = s.M // panic
	s.Pter4 = new(Pter4)
	f()
}

func pointer4() {
	pointer41()
	pointer42()
}

func pointer5() {
	var s = S{Pter4: new(Pter4)}
	var f = s.M // <=> (*s.T).M
	var g = reflect.ValueOf(&s).Elem().MethodByName("M").Interface().(func() int)
	var h = interface {
		M() int
	}(s).M
	s.Pter4.X = 3
	println(f()) // 0
	println(g()) // 3
	println(h()) // 3
}

// ====================================================
type Inter6 interface {
	M6()
}

type Pter6 struct {
	x int
}

func (t Pter6) M6() {
	println(t.x)
}

func pointer6() {
	var t = &Pter6{
		x: 1,
	}
	var i Inter6 = t

	var f = i.M6
	defer f() // 2（正确）

	// i.M 将在编译时刻被（错误地）去虚拟化为 (*t).M。
	defer i.M6() // 1（错误）

	t.x = 2
}
