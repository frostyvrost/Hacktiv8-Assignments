package service

import (
	"project-3/model"
	"project-3/pkg"
	"project-3/repo"

	"github.com/asaskevich/govalidator"
)

type taskServiceRepo interface {
	CreateTask(task *model.Task, userID uint) (*model.Task, pkg.Error)
	GetAllTasks() ([]*model.Task, pkg.Error)
	UpdateTask(task *model.TaskUpdate, taskId uint) (*model.Task, pkg.Error)
	UpdateStatusTask(statusupdate *model.TaskStatusUpdate, taskId uint) (*model.Task, pkg.Error)
	UpdateCategoryIdTask(Updatetask *model.TaskCategoryUpdate, taskId uint) (*model.Task, pkg.Error)
	DeleteTask(taskId uint) pkg.Error
}

type taskService struct{}

var TaskService taskServiceRepo = &taskService{}

func (t *taskService) CreateTask(task *model.Task, userID uint) (*model.Task, pkg.Error) {
	task.UserID = userID

	if _, err := govalidator.ValidateStruct(task); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	_, err := repo.CategoryModel.GetCategoryById(task.CategoryID)

	if err != nil {
		return nil, err
	}

	createdTask, err := repo.TaskModel.CreateTask(task)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (t *taskService) GetAllTasks() ([]*model.Task, pkg.Error) {
	var task []*model.Task

	task, err := repo.TaskModel.GetAllTasks()

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) UpdateTask(task *model.TaskUpdate, taskId uint) (*model.Task, pkg.Error) {
	if _, err := govalidator.ValidateStruct(task); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.TaskModel.UpdateTask(task, taskId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *taskService) UpdateStatusTask(statusupdate *model.TaskStatusUpdate, taskId uint) (*model.Task, pkg.Error) {
	if _, err := govalidator.ValidateStruct(statusupdate); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.TaskModel.UpdateStatusTask(statusupdate, taskId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *taskService) UpdateCategoryIdTask(Updatetask *model.TaskCategoryUpdate, taskId uint) (*model.Task, pkg.Error) {
	if _, err := govalidator.ValidateStruct(Updatetask); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	_, err := repo.CategoryModel.GetCategoryById(Updatetask.CategoryID)

	if err != nil {
		return nil, err
	}

	result, err := repo.TaskModel.UpdateCategoryIdTask(Updatetask, taskId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *taskService) DeleteTask(taskId uint) pkg.Error {
	err := repo.TaskModel.DeleteTask(taskId)

	if err != nil {
		return err
	}

	return nil
}
