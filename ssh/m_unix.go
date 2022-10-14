package theoryssh

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Unix() {
	// func (in io.Reader, out io.Write, args []string )
	// app1 param1 | app2 param2
	// pipe(bind(app1, param1), bind(app2, param2))
	args := os.Args[1:]
	for _, v := range args {
		fmt.Println(v)
	}
	p := pipe(bind(app1, args), app2)
	p(os.Stdin, os.Stdout)
}

func bind(
	app func(in io.Reader, out io.Writer, args []string),
	args []string,
) func(in io.Reader, out io.Writer) {
	return func(in io.Reader, out io.Writer) {
		app(in, out, args)
	}
}

func pipe(app1 func(in io.Reader, out io.Writer), app2 func(in io.Reader, out io.Writer)) func(in io.Reader, out io.Writer) {
	return func(in io.Reader, out io.Writer) {
		pr, pw := io.Pipe()
		defer pw.Close()
		go func() {
			defer pr.Close()
			fmt.Println("app2 堵塞....")
			app2(pr, out)
			fmt.Println("app2 完成....")
		}()
		app1(in, pw)
	}
}

func app1(in io.Reader, out io.Writer, args []string) {
	for _, v := range args {
		file, err := os.Open(v)
		if err != nil {
			continue
		}
		defer file.Close()
		buf := bufio.NewReader(file)
		for i := 1; ; i++ {
			line, err := buf.ReadBytes('\n')
			if err != nil {
				break
			}
			linenum := strconv.Itoa(i)
			nline := []byte(linenum + " ")
			fmt.Println("nline:", string(nline))
			fmt.Println("line:", string(line))
			nline = append(nline, line...)
			fmt.Println("nline:", string(nline))
			out.Write(nline)
		}
	}
}

func app2(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, 10)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			break
		}
		t := bytes.ToUpper(p[:n])
		fmt.Println("t:", string(t))
		out.Write(t)
	}
}
