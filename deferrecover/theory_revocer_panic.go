package main

import (
	"fmt"
	"os"
)

func L1() {
	defer func() {
		fmt.Println("L1 defer func invoked")
	}()
	fmt.Println("L1  invoked")
	L2()
	fmt.Println("do something after L2 in L1")
}

func L2() {
	defer func() {
		fmt.Println("L2 defer func invoked")
		if e := recover(); e != nil {
			fmt.Println("i get you")
		}
	}()
	fmt.Println("L2  invoked")
	L3()
	fmt.Println("do something after L3 in L2")
	panic("runtime exception")
}

func L3() {}

func base1() {
	// this defer&reovce because of main function's required
	defer func() {
		// Capture any errors that occur and do not propagate in the opposite direction of the call stack
		if e := recover(); e != nil {
			fmt.Println(e.(error).Error())
		}
	}()
	// base1 function:
	// I hope to end all work immediately in case of an error at ant time at ant function
	group1()
	group2()
}

func baseA() {
	// this defer&reovce because of main function's required
	defer func() {
		// Capture any errors that occur and do not propagate in the opposite direction of the call stack
		if e := recover(); e != nil {
			fmt.Println(e.(error).Error())
		}
	}()
	// baseA fucntion:
	// I hope to Encounter an error and continue to work at groupA()
	groupA()
	// groupA() error occured but not impacted on next groupB() execution
	// I hope to end all work immediately in case of an error at ant time at groupB()
	groupB()
}
func group1() {
	// panic occured case1: err deal
	field1, err := field1()
	// panic occured case2: index out of range
	fmt.Println(field1[4], err)
	// panic occured case3: err deal of origin package
	f, err := os.Open("test.txt")
	fmt.Println(f, err)
	// panic occured case4: conversion
	var case3 interface{}
	fmt.Println(case3.(string))
}

func group2() {
	// any other panic errors occured
}

func field1() ([]string, error) {
	field1 := []string{"test"}
	return field1, nil
}

func groupA() {
	// this defer&reovce because of BaseA function's required
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e.(error).Error())
		}
	}()
	// fieldA()
	panic("errors occured")
}

func groupB() {
	// any work
}
