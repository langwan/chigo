package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {

	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		for {
			threadCreate := pprof.Lookup("threadcreate")
			fmt.Printf("threads = %d\n", threadCreate.Count())
			time.Sleep(200 * time.Millisecond)
		}
	}()
	for i := 0; i < 100; i++ {
		go doing()
		time.Sleep(220 * time.Millisecond)
	}

	wait := make(chan struct{})
	<-wait
}

func doing() {
	runtime.LockOSThread()
	//defer runtime.UnlockOSThread()
	exec.Command("bash", "-c", "sleep 3").Output()
}
