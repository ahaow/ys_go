package response

type HelperInfoResponse struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Intro     string `json:"intro"`
	Role      int    `json:"role"`
	Address   string `json:"address"`
}
