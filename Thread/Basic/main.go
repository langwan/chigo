package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os/exec"
	"sync"
)

func main() {

	//debug.SetMaxThreads(2)
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
	tid := unix.Gettid()
	fmt.Printf("task id = %d, thread id = %d\n", id, tid)
	exec.Command("bash", "-c", "sleep 3").Output()
	wg.Done()
}
