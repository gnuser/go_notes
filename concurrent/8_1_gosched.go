package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(1) // 只使用一个调度器
	exit := make(chan struct{})
	go func() {
		defer close(exit)

		// b任务不一定被执行, 修改GOMAXPROCS为0或大于1后也有几率不执行
		go func() {
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)
			if i == 1 {
				runtime.Gosched() // 将当前任务放回队列，等待下次调度
			}
		}
	}()

	<-exit
}
