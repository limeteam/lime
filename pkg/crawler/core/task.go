package core

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

// Task is a task define
type Task struct {
	ID int
	TaskRule
	TaskConfig
}

// NewTask return a new task object
func NewTask(id int, rule TaskRule, config TaskConfig) *Task {
	return &Task{
		ID:         id,
		TaskRule:   rule,
		TaskConfig: config,
	}
}

var (
	ctlMu  = &sync.RWMutex{}
	ctlMap = make(map[int]context.CancelFunc)
)

func addTaskCtrl(taskID int, cancelFunc context.CancelFunc) error {
	ctlMu.Lock()
	defer ctlMu.Unlock()

	if _, ok := ctlMap[taskID]; ok {
		return errors.WithStack(fmt.Errorf("duplicate taskID:%d", taskID))
	}
	ctlMap[taskID] = cancelFunc

	return nil
}

// CancelTask cancel a task by taskID
func CancelTask(taskID int) bool {
	ctlMu.Lock()
	defer ctlMu.Unlock()

	cancel, ok := ctlMap[taskID]
	if !ok {
		return false
	}
	cancel()
	delete(ctlMap, taskID)

	return true
}
