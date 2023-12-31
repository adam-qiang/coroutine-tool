<p align="center">
<a href="https://pkg.go.dev/github.com/adam-qiang/coroutine-tool"><img src="https://pkg.go.dev/badge/github.com/adam-qiang/coroutine-tool.svg" alt="Go Reference"></a>
<a href="https://en.wikipedia.org/wiki/MIT_License" rel="nofollow"><img alt="MIT" src="https://img.shields.io/badge/license-MIT-blue.svg" style="max-width:100%;"></a>
</p>

---

# coroutine-tool

利用GO语言的原生协程并发处理任务

## 安装

```go
 go get github.com/adam-qiang/coroutine-tool@latest
```

## Demo

```go
package main

import (
	"coroutine-github.com/adam-qiang/coroutine-tool/pool"
	"coroutine-github.com/adam-qiang/coroutine-tool/task"
	"time"
)

func main() {
	//创建一个任务
	task1 := task.CreateTask(func() error {
		time.Sleep(time.Second * 1)
		return nil
	})
	task2 := task.CreateTask(func() error {
		time.Sleep(time.Second * 10)
		return nil
	})

	task3 := task.CreateTask(func() error {
		return nil
	})

	//创建一个协程池
	createPool := pool.CreatePool(4)

	//向协程池中放入任务
	go func() {
		for i := 0; i < 50; i++ {
			createPool.EntryChannel <- task1
			createPool.EntryChannel <- task2
			createPool.EntryChannel <- task3
		}

	}()

	//执行线程池
	createPool.Run(3*50, true, "")
}

```