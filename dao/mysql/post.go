package mysql

import (
	"web_app/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post(post_id, title, content, author_id, community_id) values(?, ?, ?, ?, ?)"
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return err
	}

	return nil
}

// GetPostById 查询帖子详情
func GetPostById(id int64) (*models.Post, error) {
	sqlStr := "select post_id, title, content, author_id, community_id from post where post_id = ?"
	row := db.QueryRow(sqlStr, id)
	p := &models.Post{}
	err := row.Scan(&p.ID, &p.Title, &p.Content, &p.AuthorID, &p.CommunityID)
	if err != nil {
		return nil, err
	}

	return p, nil
}
