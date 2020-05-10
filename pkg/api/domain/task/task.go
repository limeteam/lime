package task

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"lime/pkg/api/model"
	"lime/pkg/common"
	"lime/pkg/crawler/core"
)

var MtsCh = make(chan common.MTS, 1)

func GetMTSChan() chan common.MTS {
	return MtsCh
}
func CheckTaskRuning(taskID int) error {
	if taskLock.IsRunning(taskID) {
		return errors.New("other operation is running")
	}
	defer taskLock.Complete(taskID)
	return nil
}

func CheckTaskStatus(taskID int, task model.Task) error {
	if !taskCanBeStart(task) {
		errs := fmt.Sprintf("StartTaskReq taskID status is non-conformance , taskID: %v", taskID)
		return errors.New(errs)
	}
	return nil
}

func CheckTaskCanBeUpdate(taskID int, task model.Task) error {
	if !taskCanBeUpdate(task) {
		errs := fmt.Sprintf("StartTaskReq taskID status is non-conformance , taskID: %v", taskID)
		return errors.New(errs)
	}
	return nil
}

func CheckRestartTaskStatus(taskID int, task model.Task) error {
	if !taskCanBeRestart(task) {
		errs := fmt.Sprintf("StartTaskReq taskID status is non-conformance , taskID: %v", taskID)
		return errors.New(errs)
	}
	return nil
}

func CheckCronSpec(taskID int, task model.Task) error {
	if task.CronSpec != "" {
		errs := fmt.Sprintf("RestartTaskReq taskID is crontab task, taskID: %d", taskID)
		return errors.New(errs)
	}
	return nil
}

func CheckCronSpecForRestart(taskID int, task model.Task) error {
	// only allow crontab task
	if task.CronSpec != "" {
		errs := fmt.Sprintf("RestartTaskReq taskID is crontab task, taskID: %d", taskID)
		return errors.New(errs)
	}
	return nil
}

func CreateTaskModel(taskID int, task model.Task) error {
	//create task
	spiderTask, err := GetSpiderTaskByModel(task)
	if err != nil {
		errs := fmt.Sprintf("StartTaskReq get model task fail, taskID: %v", taskID)
		return errors.New(errs)
	}
	// run Task Model
	s := core.New(spiderTask, GetMTSChan())
	if err := s.Run(); err != nil {
		errs := fmt.Sprintf("StartTaskReq run task fail, taskID: %v , err: %+v", taskID, err)
		return errors.New(errs)
	}
	return nil
}

func CreateCronTask(taskId int, taskCronSpec string) error {
	log.Infof("restarting task, taskID:%v", taskId)
	ct, err := NewCronTask(taskId, taskCronSpec, GetMTSChan())
	if err != nil {
		log.Errorf("new cron task failed! err:%+v", err)
		return err
	}

	if err := ct.Start(); err != nil {
		log.Errorf("start cron task failed! err:%+v", err)
		return err
	}
	return nil
}

func CronTaskStopAndCreate(taskID int, oldtask model.Task, newtaskCronSpec string) error {
	if oldtask.CronSpec == newtaskCronSpec {
		return nil
	}
	if oldtask.CronSpec != "" {
		// stop cron task
		if ct := GetCronTask(taskID); ct != nil {
			ct.Stop()
		}
	}
	if newtaskCronSpec != "" && (oldtask.Status == common.TaskStatusCompleted || oldtask.Status == common.TaskStatusUnexceptedExited) {
		return CreateCronTask(taskID, newtaskCronSpec)
	}
	return nil
}

func taskCanBeStart(task model.Task) bool {
	return task.Status == common.TaskStatusStopped ||
		task.Status == common.TaskStatusUnexceptedExited ||
		task.Status == common.TaskStatusCompleted ||
		task.Status == common.TaskStatusRunningTimeout
}

func taskCanBeRestart(task model.Task) bool {
	return task.Status == common.TaskStatusStopped
}

func taskCanBeUpdate(task model.Task) bool {
	return task.Status == common.TaskStatusStopped ||
		task.Status == common.TaskStatusUnexceptedExited ||
		task.Status == common.TaskStatusCompleted ||
		task.Status == common.TaskStatusRunningTimeout ||
		task.Status == common.TaskStatusRunning

}
