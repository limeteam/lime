package controllers

import (
	"github.com/gin-gonic/gin"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/service"
	"lime/pkg/api/utils/e"
)

var taskService = service.TaskService{}

type TaskController struct {
	BaseController
}

func (C *TaskController) List(c *gin.Context) {
	var listDto dto.TaskListDto
	if C.BindAndValidate(c, &listDto) {
		data, total := taskService.List(listDto)
		C.Resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      listDto.Page,
			"page_size": listDto.Limit,
		})
	}
}

func (C *TaskController) Get(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if C.BindAndValidate(c, &gDto) {
		data := taskService.InfoOfId(gDto)
		if data.Id < 1 {
			C.Fail(c, e.ErrIdData)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *TaskController) Create(c *gin.Context) {
	var taskDto dto.TaskCreateDto
	if C.BindAndValidate(c, &taskDto) {
		taskModel, err := taskService.Create(taskDto)
		if err != nil {
			C.Fail(c, e.ErrAddFail)
		}
		C.Resp(c, map[string]interface{}{
			"id": taskModel.Id,
		})
	}
}

func (C *TaskController) Delete(c *gin.Context) {
	var taskDto dto.GeneralDelDto
	if C.BindAndValidate(c, &taskDto) {
		affected := taskService.Delete(taskDto)
		if affected <= 0 {
			C.Fail(c, e.ErrDelFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": "ok.DeletedDone",
			"id":     taskDto.Id,
		})
	}
}

func (C *TaskController) Edit(c *gin.Context) {
	var taskDto dto.TaskEditDto
	if C.BindAndValidate(c, &taskDto) {
		err := taskService.Update(taskDto)
		if err != nil {
			C.Fail(c, e.ErrEditFail)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": "ok.UpdateDone",
			"id":     taskDto.Id,
		})
	}
}

func (C *TaskController) StopTask(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if C.BindAndValidate(c, &gDto) {
		data := taskService.StopTask(gDto)
		if data != nil {
			c.AbortWithError(400, data)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}

}

func (C *TaskController) StartTask(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if C.BindAndValidate(c, &gDto) {
		data := taskService.RunTask(gDto)
		if data != nil {
			c.AbortWithError(400, data)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (C *TaskController) RestartTask(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if C.BindAndValidate(c, &gDto) {
		data := taskService.RestartTask(gDto)
		if data != nil {
			c.AbortWithError(400, data)
			return
		}
		C.Resp(c, map[string]interface{}{
			"result": data,
		})
	}
}
