package service

import (
	"net/http"
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/repo"
)

type commentServiceMock struct {
}

type CommentService interface {
	AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, pkg.Error)
	GetComments() (*dto.GetCommentResponse, pkg.Error)
	DeleteComment(commentId int) (*dto.GetCommentResponse, pkg.Error)
	UpdateComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.GetCommentResponse, pkg.Error)
}

type commentServiceImpl struct {
	pr repo.PhotoRepository
	cr repo.CommentRepository
}

func NewCommentService(commentRepo repo.CommentRepository, photoRepo repo.PhotoRepository) CommentService {
	return &commentServiceImpl{
		pr: photoRepo,
		cr: commentRepo,
	}
}

// AddComment implements CommentService.
func (c *commentServiceImpl) AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, pkg.Error) {

	err := pkg.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	_, err = c.pr.GetPhotoId(commentPayload.PhotoId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	comment := &models.Comment{
		UserId:  userId,
		PhotoId: commentPayload.PhotoId,
		Message: commentPayload.Message,
	}

	response, err := c.cr.AddComment(comment)

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusCreated,
		Message:    "new comment successfully added",
		Data:       response,
	}, nil
}

// GetComments implements CommentService.
func (c *commentServiceImpl) GetComments() (*dto.GetCommentResponse, pkg.Error) {

	data, err := c.cr.GetComments()

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "comments successfully fetched",
		Data:       data,
	}, nil
}

// DeleteComment implements CommentService.
func (c *commentServiceImpl) DeleteComment(commentId int) (*dto.GetCommentResponse, pkg.Error) {

	err := c.cr.DeleteComment(commentId)

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "Your comment has been successfully deleted",
		Data:       nil,
	}, nil
}

// UpdateComment implements CommentService.
func (c *commentServiceImpl) UpdateComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.GetCommentResponse, pkg.Error) {

	err := pkg.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	comment := &models.Comment{
		Message: commentPayload.Message,
	}

	data, err := c.cr.UpdateComment(commentId, comment)

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "comment has been successfully updated",
		Data:       data,
	}, nil
}

var (
	AddComment    func(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, pkg.Error)
	GetComments   func() (*dto.GetCommentResponse, pkg.Error)
	DeleteComment func(commentId int) (*dto.GetCommentResponse, pkg.Error)
	UpdateComment func(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.GetCommentResponse, pkg.Error)
)

func NewCommentServiceMock() CommentService {
	return &commentServiceMock{}
}

// AddComment implements CommentService.
func (csm *commentServiceMock) AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, pkg.Error) {
	return AddComment(userId, commentPayload)
}

// GetComments implements CommentService.
func (csm *commentServiceMock) GetComments() (*dto.GetCommentResponse, pkg.Error) {
	return GetComments()
}

// DeleteComment implements CommentService.
func (csm *commentServiceMock) DeleteComment(commentId int) (*dto.GetCommentResponse, pkg.Error) {
	return DeleteComment(commentId)
}

// UpdateComment implements CommentService.
func (csm *commentServiceMock) UpdateComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.GetCommentResponse, pkg.Error) {
	return UpdateComment(commentId, commentPayload)
}
