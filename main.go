package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	list := gdb.List{}
	var data = []map[string]interface{}{
		{"CREATED_TIME": "2022-09-20 17:54:57.2125326 +0800 CST m=+0.030851101",
			"ACCOUNT_NAME": "user_3",
			"PWD_RESET":    0,
			"ENABLED":      0,
			"CREATED_BY":   "",
			"UPDATED_TIME": nil,
			"ID":           5556,
			"UPDATED_BY":   "",
			"DELETED":      0,
		},
		{
			"UPDATED_BY":   "",
			"ID":           22,
			"PWD_RESET":    0,
			"DELETED":      0,
			"CREATED_BY":   "",
			"UPDATED_TIME": nil,
			"CREATED_TIME": "2022-09-20 17:54:57.2125326 +0800 CST m=+0.030851101",
			"ENABLED":      0,
			"ACCOUNT_NAME": "user_3",
		},
	}
	list = append(list, data...)
	g.Dump("===========================list===========================", list)
	var keys []string
	var values []string
	var selvalues []string
	charL, charR := `"`, `"`
	valuecharL, valuecharR := "'", "'"
	for k := range list[0] {
		keys = append(keys, k)
	}

	for _, column := range keys {
		fmt.Println("===========================")
		fmt.Println(list[0])
		fmt.Println("column:", column)
		fmt.Println("list[0][name]:", list[0][column])
		if list[0][column] == nil {
			fmt.Println("list[0][column]:", list[0][column])

			continue
		}
		va := reflect.ValueOf(list[0][column])
		ty := reflect.TypeOf(list[0][column])
		d := ""
		switch ty.Kind() {
		case reflect.String:
			d = va.String()
		case reflect.Int:
			d = strconv.FormatInt(va.Int(), 10)
		case reflect.Int64:
			d = strconv.FormatInt(va.Int(), 10)
		}
		selvalues = append(selvalues, fmt.Sprintf(valuecharL+"%s"+valuecharR+" AS "+charL+"%s"+charR, d, column))
	}

	for _, mapper := range list[1:] {
		var element []string
		for _, column := range keys {
			if mapper[column] == nil {
				continue
			}
			va := reflect.ValueOf(mapper[column])
			ty := reflect.TypeOf(mapper[column])
			switch ty.Kind() {
			case reflect.String:
				element = append(element, valuecharL+va.String()+valuecharR)
			case reflect.Int:
				element = append(element, strconv.FormatInt(va.Int(), 10))
			case reflect.Int64:
				element = append(element, strconv.FormatInt(va.Int(), 10))
			}
		}
		values = append(values, fmt.Sprintf(`UNION ALL SELECT %s FROM DUAL`, strings.Join(element, ",")))
	}
	fmt.Println(strings.Join(selvalues, ","))

	fmt.Println(strings.Join(values, " "))
	// pointer6()
	// a := tt(Ttet)
	// a(5, 6)
	// theoryContext()
	// theoryReturn()
	// theoryRountine()
	// interSt()
	// interIn()
	// slice_interface()
	// theoryStruct()
	// main function: requirment
	// I hope to Never break down at any time
	// base1()
	// baseA()
	// L1()
	// GMP()
	// f := Foo{
	// 	I: implOfI{},
	// 	J: implOfJ{},
	// }
	// println(f.String())

	// type A interface {
	// 	a()
	// 	String() string
	// }
	// type B interface {
	// 	b()
	// 	String() string
	// }
	// //  ----- (1)go1.14
	// type O interface {
	// 	A
	// 	B
	// }
	// // ---- (2)go1.13
	// type O2 interface {
	// 	a()
	// 	b()
	// 	String() string
	// }
	// zheng("test")
	// \w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}

	// reg := regexp.MustCompile(`^([A-Za-z]|[\u4e00-\u9fa5])+(\d|_|-|[A-Za-z]|[\u4e00-\u9fa5])+`)
	// fmt.Println(reg.MatchString("QWE"))

	// fmt.Println(Multiple3And5(10))
	// Init()
	// if cmd.UUID == "" || cmd.File == "" || cmd.CldType == "" || cmd.CldID == "" {
	// 	fmt.Println("以下命令行参数必填，且信息均为真实有效")
	// 	flag.PrintDefaults()
	// 	return
	// }
	// fmt.Println(fmt.Sprintf("command: %+v", cmd))
	// mysqlRule7()
}
