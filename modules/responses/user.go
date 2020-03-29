package responses

import "github.com/KouT127/attendance-management/models"

type UserResp struct {
	Id       string `json:"id"`
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
	User       UserResp        `json:"user"`
	Attendance *AttendanceResp `json:"attendance"`
}

func toUserResp(user *models.User) UserResp {
	resp := UserResp{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
	}
	return resp
}

func ToUserResult(user *models.User) *UserResult {
	res := &UserResult{}
	res.IsSuccessful = true
	res.User = toUserResp(user)
	return res
}

func ToUserMineResult(user *models.User, attendance *models.Attendance) *UserMineResult {
	res := &UserMineResult{}
	res.IsSuccessful = true
	res.User = toUserResp(user)
	if attendance != nil {
		res.Attendance = toAttendanceResp(attendance)
	}
	return res
}
