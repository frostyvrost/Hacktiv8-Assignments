package pg

import (
	"database/sql"
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/repo"
	"time"
)

type socialMediaWithUserAndPhoto struct {
	SocialMediaId             int
	SocialMediaName           string
	SocialMediaSocialMediaUrl string
	SocialMediaUserId         int
	SocialMediaCreatedAt      time.Time
	SocialMediaUpdatedAt      time.Time
	UserId                    int
	UserUsername              string
	UserEmail                 string
	UserPassword              string
	UserAge                   uint
	UserCreatedAt             time.Time
	UserUpdatedAt             time.Time
	PhotoId                   sql.NullInt64
	PhotoTitle                sql.NullString
	PhotoCaption              sql.NullString
	PhotoPhotoUrl             sql.NullString
	PhotoUserId               sql.NullInt64
	PhotoCreatedAt            sql.NullTime
	PhotoUpdatedAt            sql.NullTime
}

func (s *socialMediaWithUserAndPhoto) socialMediaWithUserAndPhotoToAggregate() repo.SocialMediaUserPhotoMapped {
	return repo.SocialMediaUserPhotoMapped{
		SocialMedia: models.SocialMedia{
			Id:             s.SocialMediaId,
			Name:           s.SocialMediaName,
			SocialMediaUrl: s.SocialMediaSocialMediaUrl,
			UserId:         s.SocialMediaUserId,
			CreatedAt:      s.SocialMediaCreatedAt,
			UpdatedAt:      s.SocialMediaUpdatedAt,
		},
		User: models.User{
			Id:        s.UserId,
			Username:  s.UserUsername,
			Email:     s.UserEmail,
			Password:  s.UserPassword,
			Age:       s.UserAge,
			CreatedAt: s.UserCreatedAt,
			UpdatedAt: s.UserUpdatedAt,
		},
		Photo: models.Photo{
			Id:        int(s.PhotoId.Int64),
			Title:     s.PhotoTitle.String,
			Caption:   s.PhotoCaption.String,
			PhotoUrl:  s.PhotoPhotoUrl.String,
			UserId:    int(s.PhotoUserId.Int64),
			CreatedAt: s.PhotoCreatedAt.Time,
			UpdatedAt: s.PhotoUpdatedAt.Time,
		},
	}
}

type socialMediaRepositoryImpl struct {
	db *sql.DB
}

const (
	addSocialMediaQuery = `
		INSERT INTO
			social_media
				(
					name,
					social_media_url,
					user_id
				)
		VALUES
			(
				$1, $2, $3
			)
		RETURNING
			id, name, social_media_url, user_id, created_at
	`
	updateSocialMediaQuery = `
		UPDATE
			social_media
		SET
			name = $2,
			social_media_url = $3,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, name, social_media_url, user_id, updated_at
	`
	deleteSocialMediaQuery = `
		DELETE FROM
			social_media
		WHERE
			id = $1
	`

	getSocialMediaQuery = `
		SELECT
			s.id,
			s.name,
			s.social_media_url,
			s.user_id,
			s.created_at,
			s.updated_at,
			u.id,
			u.username,
			MIN(p.photo_url) AS photo_url
		FROM
			social_media AS s
		LEFT JOIN
			users AS u
		ON
			s.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			p.user_id = s.user_id
		GROUP BY 
			s.id, s.name, s.social_media_url, s.user_id, s.created_at, s.updated_at, u.id, u.username
		ORDER BY
			s.id
		ASC
	`

	getSocialMediaByIdQuery = `
		SELECT
			s.id,
			s.name,
			s.social_media_url,
			s.user_id,
			s.created_at,
			s.updated_at,
			u.id,
			u.username,
			p.photo_url
		FROM
			social_media AS s
		LEFT JOIN
			users AS u
		ON
			s.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			p.id = s.user_id
		WHERE
			s.id = $1
	`
)

func NewSocialMediaRepository(db *sql.DB) repo.SocialMediaRepository {
	return &socialMediaRepositoryImpl{
		db: db,
	}
}

// AddSocialMedia implements repo.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) AddSocialMedia(socialMediaPayload *models.SocialMedia) (*dto.NewSocialMediaResponse, pkg.Error) {
	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	row := tx.QueryRow(addSocialMediaQuery, socialMediaPayload.Name, socialMediaPayload.SocialMediaUrl, socialMediaPayload.UserId)

	var socialMedia dto.NewSocialMediaResponse
	err = row.Scan(
		&socialMedia.Id,
		&socialMedia.Name,
		&socialMedia.SocialMediaUrl,
		&socialMedia.UserId,
		&socialMedia.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	return &socialMedia, nil
}

// DeleteSocialMedia implements repo.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) DeleteSocialMedia(socialMediaId int) pkg.Error {
	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return pkg.NewInternalServerError("Something went wrong!")
	}

	_, err = tx.Exec(deleteSocialMediaQuery, socialMediaId)

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

// UpdateSocialMedia implements repo.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) UpdateSocialMedia(socialMediaId int, socialMediaPayload *models.SocialMedia) (*dto.SocialMediaUpdateResponse, pkg.Error) {
	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	row := tx.QueryRow(updateSocialMediaQuery, socialMediaId, socialMediaPayload.Name, socialMediaPayload.SocialMediaUrl)

	var socialMedia dto.SocialMediaUpdateResponse
	err = row.Scan(
		&socialMedia.Id,
		&socialMedia.Name,
		&socialMedia.SocialMediaUrl,
		&socialMedia.UserId,
		&socialMedia.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	return &socialMedia, nil
}

// GetSocialMediaById implements repo.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) GetSocialMediaById(socialMediaId int) (*dto.GetSocialMedia, pkg.Error) {
	row := s.db.QueryRow(getSocialMediaByIdQuery, socialMediaId)

	var socialMedia socialMediaWithUserAndPhoto
	err := row.Scan(
		&socialMedia.SocialMediaId,
		&socialMedia.SocialMediaName,
		&socialMedia.SocialMediaSocialMediaUrl,
		&socialMedia.SocialMediaUserId,
		&socialMedia.SocialMediaCreatedAt,
		&socialMedia.SocialMediaUpdatedAt,
		&socialMedia.UserId,
		&socialMedia.UserUsername,
		&socialMedia.PhotoPhotoUrl,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.NewNotFoundError("social media not found")
		}
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	result := repo.SocialMediaUserPhotoMapped{}
	return result.HandleMappingSocialMediaWithUserAndPhotoById(repo.SocialMediaUserPhoto(socialMedia.socialMediaWithUserAndPhotoToAggregate())), nil
}

// GetSocialMedias implements repo.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) GetSocialMedias() ([]*dto.GetSocialMedia, pkg.Error) {

	rows, err := s.db.Query(getSocialMediaQuery)

	if err != nil {
		return nil, pkg.NewInternalServerError("Something went wrong!")
	}

	var socialMedias []repo.SocialMediaUserPhoto

	for rows.Next() {
		var socialMedia socialMediaWithUserAndPhoto
		err = rows.Scan(
			&socialMedia.SocialMediaId,
			&socialMedia.SocialMediaName,
			&socialMedia.SocialMediaSocialMediaUrl,
			&socialMedia.SocialMediaUserId,
			&socialMedia.SocialMediaCreatedAt,
			&socialMedia.SocialMediaUpdatedAt,
			&socialMedia.UserId,
			&socialMedia.UserUsername,
			&socialMedia.PhotoPhotoUrl,
		)

		if err != nil {
			return nil, pkg.NewInternalServerError("Something went wrong!")
		}

		socialMedias = append(socialMedias, repo.SocialMediaUserPhoto(socialMedia.socialMediaWithUserAndPhotoToAggregate()))
	}

	result := repo.SocialMediaUserPhotoMapped{}
	return result.HandleMappingSocialMediaWithUserAndPhoto(socialMedias), nil
}
