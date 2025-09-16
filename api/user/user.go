package user

import (
	"net/http"
	"strconv"
	user_dao "ys_go/dao/user"
	"ys_go/forms"
	"ys_go/utils/jwt"
	"ys_go/utils/pwd"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	registerForm := forms.UserRegisterForm{}

	if err := c.ShouldBind(&registerForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}

	_, err := user_dao.GetUserByMobile(registerForm.Mobile)

	if err == nil {
		// 用户存在 → 不允许注册
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户已存在"})
		return
	} else if err.Error() != "用户不存在" {
		// 数据库错误
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "查询失败"})
		return
	}

	rsp, _ := user_dao.CreateUser(&registerForm)
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建成功",
		"id":  rsp.Id,
	})
}

func Update(c *gin.Context) {
	updateForm := forms.UserUpdateForm{}
	if err := c.ShouldBind(&updateForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}
	err := user_dao.UpdateUser(&updateForm)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "用户更新成功",
	})
}

func Login(c *gin.Context) {
	passwordLoginForm := forms.UserPassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	rsp, err := user_dao.GetUserByMobile(passwordLoginForm.Mobile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		if pwd.CompareHashAndPassword(rsp.Password, passwordLoginForm.PassWord) {
			token, err := jwt.GenerateJWT(rsp.Mobile, uint(rsp.Id))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "生成token失败",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"id":        rsp.Id,
				"nick_name": rsp.Nickname,
				"token":     token,
			})

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "密码错误",
			})
			return
		}
	}
}

func List(c *gin.Context) {
	pageNo := c.DefaultQuery("pageNo", "1")
	pageNoInt, _ := strconv.Atoi(pageNo)

	pageSize := c.DefaultQuery("pageSize", "10")
	pageSizeInt, _ := strconv.Atoi(pageSize)

	rsp, err := user_dao.GetUserList(&forms.PageForm{
		PageNo:   pageNoInt,
		PageSize: pageSizeInt,
	})

	if err != nil {
		zap.S().Errorw("[GetUserList] 查询用户列表失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "查询用户列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  rsp.Data,
		"total": rsp.Total,
	})
}
