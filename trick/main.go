package trick

import (
	"fmt"
	"strings"
	"time"
)

func recuse(arr []string, str string) map[string]string {
	strs := strings.SplitAfter(str, "failed=")
	fmt.Println(strs)
	result := make(map[string]string, 0)
	for k, v := range strs {
		for _, vv := range arr {
			if strings.Contains(v, vv) {
				result[vv] = strs[k+1][:1]
				break
			}
		}
	}
	fmt.Println(result)
	return result
}

func trick() {
	// makeBuff()
	// recuse([]string{"h1", "h2", "h3"}, "dsadh1:kihifailed=1kghyhjkdffyjh2:khjklhonefailed=3")
	Pro()

	// deferCall()
	<-time.After(time.Second * 10)
}

func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	panic("触发panic")
}
