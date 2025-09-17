package helper

import (
	"net/http"
	"strconv"
	helper_dao "ys_go/dao/helper"
	"ys_go/forms"
	"ys_go/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func Update(c *gin.Context) {
	updateForm := forms.HelperUpdateRequest{}

	if err := c.ShouldBind(&updateForm); err != nil {
		response.FailWithMsg("请求参数错误", c)
		return
	}

	_, err := helper_dao.GetHelerById(int64(updateForm.Id))

	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	err = helper_dao.UpdateHelper(&updateForm)

	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithMsg("更新成功", c)
}

func Delete(c *gin.Context) {
	deleteForm := forms.HelperDeleteRequest{}
	if err := c.ShouldBind(&deleteForm); err != nil {
		response.FailWithMsg("请求参数错误", c)
		return
	}
	_, err := helper_dao.GetHelerById(int64(deleteForm.Id))
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	err = helper_dao.DeleteHelper(&deleteForm)

	if err != nil {
		response.FailWithMsg(err.Error(), c)
	}
	response.OkWithMsg("删除活动集成功", c)

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

func List(c *gin.Context) {
	pageNo := c.DefaultQuery("pageNo", "1")
	pageNoInt, _ := strconv.Atoi(pageNo)

	pageSize := c.DefaultQuery("pageSize", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)

	rsp, err := helper_dao.GetHelperList(&forms.PageForm{
		PageNo:   pageNoInt,
		PageSize: pageSizeInt,
	})

	if err != nil {
		zap.S().Errorw("[GetHelperList] 查询活动集列表失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "查询用户列表失败",
		})
		return
	}

	response.OkWithList(rsp.Data, int64(rsp.Total), c)
}
