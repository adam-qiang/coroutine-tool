/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-23 00:47:41
 */

package task

// Task
// @Description: 任务结构体定义
type Task struct {
	function func() error
}

// CreateTask 创建任务方法（对外）
// @param f func() error
// @return *Task
func CreateTask(f func() error) *Task {
	newFunc := Task{
		function: f,
	}

	return &newFunc
}

// ExecuteTask 执行任务（对协程池内部的）
// @receiver t *Task
func (t *Task) ExecuteTask() {
	//执行任务中定义的任务内容
	err := t.function()
	if err != nil {
		return
	}
}
