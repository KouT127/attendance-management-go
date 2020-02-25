package responses

import "github.com/KouT127/attendance-management/models"

type UserResp struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageUrl string `json:"imageUrl"`
}

func NewUserResp(user *models.User) *UserResp {
	resp := &UserResp{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
	}
	return resp
}
