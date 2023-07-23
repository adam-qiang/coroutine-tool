/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-23 00:12:25
 */

package pool

import (
	"coroutine-github.com/adam-qiang/coroutine-tool/progress_bar"
	"coroutine-github.com/adam-qiang/coroutine-tool/task"
	"sync"
)

type Pool struct {
	//对外的Task进入通道
	EntryChannel chan *task.Task

	//对内的Task通道，用来去消费Task
	JobChannel chan *task.Task

	//用来定义协程池内消费任务的worker数，即可理解为需要启动多少个协程来消费任务
	PoolWorkerNum int
}

var wg sync.WaitGroup
var taskPoolNum int
var runTaskNum int

// CreatePool 创建一个协程池
// @param poolWorkerNum int
// @return *Pool
func CreatePool(poolWorkerNum int) *Pool {
	//创建协程池
	pool := Pool{
		EntryChannel:  make(chan *task.Task),
		JobChannel:    make(chan *task.Task),
		PoolWorkerNum: poolWorkerNum,
	}

	return &pool
}

// Run 运行协程池
// @receiver p *Pool
// @param taskNum int
// @param isShowOption bool
// @param graph string
func (p *Pool) Run(taskNum int, isShowOption bool, graph string) {
	wg.Add(taskNum)
	taskPoolNum = taskNum
	runTaskNum = 0

	//根据协程池设置的worker数启动协程
	for i := 1; i <= p.PoolWorkerNum; i++ {
		go p.worker()
	}

	var bar progress_bar.Bar
	if isShowOption {
		if graph == "" {
			bar.NewBar(0, int64(taskNum))
		} else {
			bar.NewBarWithGraph(0, int64(taskNum), graph)
		}

	}

	num := 1

	//将外部进入的任务放入到协程池内通道
	for t := range p.EntryChannel {
		p.JobChannel <- t

		//运行进度条
		if isShowOption {
			bar.Run(int64(num))
		}

		if num == taskNum {
			//等待任务结束
			wg.Wait()

			//关闭
			close(p.EntryChannel)
			close(p.JobChannel)
		}
		num++
	}
}

// 协程消费
// @receiver p *Pool
func (p *Pool) worker() {
	for t := range p.JobChannel {
		if runTaskNum == taskPoolNum {
			return
		}

		runTaskNum++
		t.ExecuteTask()
		wg.Done()
	}
}
