package service

import (
	"lime/pkg/api/dao"
	"lime/pkg/api/domain/task"
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"lime/pkg/common"
	"lime/pkg/crawler/core"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var TaskDao = dao.TaskDao{}

// Service
type TaskService struct{}

// InfoOfId
func (us TaskService) InfoOfId(dto dto.GeneralGetDto) model.Task {
	return TaskDao.Get(dto.Id)
}

func (us TaskService) GetAll() []model.Task {
	return TaskDao.GetAll()
}

// List
func (us TaskService) List(dto dto.TaskListDto) ([]model.Task, int64) {
	return TaskDao.List(dto)
}

// Create
func (us TaskService) Create(dto dto.TaskCreateDto) (model.Task, error) {
	TaskModel := model.Task{
		TaskName:          dto.TaskName,
		TaskRuleName:      dto.TaskRuleName,
		TaskDesc:          dto.TaskDesc,
		Status:            common.TaskStatusRunning,
		Counts:            1,
		CronSpec:          dto.CronSpec,
		OutputType:        "book",
		OptUserAgent:      dto.OptUserAgent,
		OptMaxDepth:       dto.OptMaxDepth,
		OptAllowedDomains: dto.OptAllowedDomains,
		OptURLFilters:     dto.OptURLFilters,
		OptMaxBodySize:    dto.OptMaxBodySize,
		OptRequestTimeout: dto.OptRequestTimeout,
		AutoMigrate:       dto.AutoMigrate,
		LimitEnable:       dto.LimitEnable,
		LimitDomainRegexp: dto.LimitDomainRegexp,
		LimitDomainGlob:   dto.LimitDomainGlob,
		LimitDelay:        dto.LimitDelay,
		LimitRandomDelay:  dto.LimitRandomDelay,
		LimitParallelism:  dto.LimitParallelism,
		ProxyURLs:         dto.ProxyURLs,
	}
	c := TaskDao.Create(&TaskModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	spiderTask, err := task.GetSpiderTaskByModel(TaskModel)
	if err != nil {
		err := fmt.Sprintf("spider get model task failed! err:%+v", err)
		return TaskModel, errors.New(err)
	}
	s := core.New(spiderTask, task.GetMTSChan())
	if err := s.Run(); err != nil {
		err := fmt.Sprintf("spider run task failed! err:%+v", err)
		return TaskModel, errors.New(err)
	}

	if TaskModel.CronSpec != "" {
		log.Infof("starting cron task:%s", TaskModel.CronSpec)
		ct, err := task.NewCronTask(TaskModel.Id, TaskModel.CronSpec, task.GetMTSChan())
		if err != nil {
			err := fmt.Sprintf("new cron task failed! err:%+v", err)
			return TaskModel, errors.New(err)
		} else {
			if err := ct.Start(); err != nil {
				err := fmt.Sprintf("start cron task failed! err:%+v", err)
				return TaskModel, errors.New(err)
			}
		}
	}
	return TaskModel, nil
}

// Update
func (us TaskService) Update(dto dto.TaskEditDto) error {
	taskID := dto.Id
	if err := task.CheckTaskRuning(taskID); err != nil {
		return err
	}
	TaskModel := TaskDao.Get(taskID)

	// check task status
	if err := task.CheckTaskCanBeUpdate(taskID, TaskModel); err != nil {
		return err
	}
	// recreate crontab task if cronspec is change
	if err := task.CronTaskStopAndCreate(taskID, TaskModel, dto.CronSpec); err != nil {
		errs := fmt.Sprintf("UpdateTaskReq recreate crontab task fail, taskID: %v , err: %+v", taskID, err)
		return errors.New(errs)
	}

	TaskModel.TaskName = dto.TaskName
	TaskModel.TaskRuleName = dto.TaskRuleName
	TaskModel.TaskDesc = dto.TaskDesc
	TaskModel.Status = common.TaskStatusRunning
	TaskModel.Counts = TaskModel.Counts + 1
	TaskModel.CronSpec = dto.CronSpec
	TaskModel.OutputType = "book"
	TaskModel.OptUserAgent = dto.OptUserAgent
	TaskModel.OptMaxDepth = dto.OptMaxDepth
	TaskModel.OptAllowedDomains = dto.OptAllowedDomains
	TaskModel.OptURLFilters = dto.OptURLFilters
	TaskModel.OptMaxBodySize = dto.OptMaxBodySize
	TaskModel.OptRequestTimeout = dto.OptRequestTimeout
	TaskModel.AutoMigrate = dto.AutoMigrate
	TaskModel.LimitEnable = dto.LimitEnable
	TaskModel.LimitDomainRegexp = dto.LimitDomainRegexp
	TaskModel.LimitDomainGlob = dto.LimitDomainGlob
	TaskModel.LimitDelay = dto.LimitDelay
	TaskModel.LimitRandomDelay = dto.LimitRandomDelay
	TaskModel.LimitParallelism = dto.LimitParallelism
	TaskModel.ProxyURLs = dto.ProxyURLs
	TaskDao.Update(&TaskModel)
	return nil
}

// Delete
func (us TaskService) Delete(dto dto.GeneralDelDto) int64 {
	TaskModel := model.Task{
		Id: dto.Id,
	}
	c := TaskDao.Delete(&TaskModel)
	return c.RowsAffected
}

//执行非计划任务
func (us TaskService) RunTask(dto dto.GeneralGetDto) error {
	taskID := dto.Id
	if err := task.CheckTaskRuning(taskID); err != nil {
		return err
	}
	// query task info from db
	taskModel := TaskDao.Get(taskID)
	if taskModel.Id < 1 {
		return errors.New("StartTaskReq query model task fail")
	}
	// not allow crontab task
	if err := task.CheckCronSpec(taskModel.Id, taskModel); err != nil {
		return err
	}
	// check task status
	if err := task.CheckTaskStatus(taskModel.Id, taskModel); err != nil {
		return err
	}
	// create & run Task
	if err := task.CreateTaskModel(taskModel.Id, taskModel); err != nil {
		return err
	}
	// update task status
	taskModel.Status = common.TaskStatusRunning
	affected := TaskDao.Update(&taskModel)
	if affected.Error != nil {
		core.CancelTask(taskID)
		return errors.New("StartTaskReq update task status")
	}
	return nil
}

//停止非计划任务
func (us TaskService) StopTask(dto dto.GeneralGetDto) error {
	taskID := dto.Id
	if err := task.CheckTaskRuning(taskID); err != nil {
		return err
	}
	// stop cron task
	if ct := task.GetCronTask(taskID); ct != nil {
		ct.Stop()
	}
	// cancel spider task
	core.CancelTask(taskID)
	task := TaskDao.Get(taskID)
	if task.Id < 1 {
		return errors.New("StartTaskReq query model task fail")
	}
	task.Status = common.TaskStatusStopped
	affected := TaskDao.Update(&task)
	if affected.Error != nil {
		return errors.New("StartTaskReq update task status")
	}
	return nil
}

//重启任务
func (us TaskService) RestartTask(dto dto.GeneralGetDto) error {
	taskID := dto.Id
	if err := task.CheckTaskRuning(taskID); err != nil {
		return err
	}
	// query task info from db
	taskModel := TaskDao.Get(taskID)
	if taskModel.Id < 1 {
		return errors.New("StartTaskReq query model task fail")
	}
	// not allow crontab task
	if err := task.CheckCronSpecForRestart(taskModel.Id, taskModel); err != nil {
		return err
	}
	// check task status
	if err := task.CheckRestartTaskStatus(taskModel.Id, taskModel); err != nil {
		return err
	}
	// create crontab task
	err := task.CreateCronTask(taskID, taskModel.CronSpec)
	if err != nil {
		errs := fmt.Sprintf("RestartTaskReq run task fail, taskID: %v", taskID)
		return errors.New(errs)
	}
	// update task status
	taskModel.Status = common.TaskStatusCompleted
	affected := TaskDao.Update(&taskModel)
	if affected.Error != nil {
		return errors.New("RestartTaskReq update task status")
	}
	return nil
}

//用户调试
func (us TaskService) ExecTask(dto dto.GeneralGetDto) error {
	//taskID := dto.Id
	//if err := task.CheckTaskRuning(taskID); err != nil{
	//	return err
	//}
	//spiderTask, err := task.GetSpiderTaskByModel(task)
	//if err != nil {
	//	log.Errorf("StartTaskReq get model task fail, taskID: %v", taskID)
	//	errs := fmt.Sprintf("StartTaskReq get model task fail, taskID: %v", taskID)
	//	return errors.New(errs)
	//}
	//// run Task Model
	//s := core.New(spiderTask, task.GetMTSChan())
	//log.Info("start run")
	//if err := s.Run(); err != nil {
	//	log.Errorf("StartTaskReq run task fail, taskID: %v , err: %+v", taskID, err)
	//	errs := fmt.Sprintf("StartTaskReq run task fail, taskID: %v , err: %+v", taskID, err)
	//	return errors.New(errs)
	//}
	return nil
}

// check task status
// 1. set status to TaskStatusRunning if status is TaskStatusUnexceptedExited
// 2. restart task if status is TaskStatusCompleted, TaskStatusRunning
func CheckTask() {
	log.Infof("starting check task goroutine")
	tasks := TaskDao.GetAll()
	if tasks == nil {
		log.Infof("no task found, exit service.CheckTask method")
		return
	}
	for _, taskModel := range tasks {
		if taskModel.Status == common.TaskStatusRunning {
			log.Infof("set task status to TaskStatusUnexceptedExited, taskID:%v", taskModel.Id)
			taskModel.Status = common.TaskStatusUnexceptedExited
			affected := TaskDao.Update(&taskModel)
			if affected.Error != nil {
				log.Error("update task status")
				continue
			}
		}
		if (taskModel.Status == common.TaskStatusCompleted || taskModel.Status == common.TaskStatusRunning ||
			taskModel.Status == common.TaskStatusUnexceptedExited) && taskModel.CronSpec != "" {
			if err := task.CreateCronTask(taskModel.Id, taskModel.CronSpec); err != nil {
				log.Errorf("restart task err:%+v", err)
				continue
			}
		}
	}
}

func ManageTaskStatus() {
	log.Infof("starting manage task status goroutine")
	for {
		select {
		case mts := <-task.MtsCh:
			taskModel := TaskDao.Get(mts.ID)
			if taskModel.Id < 1 {
				log.Errorf("query model task err: %+v", taskModel)
				break
			}
			taskModel.Status = mts.Status
			if mts.Status == common.TaskStatusCompleted {
				taskModel.Counts += 1
			}
			affected := TaskDao.Update(&taskModel)
			if affected.Error != nil {
				log.Error("update task status")
				break
			}
		}
	}
}
