package controllers

import (
	"lime/pkg/api/dto"
	"lime/pkg/api/service"
	"github.com/gin-gonic/gin"
)

var taskService = service.TaskService{}

type TaskController struct {
	BaseController
}

func (t *TaskController) List(c *gin.Context) {
	var listDto dto.TaskListDto
	if t.BindAndValidate(c, &listDto) {
		data, total := taskService.List(listDto)
		resp(c, map[string]interface{}{
			"list":      data,
			"total":     total,
			"page":      listDto.Page,
			"page_size": listDto.Limit,
		})
	}
}

func (t *TaskController) Get(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if t.BindAndValidate(c, &gDto) {
		data := taskService.InfoOfId(gDto)
		if data.Id < 1 {
			fail(c, ErrIdData)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (t *TaskController) Create(c *gin.Context) {
	var taskDto dto.TaskCreateDto
	if t.BindAndValidate(c, &taskDto) {
		taskModel, err := taskService.Create(taskDto)
		if err != nil {
			fail(c, ErrAddFail)
		}
		resp(c, map[string]interface{}{
			"id": taskModel.Id,
		})
	}
}

func (t *TaskController) Delete(c *gin.Context) {
	var taskDto dto.GeneralDelDto
	if t.BindAndValidate(c, &taskDto) {
		affected := taskService.Delete(taskDto)
		if affected <= 0 {
			fail(c, ErrDelFail)
			return
		}
		resp(c, map[string]interface{}{
			"result": "ok.DeletedDone",
			"id":     taskDto.Id,
		})
	}
}

func (t *TaskController) Edit(c *gin.Context) {
	var taskDto dto.TaskEditDto
	if t.BindAndValidate(c, &taskDto) {
		err := taskService.Update(taskDto)
		if err != nil {
			fail(c, ErrEditFail)
			return
		}
		resp(c, map[string]interface{}{
			"result": "ok.UpdateDone",
			"id":     taskDto.Id,
		})
	}
}

func (t *TaskController) StopTask(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if t.BindAndValidate(c, &gDto) {
		data := taskService.StopTask(gDto)
		if data != nil {
			c.AbortWithError(400, data)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}

}

func (t *TaskController) StartTask(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if t.BindAndValidate(c, &gDto) {
		data := taskService.RunTask(gDto)
		if data != nil {
			c.AbortWithError(400, data)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}

func (t *TaskController) RestartTask(c *gin.Context) {
	var gDto dto.GeneralGetDto
	if t.BindAndValidate(c, &gDto) {
		data := taskService.RestartTask(gDto)
		if data != nil {
			c.AbortWithError(400, data)
			return
		}
		resp(c, map[string]interface{}{
			"result": data,
		})
	}
}
