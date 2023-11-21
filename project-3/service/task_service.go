package service

import (
	"net/http"
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
	"project-3/repo"
)

type TaskService interface {
	Create(userId int, taskPayLoad *dto.NewTasksRequest) (*dto.NewTasksResponse, pkg.MessageErr)
	Get() (*dto.GetResponseTasks, pkg.MessageErr)
	UpdateTask(taskId int, taskPayLoad *dto.UpdateTaskRequest) (*dto.UpdateResponseTask, pkg.MessageErr)
	UpdateTaskByStatus(taskId int, taskPayLoad *dto.UpdateTaskRequestByStatus) (*dto.UpdateResponseTask, pkg.MessageErr)
	UpdateTaskByCategoryId(taskId int, taskPayLoad *dto.UpdateCategoryIdRequest) (*dto.UpdateCategoryId, pkg.MessageErr)
	DeleteTaskById(taskId int) (*dto.DeleteTaskByIdResponse, pkg.MessageErr)
}

type taskService struct {
	taskRepo     repo.TaskRepository
	categoryRepo repo.CategoryRepository
	userRepo     repo.UserRepository
}

func NewTaskService(taskRepo repo.TaskRepository, categoryRepo repo.CategoryRepository, userRepo repo.UserRepository) TaskService {
	return &taskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (ts *taskService) Create(userId int, taskPayLoad *dto.NewTasksRequest) (*dto.NewTasksResponse, pkg.MessageErr) {
	err := pkg.ValidateStruct(taskPayLoad)

	if err != nil {
		return nil, err
	}

	_, err = ts.categoryRepo.CheckCategoryId(taskPayLoad.CategoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	task := &models.Task{
		UserId:      userId,
		Title:       taskPayLoad.Title,
		Description: taskPayLoad.Description,
		CategoryId:  taskPayLoad.CategoryId,
	}
	response, err := ts.taskRepo.CreateNewTask(task)

	if err != nil {
		return nil, err
	}

	response = &dto.NewTasksResponse{
		Id:          response.Id,
		Title:       response.Title,
		Description: response.Description,
		Status:      response.Status,
		UserId:      response.UserId,
		CategoryId:  response.CategoryId,
		CreatedAt:   response.CreatedAt,
	}

	return response, nil
}

func (ts *taskService) Get() (*dto.GetResponseTasks, pkg.MessageErr) {
	tasks, err := ts.taskRepo.GetTask()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	response := dto.GetResponseTasks{
		StatusCode: http.StatusOK,
		Message:    "successfully fetching task",
		Data:       tasks,
	}

	return &response, nil
}

func (ts *taskService) UpdateTask(taskId int, taskPayLoad *dto.UpdateTaskRequest) (*dto.UpdateResponseTask, pkg.MessageErr) {
	err := pkg.ValidateStruct(taskPayLoad)

	if err != nil {
		return nil, err
	}

	updateTask, err := ts.taskRepo.GetTaskById(taskId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if updateTask.Id != taskId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	task := &models.Task{
		Id:          taskId,
		Title:       taskPayLoad.Title,
		Description: taskPayLoad.Description,
	}

	response, err := ts.taskRepo.UpdateTaskById(task)

	if err != nil {
		return nil, err
	}

	return &dto.UpdateResponseTask{
		StatusCode: http.StatusOK,
		Message:    "Task has been successfully updated",
		Data:       response,
	}, nil
}

func (ts *taskService) UpdateTaskByStatus(taskId int, taskPayLoad *dto.UpdateTaskRequestByStatus) (*dto.UpdateResponseTask, pkg.MessageErr) {
	err := pkg.ValidateStruct(taskPayLoad)

	if err != nil {
		return nil, err
	}

	updateTask, err := ts.taskRepo.GetTaskById(taskId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if updateTask.Id != taskId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	task := &models.Task{
		Id:     taskId,
		Status: taskPayLoad.Status,
	}

	response, err := ts.taskRepo.UpdateTaskByStatus(task)

	if err != nil {
		return nil, err
	}

	return &dto.UpdateResponseTask{
		StatusCode: http.StatusOK,
		Message:    "Status has been successfully updated",
		Data:       response,
	}, nil
}

func (ts *taskService) UpdateTaskByCategoryId(taskId int, taskPayLoad *dto.UpdateCategoryIdRequest) (*dto.UpdateCategoryId, pkg.MessageErr) {
	err := pkg.ValidateStruct(taskPayLoad)

	if err != nil {
		return nil, err
	}

	updateTask, err := ts.taskRepo.GetTaskById(taskId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if updateTask.Id != taskId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	_, err = ts.categoryRepo.CheckCategoryId(taskPayLoad.CategoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequest("invalid user")
		}
		return nil, err
	}

	task := &models.Task{
		Id:         taskId,
		CategoryId: taskPayLoad.CategoryId,
	}

	response, err := ts.taskRepo.UpdateTaskByCategoryId(task)

	if err != nil {
		return nil, err
	}

	return &dto.UpdateCategoryId{
		StatusCode: http.StatusOK,
		Message:    "Category id has been successfully updated",
		Data:       response,
	}, nil
}

func (ts *taskService) DeleteTaskById(taskId int) (*dto.DeleteTaskByIdResponse, pkg.MessageErr) {
	task, err := ts.taskRepo.GetTaskById(taskId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if task.Id != taskId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	ts.taskRepo.DeleteTaskById(taskId)

	response := &dto.DeleteTaskByIdResponse{
		Message: "Task has been successfully deleted",
	}

	return response, nil
}
