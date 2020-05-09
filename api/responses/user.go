package responses

import "github.com/KouT127/attendance-management/domain/models"

type UserResp struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageUrl string `json:"imageUrl"`
}

type UserResult struct {
	CommonResponse
	User UserResp `json:"user"`
}

type UserMineResult struct {
	CommonResponse
	User UserResp `json:"user"`
}

func toUserResp(user *models.User) UserResp {
	resp := UserResp{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		ImageUrl: user.ImageURL,
	}
	return resp
}

func ToUserResult(user *models.User) *UserResult {
	res := &UserResult{}
	res.IsSuccessful = true
	res.User = toUserResp(user)
	return res
}

func ToUserMineResult(user *models.User) *UserMineResult {
	res := &UserMineResult{}
	res.IsSuccessful = true
	res.User = toUserResp(user)
	return res
}
