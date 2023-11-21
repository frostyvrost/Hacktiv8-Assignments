package pg

import (
	"database/sql"
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
	"project-3/repo"
)

const (
	createTask = `
		INSERT INTO tasks (
			user_id,
			title, 
			description,
			category_id
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id, title, description, status, user_id, category_id, created_at;
	`

	getTaskWithUser = `
		SELECT
			t.id,
			t.title,
			t.status,
			t.description,
			t.user_id,
			t.category_id,
			t.created_at,
			u.id,
			u.email,
			u.full_name
		FROM
			tasks AS t
		LEFT JOIN
			users AS u
		ON
			t.user_id = u.id
		ORDER BY
			t.id
		ASC
	`

	getTaskById = `
		SELECT 
			t.id,
			t.user_id
		FROM 
			tasks AS t
		WHERE 
			t.id = $1
	`

	updateTaskById = `
		UPDATE
			tasks
		SET
			title = $2,
			description = $3,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, title, description, status, user_id, category_id, updated_at
	`

	updateTaskByStatus = `
		UPDATE
			tasks
		SET
			status = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, title, description, status, user_id, category_id, updated_at
	`
	updateTaskByCategoryId = `
		UPDATE
			tasks
		SET
			category_id = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, title, description, status, user_id, category_id, updated_at
	`

	deleteTaskById = `
		DELETE
		FROM
			tasks
		WHERE
			id = $1
	`
)

type taskPG struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) repo.TaskRepository {
	return &taskPG{
		db: db,
	}
}


func (t *taskPG) CreateNewTask(taskPayLoad *models.Task) (*dto.NewTasksResponse, pkg.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	var task dto.NewTasksResponse

	row := tx.QueryRow(
		createTask,
		taskPayLoad.UserId,
		taskPayLoad.Title,
		taskPayLoad.Description,
		taskPayLoad.CategoryId,
	)
	err = row.Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.UserId,
		&task.CategoryId,
		&task.CreatedAt,
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

	return &task, nil
}

func (t *taskPG) GetTask() ([]repo.TaskUserMapped, pkg.MessageErr) {
	tasksUser := []repo.TaskUser{}
	rows, err := t.db.Query(getTaskWithUser)

	if err != nil {
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		var taskUser repo.TaskUser

		err := rows.Scan(
			&taskUser.Task.Id,
			&taskUser.Task.Title,
			&taskUser.Task.Status,
			&taskUser.Task.Description,
			&taskUser.Task.UserId,
			&taskUser.Task.CategoryId,
			&taskUser.Task.CreatedAt,
			&taskUser.User.Id,
			&taskUser.User.Email,
			&taskUser.User.FullName,
		)

		if err != nil {
			return nil, pkg.NewInternalServerError("something went wrong")
		}

		tasksUser = append(tasksUser, taskUser)
	}

	result := repo.TaskUserMapped{}
	return result.HandleMappingTasksUser(tasksUser), nil
}

func (t *taskPG) GetTaskById(id int) (*models.Task, pkg.MessageErr) {
	var taskUser models.Task

	err := t.db.QueryRow(getTaskById, id).Scan(&taskUser.Id, &taskUser.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewNotFoundError("task not found")
		}
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	return &taskUser, nil
}

func (t *taskPG) UpdateTaskById(taskPayLoad *models.Task) (*dto.UpdateTaskResponse, pkg.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateTaskById, taskPayLoad.Id, taskPayLoad.Title, taskPayLoad.Description)

	var taskUpdate dto.UpdateTaskResponse
	err = row.Scan(
		&taskUpdate.Id,
		&taskUpdate.Title,
		&taskUpdate.Description,
		&taskUpdate.Status,
		&taskUpdate.UserId,
		&taskUpdate.CategoryId,
		&taskUpdate.UpdatedAt,
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

	return &taskUpdate, nil
}

func (t *taskPG) UpdateTaskByStatus(taskPayLoad *models.Task) (*dto.UpdateTaskResponseByStatus, pkg.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateTaskByStatus, taskPayLoad.Id, taskPayLoad.Status)

	var taskUpdate dto.UpdateTaskResponseByStatus
	err = row.Scan(
		&taskUpdate.Id,
		&taskUpdate.Title,
		&taskUpdate.Description,
		&taskUpdate.Status,
		&taskUpdate.UserId,
		&taskUpdate.CategoryId,
		&taskUpdate.UpdatedAt,
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

	return &taskUpdate, nil
}

func (t *taskPG) UpdateTaskByCategoryId(taskPayLoad *models.Task) (*dto.UpdateCategoryIdResponse, pkg.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wrong")
	}
	var taskUpdate dto.UpdateCategoryIdResponse
	err = tx.QueryRow(updateTaskByCategoryId, taskPayLoad.Id, taskPayLoad.CategoryId).Scan(
		&taskUpdate.Id,
		&taskUpdate.Title,
		&taskUpdate.Description,
		&taskUpdate.Status,
		&taskUpdate.UserId,
		&taskUpdate.CategoryId,
		&taskUpdate.UpdatedAt,
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

	return &taskUpdate, nil
}

func (t *taskPG) DeleteTaskById(taskId int) pkg.MessageErr {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deleteTaskById, taskId)

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
