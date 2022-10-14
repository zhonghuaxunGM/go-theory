package trick

import (
	"bytes"
	"fmt"
	"os"
)

func makeBuff() {
	fmt.Println("===============case1=============")
	// 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
	// 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。
	var dou byte = '?'
	var ren rune = '人'
	var num int
	var prt *int
	num = 10
	prt = &num
	fmt.Println(prt, ren, dou)

	fmt.Println("\n\n\n\n===============case2=============buffer 创建")
	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))
	// buf21 := bytes.NewBuffer([]byte("hello", "hello"))
	buf3 := bytes.NewBuffer([]byte{'h', 'e', 'l', 'l', 'o'})
	fmt.Println(buf1, buf2, buf3)

	fmt.Println("\n\n\n\n===============case3=============Write 把slice byte放置尾部")
	world1 := []byte(" world")
	buf := bytes.NewBufferString("hello")
	buf.Write(world1)
	fmt.Println(buf.String())

	fmt.Println("\n\n\n\n===============case4=============WriteString 把字符放置buffer尾部")
	world2 := "world"
	buf = bytes.NewBufferString("hello")
	buf.WriteString(world2)
	fmt.Println(buf.String())

	fmt.Println("\n\n\n\n===============case4=============WriteByte 把btye放置buffer尾部")
	var world3 byte = '?'
	buf = bytes.NewBufferString("hello")
	buf.WriteByte(world3)
	fmt.Println(buf.String())

	fmt.Println("\n\n\n\n===============case5=============WriteTo方法，将一个缓冲器的数据写到w里，w是实现io.Writer的")
	os.Remove("./text.txt")
	file, _ := os.Create("./text.txt")
	buf = bytes.NewBufferString("hello world")
	fmt.Fprint(file, buf.String())
	buf = bytes.NewBufferString("!!!!")
	buf.WriteTo(file)

	fmt.Println("\n\n\n\n===============case6=============Read方法，返回一个容器p，读完后p就满了，缓冲器相应的减少。")
	buff := bytes.NewBuffer([]byte("test"))
	buff.Write([]byte("dsadh1:kihifailed=1kghyhjkdffyjh2:khjklhonefailed=3"))
	s3 := make([]byte, 2)
	buff.Read(s3)
	fmt.Println(buff.String())
	fmt.Println(string(s3))

	buff.Read(s3)
	fmt.Println(buff.String())
	fmt.Println(string(s3))

	fmt.Println("\n\n\n\n===============case7=============ReadString方法||ReadBytes方法，需要一个byte作为分隔符，读的时候从缓冲器里找出第一个出现的分隔符，缓冲器头部开始到分隔符之间的byte返回")
	buff = bytes.NewBuffer([]byte("test"))
	buff.Write([]byte(" done1one2"))
	newbuff, _ := buff.ReadBytes('s')
	fmt.Println(buff.String())
	fmt.Println(string(newbuff))

	fmt.Println("\n\n\n\n===============case8=============ReadFrom方法，从一个实现io.Reader接口的r，把r的内容读到缓冲器里，n返回读的数量")
	file, _ = os.Open("./text.txt")
	var buffff bytes.Buffer
	buffff.ReadFrom(file)
	fmt.Println(buffff.String())

}
