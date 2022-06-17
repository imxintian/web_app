package models

// 定义请求的参数结构体

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostVoteData struct {
	PostId    string `json:"post_id" binding:"required"`              // 帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 方向 1:赞同 -1:反对 0:取消

}
