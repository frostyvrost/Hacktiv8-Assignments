package pg

import (
	"database/sql"
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/repo"
)

type photoRepositoryImpl struct {
	db *sql.DB
}

const (
	addNewPhotoQuery = `
		INSERT INTO
			photos
				(
					title,
					caption,
					photo_url,
					user_id
				)
		VALUES
				($2, $3, $4, $1)
		RETURNING
				id, title, caption, photo_url, user_id, created_at
	`

	getUserAndPhotos = `
				SELECT
					p.id,
					p.title,
					p.caption,
					p.photo_url,
					p.user_id,
					p.created_at,
					p.updated_at,
					u.email,
					u.username
				FROM
					photos as p
				LEFT JOIN
					users AS u
				ON
					p.user_id = u.id
				ORDER BY
					p.id
				ASC
	`
	getUserAndPhotosById = `
		SELECT
			p.id,
			p.title,
			p.caption,
			p.photo_url,
			p.user_id,
			p.created_at,
			p.updated_at,
			u.email,
			u.username
		FROM
			photos as p
		LEFT JOIN
			users AS u
		ON
			p.user_id = u.id
		WHERE
				p.id = $1
		ORDER BY
			p.id
		ASC
	`

	UpdatePhotoQuery = `
		UPDATE
			photos
		SET
			title = $2,
			caption = $3,
			photo_url = $4,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
				id, title, caption, photo_url, user_id, updated_at
	`

	deletePhotoById = `
		DELETE FROM
			photos
		WHERE
			id = $1		
	`
)

func NewPhotoRepository(db *sql.DB) repo.PhotoRepository {
	return &photoRepositoryImpl{
		db: db,
	}
}

// AddPhoto implements repo.PhotoRepository.
func (photoRepo *photoRepositoryImpl) AddPhoto(photoPayload *models.Photo) (*dto.PhotoResponse, pkg.Error) {

	tx, err := photoRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	row := tx.QueryRow(addNewPhotoQuery, photoPayload.UserId, photoPayload.Title, photoPayload.Caption, photoPayload.PhotoUrl)
	var photo dto.PhotoResponse

	err = row.Scan(
		&photo.Id,
		&photo.Title,
		&photo.Caption,
		&photo.PhotoUrl,
		&photo.UserId,
		&photo.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	return &photo, nil
}

// GetPhotos implements repo.PhotoRepository.
func (photoRepo *photoRepositoryImpl) GetPhotos() ([]repo.PhotoUserMapped, pkg.Error) {

	photosUser := []repo.PhotoUser{}
	rows, err := photoRepo.db.Query(getUserAndPhotos)

	if err != nil {
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	for rows.Next() {
		photoUser := repo.PhotoUser{}

		err = rows.Scan(
			&photoUser.Photo.Id,
			&photoUser.Photo.Title,
			&photoUser.Photo.Caption,
			&photoUser.Photo.PhotoUrl,
			&photoUser.Photo.UserId,
			&photoUser.Photo.CreatedAt,
			&photoUser.Photo.UpdatedAt,
			&photoUser.User.Email,
			&photoUser.User.Username,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, pkg.NewNotFoundError("photos not found")
			}
			return nil, pkg.NewInternalServerError("Something went wrong!")
		}

		photosUser = append(photosUser, photoUser)
	}

	result := repo.PhotoUserMapped{}
	return result.HandleMappingPhotoWithUser(photosUser), nil
}

// GetPhotoId implements repo.PhotoRepository.
func (photoRepo *photoRepositoryImpl) GetPhotoId(photoId int) (*repo.PhotoUserMapped, pkg.Error) {

	photoUser := repo.PhotoUser{}

	row := photoRepo.db.QueryRow(getUserAndPhotosById, photoId)
	err := row.Scan(
		&photoUser.Photo.Id,
		&photoUser.Photo.Title,
		&photoUser.Photo.Caption,
		&photoUser.Photo.PhotoUrl,
		&photoUser.Photo.UserId,
		&photoUser.Photo.CreatedAt,
		&photoUser.Photo.UpdatedAt,
		&photoUser.User.Email,
		&photoUser.User.Username,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewNotFoundError("photo not found")
		}
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	result := repo.PhotoUserMapped{}
	return result.HandleMappingPhotoWithUserByPhotoId(photoUser), nil
}

// UpdatePhoto implements repo.PhotoRepository.
func (photoRepo *photoRepositoryImpl) UpdatePhoto(photoId int, photoPayload *models.Photo) (*dto.PhotoUpdateResponse, pkg.Error) {
	tx, err := photoRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	row := tx.QueryRow(UpdatePhotoQuery, photoId, photoPayload.Title, photoPayload.Caption, photoPayload.PhotoUrl)

	var photo dto.PhotoUpdateResponse

	err = row.Scan(
		&photo.Id,
		&photo.Title,
		&photo.Caption,
		&photo.PhotoUrl,
		&photo.UserId,
		&photo.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	return &photo, nil
}

// DeletePhoto implements repo.PhotoRepository.
func (photoRepo *photoRepositoryImpl) DeletePhoto(photoId int) pkg.Error {
	tx, err := photoRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("Something went wrong!")
	}

	_, err = tx.Exec(deletePhotoById, photoId)

	if err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("Something went wrong!")
	}

	return nil
}
