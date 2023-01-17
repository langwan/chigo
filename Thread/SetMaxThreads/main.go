package main

import (
	"fmt"
	"os/exec"
	"runtime/debug"
	"sync"
)

// 默认情况下最大线程数是 10000, go 语言不会主动释放空闲线程
func main() {
	debug.SetMaxThreads(100)
	taskCount := 50
	wg := sync.WaitGroup{}
	wg.Add(taskCount)
	for i := 0; i < taskCount; i++ {
		go doing(i, &wg)
	}
	wg.Wait()
	fmt.Println("ok.")
}

func doing(id int, wg *sync.WaitGroup) {
	fmt.Printf("task id = %d\n", id)
	exec.Command("bash", "-c", "sleep 3").Output()
	wg.Done()
}
