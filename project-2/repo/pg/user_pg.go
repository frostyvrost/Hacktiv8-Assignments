package pg

import (
	"database/sql"
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/repo"
)

type userRepositoryImpl struct {
	db *sql.DB
}

const (
	createUserQuery = `
		INSERT INTO 
			users (username, email, age, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id, username, email, age
	`

	fetchUserByEmail = `
		SELECT
			id, 
			username, 
			email, 
			password, 
			age, 
			created_at, 
			updated_at
		FROM
			users
		WHERE
			email = $1
	`

	fetchUserById = `
		SELECT
			id, 
			username, 
			email, 
			password, 
			age, 
			created_at, 
			updated_at
		FROM
			users
		WHERE
			id = $1
	`

	updateUserQuery = `
		UPDATE 
			users
		SET
			username= $2,
			email= $3,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, username, email, age, updated_at
	`

	deleteUserQuery = `
		DELETE
		FROM
			users
		WHERE
			id = $1
	`
)

func NewUserRepository(db *sql.DB) repo.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// Create implements repo.UserRepository.
func (userRepo *userRepositoryImpl) Create(userPayload *models.User) (*dto.UserResponse, pkg.Error) {
	tx, err := userRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	var user dto.UserResponse
	err = tx.QueryRow(
		createUserQuery,
		userPayload.Username,
		userPayload.Email,
		userPayload.Age,
		userPayload.Password,
	).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Age,
	)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	return &user, nil
}

// Fetch implements repo.UserRepository.
func (userRepo *userRepositoryImpl) Fetch(email string) (*models.User, pkg.Error) {

	user := models.User{}
	err := userRepo.db.QueryRow(fetchUserByEmail, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewNotFoundError("user not found")
		}
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	return &user, nil
}

// Update implements repo.UserRepository.
func (userRepo *userRepositoryImpl) Update(userPayload *models.User) (*dto.UserUpdateResponse, pkg.Error) {

	tx, err := userRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wron")
	}

	row := tx.QueryRow(updateUserQuery, userPayload.Id, userPayload.Username, userPayload.Email)

	var user dto.UserUpdateResponse
	err = row.Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Age,
		&user.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wro")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("something went wro")
	}

	return &user, nil
}

// FetchById implements repo.UserRepository.
func (userRepo *userRepositoryImpl) FetchById(userId int) (*models.User, pkg.Error) {

	user := models.User{}
	err := userRepo.db.QueryRow(fetchUserById, userId).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewNotFoundError("user not found")
		}
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	return &user, nil
}

// Delete implements repo.UserRepository.
func (userRepo *userRepositoryImpl) Delete(userId int) pkg.Error {
	tx, err := userRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("Something went wrong!")
	}

	_, err = tx.Exec(deleteUserQuery, userId)

	if err != nil {
		return pkg.NewInternalServerError("Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		return pkg.NewInternalServerError("Something went wrong!")
	}

	return nil
}
