package repo

import (
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
	"time"
)

type user struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type TaskUser struct {
	Task models.Task
	User models.User
}

type TaskUserMapped struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"userId"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	Users       []user    `json:"user"`
}

func (tum *TaskUserMapped) HandleMappingTasksUser(taskUser []TaskUser) []TaskUserMapped {
	tasksUserMapped := make(map[int]TaskUserMapped)

	for _, eachTaskUser := range taskUser {
		taskId := eachTaskUser.Task.Id
		taskUserMapped, exists := tasksUserMapped[taskId]
		if !exists {
			taskUserMapped = TaskUserMapped{
				Id:          eachTaskUser.Task.Id,
				Title:       eachTaskUser.Task.Title,
				Description: eachTaskUser.Task.Description,
				Status:      eachTaskUser.Task.Status,
				UserId:      eachTaskUser.Task.UserId,
				CategoryId:  eachTaskUser.Task.CategoryId,
				CreatedAt:   eachTaskUser.Task.CreatedAt,
			}
		}

		user := user{
			Id:       eachTaskUser.User.Id,
			Email:    eachTaskUser.User.Email,
			FullName: eachTaskUser.User.FullName,
		}
		taskUserMapped.Users = append(taskUserMapped.Users, user)
		tasksUserMapped[taskId] = taskUserMapped
	}

	taskUsers := []TaskUserMapped{}
	for _, taskUser := range tasksUserMapped {
		taskUsers = append(taskUsers, taskUser)
	}
	return taskUsers
}

type TaskRepository interface {
	CreateNewTask(taskPayLoad *models.Task) (*dto.NewTasksResponse, pkg.MessageErr)
	GetTask() ([]TaskUserMapped, pkg.MessageErr)
	GetTaskById(id int) (*models.Task, pkg.MessageErr)
	UpdateTaskById(taskPayLoad *models.Task) (*dto.UpdateTaskResponse, pkg.MessageErr)
	UpdateTaskByStatus(taskPayLoad *models.Task) (*dto.UpdateTaskResponseByStatus, pkg.MessageErr)
	UpdateTaskByCategoryId(taskPayLoad *models.Task) (*dto.UpdateCategoryIdResponse, pkg.MessageErr)
	DeleteTaskById(taskId int) pkg.MessageErr
}
