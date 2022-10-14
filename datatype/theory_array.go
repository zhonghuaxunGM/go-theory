package main

import "fmt"

func Arr() {
	fmt.Println("=========数组与切片===========")
	arr := [2]int{1}
	// 固定长度
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	// 默认值
	fmt.Println(arr)

	sli := []int{1}
	fmt.Println(len(sli))
	fmt.Println(cap(sli))
	fmt.Println(sli)
	sli = append(sli, 2, 3)
	// 自增
	fmt.Println(len(sli))
	// cap为前一数值的2倍
	fmt.Println(cap(sli))
	fmt.Println(sli)
	fmt.Println("=========切片为添加元素===========")
	sli2 := []int{1, 2, 3}
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	fmt.Println(sli2)
	// 切片为添加元素
	sli2 = append(sli2, []int{4, 5}...)
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	fmt.Println(sli2)
	// cap为前一数值的2倍
	sli2 = append(sli2, []int{6, 7}...)
	fmt.Println(len(sli2))
	fmt.Println(cap(sli2))
	fmt.Println(sli2)

	fmt.Println("=========copy slice===========")
	nnslice := make([]int, len(sli2), cap(sli2))
	fmt.Println(copy(nnslice, sli2))
	fmt.Println("====================", nnslice)
	nnslice[0] = -100
	fmt.Println("====================", nnslice)
	nnslice = append(nnslice, -200)
	fmt.Println("====================", nnslice)

	fmt.Println("=========newslice===========")
	newslice := sli2[1:3]
	fmt.Println("====================", newslice)
	fmt.Println(sli2)
	newslice[0] = -1
	fmt.Println("====================", newslice)
	fmt.Println(sli2)
	newslice = append(newslice, -2)
	fmt.Println("====================", newslice)
	fmt.Println(sli2)
}

// 使用切片操作符切取切片时，上界是切片的容量，而非长度。
func cap1() {
	array := [10]uint32{1, 2, 3, 4, 5}
	s1 := array[:5]

	s2 := s1[5:10]

	fmt.Println(s1)
	fmt.Println(s2)

	s1 = append(s1, 6)
	fmt.Println(s1)
	fmt.Println(s2)
}

// 如果 v 的类型是数组或指向数组的指针，且表达式 v 没有包含 channel 接收或（非常量）函数调用，则返回值也是一个常量。
// 这种情况下，不会对 v 进行求值（即编译期就能确定）。否则返回值不是常量，且会对 v 进行求值（即得运行时确定）。
func lengthNotCalculate() {
	var x *struct {
		s [][32]byte
	}

	println(len(x.s[99]))

	var testdata *struct {
		a *[7]int
	}
	for i, _ := range testdata.a {
		fmt.Println(i)
	}

	for i, j := range testdata.a {
		fmt.Println(i, j)
	}
}

// https://golang.org/doc/effective_go 官方文档两句话
// 1 、If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array.
// 2 、We must return the slice afterwards because, although Append can modify the elements of slice, the slice itself (the run-time data structure holding the pointer, length, and capacity) is passed by value.
func sliceLength() {
	sliceA := make([]int, 3, 4)
	sliceA[0] = 0
	sliceA[1] = 1
	sliceA[2] = 2
	fmt.Println(sliceA)

	// 这种情况下 append 对原数组生效，只是由于 len 没有改变而无法呈现出 append 的项。
	changeSlice(sliceA)
	fmt.Println(sliceA)
	fmt.Printf("%d %p main\n", len(sliceA), sliceA)

	fmt.Println(sliceA[:4])
}

func changeSlice(slicePass []int) {
	slicePass = append(slicePass, 3)
	fmt.Printf("%d %p pass\n", len(slicePass), slicePass)
}

func sliceLength2() {
	sliceA := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceA)
	fmt.Printf("%d %p main\n", len(sliceA), sliceA)
	changeSliceA(sliceA)
	fmt.Println(sliceA)
	fmt.Printf("%d %p main\n", len(sliceA), sliceA)
}

func changeSliceA(slicePass []int) {
	slicePass = append(slicePass, 6)
	fmt.Printf("%d %p pass\n", len(slicePass), slicePass)
}
