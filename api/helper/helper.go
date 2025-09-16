package helper

import (
	"strconv"
	helper_dao "ys_go/dao/helper"
	"ys_go/forms"
	"ys_go/response"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	createForm := forms.HelperCreateRequest{}

	if err := c.ShouldBind(&createForm); err != nil {
		response.FailWithMsg("请求参数错误", c)
		return
	}

	// 先查是否存在
	_, err := helper_dao.GetHelerByName(createForm.Name)
	if err == nil {
		// 找到了 → 名称已存在
		response.FailWithMsg("该名称已存在", c)
		return
	} else if err.Error() != "没有找到相关模版集" {
		response.FailWithMsg(err.Error(), c)
		return
	}

	rsp, err := helper_dao.CreateHelper(&createForm)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(map[string]any{
		"id":   rsp.Id,
		"name": rsp.Name,
	}, c)
}

func Info(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		response.FailWithMsg("请求参数错误", c)
		return
	}
	rsp, err := helper_dao.GetHelerById(i)

	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData(rsp, c)
}
