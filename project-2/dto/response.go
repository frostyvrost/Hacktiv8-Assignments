package dto

import "time"

// User Responses
type UserUpdateResponse struct {
	Id        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"monday"`
	Email     string    `json:"email" example:"monday.day@weeekly.com"`
	Age       uint      `json:"age" example:"21"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type UserResponse struct {
	Id       int    `json:"id" example:"1"`
	Username string `json:"username" example:"monday"`
	Email    string `json:"email" example:"monday.day@email.com"`
	Age      uint   `json:"age" example:"21"`
}

type TokenResponse struct {
	Token string `json:"token" example:"random string"`
}

type GetUserResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// Photo Responses
type PhotoResponse struct {
	Id        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"monday awesome"`
	Caption   string    `json:"caption" example:"Hello I'm Monday from weeekly, hopefully You can do this!"`
	PhotoUrl  string    `json:"photo_url" example:"https://www.pinterest.com/pin/807973989398829161/"`
	UserId    int       `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type PhotoUpdateResponse struct {
	Id        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"monday awesome"`
	Caption   string    `json:"caption" example:"Hello I'm Monday from Weeekly, stay strong!"`
	PhotoUrl  string    `json:"photo_url" example:"https://www.pinterest.com/pin/807973989398829161/"`
	UserId    int       `json:"user_id" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type GetPhotoResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// Comment Responses
type NewCommentResponse struct {
	Id        int       `json:"id" example:"1"`
	UserId    int       `json:"user_id" example:"1"`
	PhotoId   int       `json:"photo_id" example:"1"`
	Message   string    `json:"message" example:"so beautifull"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type GetCommentResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// Social Media Responses
type NewSocialMediaResponse struct {
	Id             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"Monday Weeekly Official"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://www.instagram.com/_weeekly/"`
	UserId         int       `json:"user_id" example:"1"`
	CreatedAt      time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type SocialMediaUpdateResponse struct {
	Id             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"Monday Weeekly Official"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://www.instagram.com/_weeekly/"`
	UserId         int       `json:"user_id" example:"1"`
	UpdatedAt      time.Time `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type SocialMediaUser struct {
	Id              int    `json:"id" example:"1"`
	Username        string `json:"username" example:"monday"`
	ProfileImageUrl string `json:"profile_image_url" example:"https://www.pinterest.com/pin/807973989398829161/"`
}

type GetSocialMedia struct {
	Id             int             `json:"id" example:"1"`
	Name           string          `json:"name" example:"Monday Weeekly Official"`
	SocialMediaUrl string          `json:"social_media_url" example:"https://www.instagram.com/_weeekly/"`
	UserId         int             `json:"user_id" example:"1"`
	CreatedAt      time.Time       `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
	UpdatedAt      time.Time       `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
	User           SocialMediaUser `json:"user"`
}

type GetSocialMediaResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type GetSocialMediaHttpResponse struct {
	StatusCode  int               `json:"status_code" example:"200"`
	Message     string            `json:"message" example:"social media successfully fetched"`
	SocialMedia []*GetSocialMedia `json:"social_media"`
}
