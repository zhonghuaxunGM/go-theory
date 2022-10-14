package trick

import (
	"fmt"
	"reflect"
)

func Pro() {
	a := &Data{
		Target: &Target{
			ID: "testID",
			// ValueCn: new(string),
			// ValueEn: new(string),
		},
		Link: Link{
			ID: "testID",
			// Schema: "db_cmp",
			Table:  "t_sys_enum",
			Column: []string{"enum_value_cn", "enum_value_en"},
		},
	}
	Watch(a, Rule)
	// plog.InfoStruct("tid", a)
}

type Data struct {
	*Target
	Link
}

type PtrToStr *string
type Target struct {
	ID      string
	Code    string
	ValueCn string   `stru:"enum_value_cn"`
	ValueEn PtrToStr `stru:"enum_value_en"`
}

type Link struct {
	ID string `json:"id"`
	// Schema    string   `json:"schema"`
	Table     string            `json:"table" zh:"table"`
	Column    []string          `json:"column" zh:"column"`
	Condition map[string]string `json:"condit" zh:"condit"`
	// sort
	// order
}

type Action func(*Data)

func Rule(p *Data) {
	if p.Target.ID == p.Link.ID {
		m := make(map[string]string, 0)
		m["code"] = "cloud_type"
		p.Link.Condition = m
		d := p.Link.Condition["code"]
		p.Code = d
	}
}

func Watch(d *Data, opts ...Action) error {
	for _, o := range opts {
		o(d)
	}
	typ := reflect.TypeOf(d)
	val := reflect.ValueOf(d)
	// 类型func 的方法 .kind() type返回指针 value返回function
	fmt.Println("typ.Kind():", typ.Kind())
	fmt.Println("val.Kind():", val.Kind())
	if typ.Kind() == reflect.Ptr {
		fmt.Println(fmt.Sprintf("形参i 类型%s 是指针", typ.Elem().Name()))
	}
	l := d.Link
	linktype := reflect.TypeOf(l)
	linkvalue := reflect.ValueOf(l)
	var tab string
	var colu []string
	var condi map[string]string

	for i := 0; i < linktype.NumField(); i++ {
		// fmt.Println("field name||", linktype.Field(i).Name)
		if linktype.Field(i).PkgPath != "" {
			fmt.Println("field pkg path||", linktype.Field(i).PkgPath)
		}
		// fmt.Println("field type||", linktype.Field(i).Type)
		// fmt.Println("field tag||", linktype.Field(i).Tag)
		if v, ok := linktype.Field(i).Tag.Lookup("zh"); ok {
			// fmt.Println("zh ==tag : value==||", v)
			if v == "table" {
				tab = linkvalue.Field(i).String()
			} else if v == "column" {
				// 如何用reflect 获取值的slice
				ss := linkvalue.Field(i).Slice(0, 1)
				for j := 0; j < ss.Len(); j++ {
					colu = append(colu, ss.Index(j).String())
				}
				// colu = d.Column
			} else if v == "condit" {
				condi = l.Condition
			}
		}
	}
	fmt.Println(tab)
	fmt.Println(colu)
	fmt.Println(condi)
	fmt.Println("===========================================")

	s := d.Target
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	// panic: reflect: call of reflect.Value.Elem on struct Value
	// Target 若不是指针，则不应使用elem（）。而是直接field 或numfield（）
	// 若是struct 是个拷贝，里面的类型或值理应是只读的
	for i := 0; i < v.Elem().NumField(); i++ {
		// fmt.Println("t.Elem():", t.Name())
		// fmt.Println("v.Elem():", .Type().Name())
		// type.Elem() struct Name
		fmt.Println("t.Elem()a:reflect type", t.Elem())
		fmt.Println("v.Elem()b:reflect value", v.Elem())
		// type Field StructField
		fmt.Println("Field(i)1 StructField:", t.Elem().Field(i))
		// value Field 对应的值的内容
		fmt.Println("Field(i)2 value:", v.Elem().Field(i))
		fmt.Println("==================")

		// struct name
		fmt.Println("Field(i)3: struct name", t.Elem().Field(i).Name)
		// struct data name 殊途同归
		fmt.Println("Field(i)4:", t.Elem().Field(i).Type.Name())
		// value Field 对应的值的内容 殊途同归 基础类型的指针没有名字
		fmt.Println("Field(i)5:", v.Elem().Field(i).Type().Name())

		// 殊途同归
		fmt.Println("Field(i).Type6:", t.Elem().Field(i).Type)
		fmt.Println("Field(i).Type7:", v.Elem().Field(i).Type())
		fmt.Println("==================")

		tag, ok := t.Elem().Field(i).Tag.Lookup("stru")
		fmt.Println("Field(i).Type8:", tag, ok)
		if ok {
			if v.Elem().Field(i).Kind() == reflect.Ptr {
				// fmt.Println(v.Elem().Field(i).Elem().Kind())
				if v.Elem().Field(i).IsNil() {
					fmt.Println(v.Elem().Field(i).CanAddr())
					// v.Elem().Field(i).Set(reflect.New(v.Elem().Field(i).Elem().Type()))
					v.Elem().Field(i).Set(reflect.New(t.Elem().Field(i).Type.Elem()))
					v.Elem().Field(i).Elem().SetString("123123123")
				}
			} else {
				v.Elem().Field(i).SetString("ueues")
			}
		}
		fmt.Println("===========================================")
	}
	fmt.Println(d.Target.ValueCn)
	fmt.Println(*d.Target.ValueEn)
	return nil
}
