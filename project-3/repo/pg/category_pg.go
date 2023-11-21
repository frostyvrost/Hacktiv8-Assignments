package pg

import (
	"database/sql"
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
	"project-3/repo"
	"time"
)

type categoryWithTask struct {
	CategoryId        int
	CategoryType      string
	CategoryCreatedAt time.Time
	CategoryUpdatedAt time.Time
	TaskId            sql.NullInt64
	TaskTitle         sql.NullString
	TaskDescription   sql.NullString
	TaskStatus        sql.NullBool
	TaskUserId        sql.NullInt64
	TaskCategoryId    sql.NullInt64
	TaskCreatedAt     sql.NullTime
	TaskUpdatedAt     sql.NullTime
}

func (c *categoryWithTask) categoryWithTaskToModels() repo.CategoryTask {
	return repo.CategoryTask{
		Category: models.Category{
			Id:        c.CategoryId,
			Type:      c.CategoryType,
			CreatedAt: c.CategoryCreatedAt,
			UpdatedAt: c.CategoryUpdatedAt,
		},
		Task: models.Task{
			Id:          int(c.TaskId.Int64),
			Title:       c.TaskTitle.String,
			Description: c.TaskDescription.String,
			Status:      c.TaskStatus.Bool,
			UserId:      int(c.TaskUserId.Int64),
			CategoryId:  int(c.TaskCategoryId.Int64),
			CreatedAt:   c.TaskCreatedAt.Time,
			UpdatedAt:   c.TaskUpdatedAt.Time,
		},
	}
}

const (
	createCategory = `
		INSERT INTO "categories"
		(
			type
		)
		VALUES ($1)
		RETURNING
			id, type, created_at;
	`

	getCategoryWithTask = `
		SELECT
			c.id,
			c.type,
			c.updated_at,
			c.created_at,
			t.id,
			t.title,
			t.description,
			t.user_id,
			t.category_id,
			t.created_at,
			t.updated_at
		FROM
			categories AS c
		LEFT JOIN
			tasks AS t
		ON
			c.id = t.category_id
		ORDER BY
			c.id
		ASC
	`

	updateCategoryById = `
		UPDATE
			categories
		SET
			type = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, type, updated_at
	`

	checkCategoryId = `
		SELECT 
			c.id 
		FROM 
			categories AS c
		WHERE
			c.id = $1
	`

	deleteCategoryById = `
		DELETE
		FROM
			categories
		WHERE
			id = $1
	`
)

type categoryPG struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) repo.CategoryRepository {
	return &categoryPG{
		db: db,
	}
}

func (c *categoryPG) Create(categoryPayLoad *models.Category) (*dto.NewCategoryResponse, pkg.MessageErr) {
	tx, err := c.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	var category dto.NewCategoryResponse

	row := tx.QueryRow(createCategory, categoryPayLoad.Type)
	err = row.Scan(&category.Id, &category.Type, &category.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) GetCategory() ([]repo.CategoryTaskMapped, pkg.MessageErr) {
	categoryTasks := []repo.CategoryTask{}
	rows, err := c.db.Query(getCategoryWithTask)

	if err != nil {
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		categoryTask := categoryWithTask{}

		err := rows.Scan(
			&categoryTask.CategoryId,
			&categoryTask.CategoryType,
			&categoryTask.CategoryCreatedAt,
			&categoryTask.CategoryUpdatedAt,
			&categoryTask.TaskId,
			&categoryTask.TaskTitle,
			&categoryTask.TaskDescription,
			&categoryTask.TaskUserId,
			&categoryTask.TaskCategoryId,
			&categoryTask.TaskCreatedAt,
			&categoryTask.TaskUpdatedAt,
		)

		if err != nil {
			return nil, pkg.NewInternalServerError("something went wrong")
		}

		categoryTasks = append(categoryTasks, categoryTask.categoryWithTaskToModels())
	}

	result := repo.CategoryTaskMapped{}
	return result.HandleMappingCategoryWithTask(categoryTasks), nil
}

func (c *categoryPG) UpdateCategory(categoryPayLoad *models.Category) (*dto.UpdateResponse, pkg.MessageErr) {
	tx, err := c.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateCategoryById, categoryPayLoad.Id, categoryPayLoad.Type)

	var categoryUpdate dto.UpdateResponse
	err = row.Scan(
		&categoryUpdate.Id,
		&categoryUpdate.Type,
		&categoryUpdate.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	return &categoryUpdate, nil
}

func (c *categoryPG) CheckCategoryId(categoryId int) (*models.Category, pkg.MessageErr) {
	category := models.Category{}
	row := c.db.QueryRow(checkCategoryId, categoryId)
	err := row.Scan(&category.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewInternalServerError("category not found")
		}
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) DeleteCategory(categoryId int) pkg.MessageErr {
	tx, _ := c.db.Begin()

	_, err := tx.Exec(deleteCategoryById, categoryId)

	if err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("something went wrong")
	}

	return nil
}
