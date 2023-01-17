package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		for {
			fmt.Printf("main thread id = %d\n", unix.Gettid())
			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		go doing()
		time.Sleep(220 * time.Millisecond)
	}
	wait := make(chan struct{})
	<-wait
}

func doing() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	exec.Command("bash", "-c", "sleep 3").Output()
}
