package user_dao

import (
	"errors"
	"time"
	"ys_go/dao"
	"ys_go/forms"
	"ys_go/global"
	"ys_go/model"
	"ys_go/response"
	"ys_go/utils/pwd"

	"go.uber.org/zap"
)

func CreateUser(req *forms.UserRegisterForm) (*response.UserInfoResponse, error) {
	var user model.User
	user.Mobile = req.Mobile
	user.NickName = req.Mobile
	user.Password = req.PassWord
	// 密码加密
	newPwd, err := pwd.GenerateFromPassword(user.Password)
	if err != nil {
		return nil, errors.New("密码设置错误")
	}
	user.Password = newPwd
	result := global.DB.Create(&user)
	if result.Error != nil {
		zap.S().Errorw("[CreateUser] 创建用户失败")
		return nil, errors.New("创建用户失败")
	}
	userInfoRsp := response.UserModelToResponse(user, true)
	return &userInfoRsp, nil
}

func UpdateUser(req *forms.UserUpdateForm) error {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return errors.New("用户不存在")
	}
	birthDay := time.Unix(int64(req.Birthday), 0)
	user.Birthday = &birthDay
	user.NickName = req.NickName
	user.Gender = req.Gender

	result = global.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByMobile(mobile string) (*response.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: mobile}).First(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := response.UserModelToResponse(user, true)
	return &userInfoRsp, nil
}

func GetUserList(req *forms.PageForm) (*response.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	resp := &response.UserListResponse{}
	resp.Total = int32(result.RowsAffected)

	// 进行分页
	global.DB.Scopes(dao.Paginate(int(req.PageNo), int(req.PageSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := response.UserModelToResponse(user, false)
		resp.Data = append(resp.Data, &userInfoRsp)
	}

	return resp, nil

}
