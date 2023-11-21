package repo

import (
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
	"time"
)

type task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CategoryTask struct {
	Category models.Category
	Task     models.Task
}

type CategoryTaskMapped struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tasks     []task    `json:"Task"`
}

func (ctm *CategoryTaskMapped) HandleMappingCategoryWithTask(categoryTask []CategoryTask) []CategoryTaskMapped {
	categoryTasksMapped := make(map[int]CategoryTaskMapped)

	for _, eachCategoryTask := range categoryTask {
		categoryId := eachCategoryTask.Category.Id
		categoryTaskMapped, exists := categoryTasksMapped[categoryId]
		if !exists {
			categoryTaskMapped = CategoryTaskMapped{
				Id:        eachCategoryTask.Category.Id,
				Type:      eachCategoryTask.Category.Type,
				CreatedAt: eachCategoryTask.Category.CreatedAt,
				UpdatedAt: eachCategoryTask.Category.UpdatedAt,
			}
		}

		task := task{
			Id:          eachCategoryTask.Task.Id,
			Title:       eachCategoryTask.Task.Title,
			Description: eachCategoryTask.Task.Description,
			UserId:      eachCategoryTask.Task.UserId,
			CategoryId:  eachCategoryTask.Task.CategoryId,
			CreatedAt:   eachCategoryTask.Task.CreatedAt,
			UpdatedAt:   eachCategoryTask.Task.UpdatedAt,
		}
		categoryTaskMapped.Tasks = append(categoryTaskMapped.Tasks, task)
		categoryTasksMapped[categoryId] = categoryTaskMapped
	}

	categoryTasks := []CategoryTaskMapped{}
	for _, categoryTask := range categoryTasksMapped {
		categoryTasks = append(categoryTasks, categoryTask)
	}

	return categoryTasks
}

type CategoryRepository interface {
	Create(categoryPayLoad *models.Category) (*dto.NewCategoryResponse, pkg.MessageErr)
	GetCategory() ([]CategoryTaskMapped, pkg.MessageErr)
	UpdateCategory(categoryPayLoad *models.Category) (*dto.UpdateResponse, pkg.MessageErr)
	CheckCategoryId(categoryId int) (*models.Category, pkg.MessageErr)
	DeleteCategory(categoryId int) pkg.MessageErr
}
