package response

type HelperInfoResponse struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Intro     string `json:"intro"`
	Role      int    `json:"role"`
	Status    int    `json:"status"`
	Address   string `json:"address"`
}

// 用户列表响应
type HelperListResponse struct {
	Total int32                 `json:"total" example:"100"`
	Data  []*HelperInfoResponse `json:"data"`
}
