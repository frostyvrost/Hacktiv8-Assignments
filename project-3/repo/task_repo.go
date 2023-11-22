package repo

import (
	"project-3/database"
	"project-3/model"
	"project-3/pkg"
)

type taskModelRepo interface {
	CreateTask(task *model.Task) (*model.Task, pkg.Error)
	GetAllTasks() ([]*model.Task, pkg.Error)
	UpdateTask(input *model.TaskUpdate, taskID uint) (*model.Task, pkg.Error)
	UpdateStatusTask(taskStatus *model.TaskStatusUpdate, taskID uint) (*model.Task, pkg.Error)
	UpdateCategoryIdTask(input *model.TaskCategoryUpdate, taskID uint) (*model.Task, pkg.Error)

	DeleteTask(taskId uint) pkg.Error
}

type taskModel struct{}

var TaskModel taskModelRepo = &taskModel{}

func (t *taskModel) CreateTask(task *model.Task) (*model.Task, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&task).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return task, nil
}

func (t *taskModel) GetAllTasks() ([]*model.Task, pkg.Error) {
	db := database.GetDB()
	var tasks []*model.Task

	err := db.Preload("User").Find(&tasks).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return tasks, nil
}

func (t *taskModel) UpdateTask(input *model.TaskUpdate, taskID uint) (*model.Task, pkg.Error) {
	db := database.GetDB()

	var task model.Task
	err := db.First(&task, taskID).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&task).Updates(input)

	return &task, nil
}

func (t *taskModel) UpdateStatusTask(taskStatus *model.TaskStatusUpdate, taskID uint) (*model.Task, pkg.Error) {
	db := database.GetDB()

	var task model.Task

	err := db.First(&task, taskID).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&task).Updates(taskStatus)

	return &task, nil
}

func (t *taskModel) UpdateCategoryIdTask(input *model.TaskCategoryUpdate, taskID uint) (*model.Task, pkg.Error) {
	db := database.GetDB()

	var task model.Task

	err := db.First(&task, taskID).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&task).Updates(input)

	return &task, nil
}

func (t *taskModel) DeleteTask(taskId uint) pkg.Error {
	db := database.GetDB()

	var task model.Task

	err := db.Where("id = ?", taskId).Delete(&task).Error

	if err != nil {
		return pkg.ParseError(err)
	}

	return nil
}
