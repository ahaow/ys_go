package forms

// 分页参数
type PageForm struct {
	PageNo   int `form:"pageNo" json:"pageNo" binding:"required" example:"1"`      // 页码
	PageSize int `form:"pageSize" json:"pageSize" binding:"required" example:"10"` // 每页数量
}

// 注册请求
type UserRegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required" example:"13800138000"`             // 手机号
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=10" example:"123456"` // 密码
}

// 登录请求
type UserPassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required" example:"13800138000"`             // 手机号
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=10" example:"123456"` // 密码
	// Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5" example:"12345"`      // 验证码
	// CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required" example:"captcha-uuid-001"` // 验证码ID
}

// 更新请求
type UserUpdateForm struct {
	Id       int32  `form:"id" json:"id" binding:"required"`
	NickName string `form:"nickname" json:"nickname" binding:"required"`
	Birthday uint64 `form:"birthday" json:"birthday"`
	Gender   string `form:"gender" json:"gender" binding:"required"`
}
