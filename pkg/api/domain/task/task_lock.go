package task

import (
	"sync"
)

var taskLock *TaskLock

type TaskLock struct {
	taskLock map[int]bool
	sync.Mutex
}

func init() {
	taskLock = &TaskLock{
		taskLock: make(map[int]bool),
	}
}

func (tl *TaskLock) IsRunning(taskid int) bool {
	tl.Lock()
	defer tl.Unlock()
	if tl.taskLock[taskid] {
		return true
	}
	tl.taskLock[taskid] = true
	return false
}

func (tl *TaskLock) Complete(taskid int) {
	tl.Lock()
	defer tl.Unlock()
	delete(tl.taskLock, taskid)
}
