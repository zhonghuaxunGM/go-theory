package main

import (
	"fmt"
	"sync"
	"time"
)

func SelectCase() {

	go func() {
		for {
			fmt.Println(time.Now())
			<-time.After(time.Duration(2 * time.Second))
		}
	}()
	arr := []string{"Test1", "Test2"}
	msgCh := make(chan string, 5)
	// defer func() {
	// 	close(msgCh)
	// }()

	retCh := make(chan string)
	sucCh := make(chan string)

	go func() {
		// 当channel被关闭时，接收者的for循环也被自动停止了
		// for v := range msgCh {
		// 	fmt.Println("msgch recv:", v)
		// }
		defer func() {
			sucCh <- "suceess"
		}()
		var strbry_closed bool
		for {
			if strbry_closed {
				return
			}
			select {
			case s, ok := <-msgCh:
				// 判断channel是否关闭，便是接受结束
				if !ok {
					fmt.Println("strbry_closed:", strbry_closed)
					strbry_closed = true
				} else {
					fmt.Println("msgch recv:", s)
				}
				// 某一个协程无法正常开始工作
				// msgCh 不能在规定时间内开始工作
			case <-time.After(time.Duration(3 * time.Second)):
				fmt.Println("time out")
				retCh <- "return call chan"
				return
			}
		}
	}()

	wg := sync.WaitGroup{}
	for _, v := range arr {
		wg.Add(1)
		go func(v string) {
			defer func() {
				wg.Done()
			}()
			// <-time.After(time.Duration(6 * time.Second))
			fmt.Println("msgch send:", v)
			msgCh <- v
		}(v)
	}

	go func() {
		wg.Wait()
		// 发送完成 关闭ch
		close(msgCh)
	}()

	select {
	// 工作时间总量时间
	case <-time.After(time.Duration(100 * time.Second)):
		return
	case r := <-retCh:
		fmt.Println(r)
		return
	case r := <-sucCh:
		fmt.Println(r)
		return
	}
}
