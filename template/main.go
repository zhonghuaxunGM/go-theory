package main

import (
	"net/http"
	"os"
	"strings"
	"text/template"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	// t1, err := template.ParseFiles("tmpl.html")
	t1, err := template.New("test").Parse(
		`<!DOCTYPE html>
		<html>
			<head>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
				<title>Go Web</title>
			</head>
			<body>
				{{ . }} {{ . }}
			</body>
		</html>`)

	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func info(w http.ResponseWriter, r *http.Request) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"longshuai", 23}
	tmpl, _ := template.New("test").Parse(
		`Name: {{.Name}}, Age: {{.Age}}`)
	_ = tmpl.Execute(os.Stdout, p)
}

func forrange(w http.ResponseWriter, r *http.Request) {
	type Friend struct {
		Fname string
	}
	type Person struct {
		UserName string
		Emails   []string
		Friends  []*Friend
	}
	p := Person{UserName: "zhxu",
		Emails:  []string{"a1@qq.com", "a2@gmail.com"},
		Friends: []*Friend{{Fname: "xiaofang"}, {Fname: "wugui"}}}

	funcMap := template.FuncMap{
		"doupper": upper,
	}
	// template.Must()
	// Must操作
	// 检测模板是否正确：大括号是否匹配，注释是否正确关闭，变量是否正确
	tmpl, _ := template.New("test2").Funcs(funcMap).Parse(
		`hello {{.UserName}}!
{{ $long := (len .UserName) }}
{{ println $long}}
{{ range .Emails }}
an email {{ . }}
{{ end }}
{{ with .Friends }}
{{ range . }}
{{ doupper .Fname}}
{{.Fname | printf "my friend name is %s" }}  
{{ end }}
{{ end }}`)
	tmpl.Execute(os.Stdout, p)
}

func upper(str string) string {
	return strings.ToUpper(str)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", tmpl)
	http.HandleFunc("/1", info)
	http.HandleFunc("/2", forrange)

	server.ListenAndServe()
}
