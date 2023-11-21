package dto

// User Requests
type NewUserRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"monday.day@email.com"`
	Age      uint   `json:"age" valid:"required~Age can't be empty, range(8|150)~Minimum age is 8" example:"21"`
	Password string `json:"password" valid:"required~Password can't be empty, length(6|255)~Minimum password is 6 length" example:"secret"`
}

type UserLoginRequest struct {
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"monday.day@email.com"`
	Password string `json:"password" valid:"required~Password can't be empty" example:"secret"`
}

type UserUpdateRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"monday.day@weeekly.com"`
}

// Photo Requests
type NewPhotoRequest struct {
	Title    string `json:"title" valid:"required~Title can't be empty" example:"monday awesome"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo URL can't be empty" example:"https://www.pinterest.com/pin/807973989398829161/"`
	Caption  string `json:"caption" example:"Hello I'm Monday from Weeekly, hopefully You can do this!"`
}

type PhotoUpdateRequest struct {
	Title    string `json:"title" valid:"required~Title can't be empty" example:"monday awesome"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo URL can't be empty" example:"https://www.pinterest.com/pin/807973989398829161/"`
	Caption  string `json:"caption" example:"Hello I'm Monday from Weeekly, stay strong!"`
}

// Comment Requests
type NewCommentRequest struct {
	PhotoId int    `json:"photo_id" example:"1"`
	Message string `json:"message" valid:"required~Message can't be empty" example:"so beautiful"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" valid:"required~Message can't be empty" example:"omg so beautiful"`
}

// Social Media Requests
type NewSocialMediaRequest struct {
	Name           string `json:"name" valid:"required~Name can't be empty" example:"Monday Weeekly Official"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social media url can't be empty" example:"https://www.instagram.com/_weeekly/"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" valid:"required~Name can't be empty" example:"Weeekly Monday Official"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social media url can't be empty" example:"https://www.instagram.com/_weeekly/"`
}
