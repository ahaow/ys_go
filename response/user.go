package response

import "ys_go/model"

// 用户信息响应
type UserInfoResponse struct {
	Id       int32  `json:"id" example:"1"`
	Password string `json:"password,omitempty" example:"hashed_password"`
	Mobile   string `json:"mobile" example:"13800138000"`
	Nickname string `json:"nickname" example:"张三"`
	BirthDay uint64 `json:"birthDay,omitempty" example:"1737052800"` // 时间戳
	Gender   string `json:"gender" example:"male"`
	Role     int32  `json:"role" example:"1"`
}

// 用户列表响应
type UserListResponse struct {
	Total int32               `json:"total" example:"100"`
	Data  []*UserInfoResponse `json:"data"`
}

// 转换 User model 为 UserInfoResponse
func UserModelToResponse(user model.User, includePassword bool) UserInfoResponse {
	userInfoRsp := UserInfoResponse{
		Id: user.ID,
		// Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}

	if includePassword {
		userInfoRsp.Password = user.Password
	}

	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}
