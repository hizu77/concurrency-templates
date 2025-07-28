package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/hizu77/concurrency-templates/internal/infra/semaphore"
)

const limit = 2

func main() {
	sema := semaphore.New(limit)

	for i := range 10 {
		sema.Acquire()
		go func() {
			defer sema.Release()
			fmt.Println(strconv.Itoa(i) + ":" + strconv.Itoa(runtime.NumGoroutine()))
		}()
	}
}
