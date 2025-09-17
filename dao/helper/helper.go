package helper_dao

import (
	"errors"
	"fmt"
	"time"
	"ys_go/dao"
	"ys_go/forms"
	"ys_go/global"
	"ys_go/model"
	"ys_go/response"
)

func GetHelerByName(name string) (*response.HelperInfoResponse, error) {
	var helper model.Helper
	result := global.DB.Where(&model.Helper{Name: name}).First(&helper)

	if result.RowsAffected == 0 {
		return nil, errors.New("没有找到相关模版集")
	}

	if result.Error != nil {
		return nil, result.Error
	}
	// 时间格式化
	var startTimeStr, endTimeStr string
	if helper.StartTime != nil {
		startTimeStr = helper.StartTime.Format("2006-01-02 15:04:05")
	}
	if helper.EndTime != nil {
		endTimeStr = helper.EndTime.Format("2006-01-02 15:04:05")
	}
	heplerInfo := response.HelperInfoResponse{
		Id:        helper.ID,
		Name:      helper.Name,
		Intro:     helper.Intro,
		StartTime: startTimeStr,
		EndTime:   endTimeStr,
		Role:      helper.Role,
	}

	return &heplerInfo, nil

}

func GetHelerById(id int64) (*response.HelperInfoResponse, error) {
	var helper model.Helper
	if err := global.DB.First(&helper, id).Error; err != nil {
		return nil, errors.New("没有找到相关活动集")
	}

	var startTimeStr, endTimeStr string
	if helper.StartTime != nil {
		startTimeStr = helper.StartTime.Format("2006-01-02 15:04:05")
	}
	if helper.EndTime != nil {
		endTimeStr = helper.EndTime.Format("2006-01-02 15:04:05")
	}
	heplerInfo := response.HelperInfoResponse{
		Id:        helper.ID,
		Name:      helper.Name,
		Intro:     helper.Intro,
		StartTime: startTimeStr,
		EndTime:   endTimeStr,
		Role:      helper.Role,
	}
	return &heplerInfo, nil
}

func CreateHelper(req *forms.HelperCreateRequest) (*response.HelperInfoResponse, error) {
	var helper model.Helper

	helper.Name = req.Name
	helper.Intro = req.Intro

	// 时间解析
	layout := "2006-01-02 15:04:05"

	if req.StartTime != "" {
		t, err := time.Parse(layout, req.StartTime)
		if err != nil {
			return nil, fmt.Errorf("开始时间格式错误: %v", err)
		}
		helper.StartTime = &t
	}

	if req.EndTime != "" {
		t, err := time.Parse(layout, req.EndTime)
		if err != nil {
			return nil, fmt.Errorf("结束时间格式错误: %v", err)
		}
		helper.EndTime = &t
	}

	// 保存数据库
	if err := global.DB.Save(&helper).Error; err != nil {
		return nil, err
	}

	// 格式化响应
	var startTimeStr, endTimeStr string
	if helper.StartTime != nil {
		startTimeStr = helper.StartTime.Format(layout)
	}
	if helper.EndTime != nil {
		endTimeStr = helper.EndTime.Format(layout)
	}

	helperInfo := response.HelperInfoResponse{
		Id:        int32(helper.ID),
		Name:      helper.Name,
		Intro:     helper.Intro,
		StartTime: startTimeStr,
		EndTime:   endTimeStr,
		Role:      helper.Role,
		Address:   helper.Address,
	}
	return &helperInfo, nil
}

func UpdateHelper(req *forms.HelperUpdateRequest) error {
	layout := "2006-01-02 15:04:05"

	var startTimePtr, endTimePtr *time.Time

	if req.StartTime != "" {
		t, err := time.Parse(layout, req.StartTime)
		if err != nil {
			return fmt.Errorf("开始时间格式错误: %v", err)
		}
		startTimePtr = &t
	}

	if req.EndTime != "" {
		t, err := time.Parse(layout, req.EndTime)
		if err != nil {
			return fmt.Errorf("结束时间格式错误: %v", err)
		}
		endTimePtr = &t
	}

	updateData := map[string]interface{}{
		"name":       req.Name,
		"intro":      req.Intro,
		"status":     req.Status,
		"role":       req.Role,
		"start_time": startTimePtr,
		"end_time":   endTimePtr,
	}

	if err := global.DB.Model(&model.Helper{}).Where("id = ?", req.Id).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func DeleteHelper(req *forms.HelperDeleteRequest) error {
	if err := global.DB.Delete(&model.Helper{}, req.Id).Error; err != nil {
		return err
	}
	return nil
}

func GetHelperList(req *forms.PageForm) (*response.HelperListResponse, error) {
	var helpers []model.Helper
	result := global.DB.Find(&helpers)

	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &response.HelperListResponse{}
	rsp.Total = int32(result.RowsAffected)

	// 进行分页
	global.DB.Scopes(dao.Paginate(int(req.PageNo), int(req.PageSize))).Find(&helpers)

	for _, h := range helpers {
		var startTimeStr, endTimeStr string
		if h.StartTime != nil {
			startTimeStr = h.StartTime.Format("2006-01-02 15:04:05")
		}
		if h.EndTime != nil {
			endTimeStr = h.EndTime.Format("2006-01-02 15:04:05")
		}
		helperInfo := response.HelperInfoResponse{
			Id:        h.ID,
			Name:      h.Name,
			Intro:     h.Intro,
			Role:      h.Role,
			Status:    h.Status,
			StartTime: startTimeStr,
			EndTime:   endTimeStr,
		}
		rsp.Data = append(rsp.Data, &helperInfo)
	}
	return rsp, nil
}
