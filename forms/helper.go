package forms

type HelperCreateRequest struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Intro     string `form:"intro" json:"intro" binding:"required"`
	StartTime string `form:"start_time" json:"start_time" binding:"required"` // 接收字符串
	EndTime   string `form:"end_time" json:"end_time" binding:"required"`
}
