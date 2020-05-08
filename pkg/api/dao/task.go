package dao

import (
	"lime/pkg/api/dto"
	"lime/pkg/api/model"
	"lime/pkg/api/utils"
	"lime/pkg/common/db"
	"github.com/jinzhu/gorm"
	"strconv"
)

type TaskDao struct {
}

func (u TaskDao) Get(id int) model.Task {
	var Task model.Task
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Task)
	return Task
}

func (u TaskDao) GetAll() []model.Task {
	var Task []model.Task
	db := db.GetGormDB()
	db.Model(&model.Task{}).Find(&Task)
	return Task
}

func (u TaskDao) List(listDto dto.TaskListDto) ([]model.Task, int64) {
	var Task []model.Task
	var total int64
	TimeType, err := strconv.Atoi(listDto.Time_type)
	db := db.GetGormDB()
	if err == nil && TimeType == 1 && utils.IsNilObject(listDto.Start_time) && utils.IsNilObject(listDto.End_time) {
		db = db.Where("created_at >=?", listDto.Start_time)
		db = db.Where("created_at <=?", listDto.End_time)
	}
	status, err := strconv.Atoi(listDto.Status)
	if err == nil && status > 0 {
		db = db.Where("status =?", listDto.Status)
	}
	db.Offset(listDto.Skip).Limit(listDto.Limit).Find(&Task)
	db.Model(&model.Task{}).Count(&total)
	return Task, total
}

func (u TaskDao) Create(Task *model.Task) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Task)
}

func (u TaskDao) Update(Task *model.Task) *gorm.DB {
	db := db.GetGormDB()
	return db.Save(Task)
}

func (u TaskDao) Delete(Task *model.Task) *gorm.DB {
	db := db.GetGormDB()
	return db.Delete(Task)
}
