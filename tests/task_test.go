package tests

import (
	"coroutine-github.com/adam-qiang/coroutine-tool/pool"
	"coroutine-github.com/adam-qiang/coroutine-tool/task"
	"testing"
	"time"
)

func TestRunTask(t *testing.T) {
	//创建一个任务
	task1 := task.CreateTask(func() error {
		//fmt.Println(time.Now())

		time.Sleep(time.Second * 1)
		return nil
	})
	task2 := task.CreateTask(func() error {
		//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		//time.Sleep(time.Second * 1)
		return nil
	})

	task3 := task.CreateTask(func() error {
		//fmt.Println(time.Now().Format("2006-01-02"))
		//time.Sleep(time.Second * 1)
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
