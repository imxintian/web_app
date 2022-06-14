package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成postId
	p.ID = snowflake.GenID()
	// 2. 存储到数据库, 并返回err
	return mysql.CreatePost(p)
}
