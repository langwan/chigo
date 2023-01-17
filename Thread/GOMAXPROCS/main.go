package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"sync"
)

// 单核和多核心执行
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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
