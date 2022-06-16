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

/*
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int64     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
*/
// GetPostById 查询帖子详情
func GetPostById(id int64) (*models.Post, error) {
	sqlStr := "select  post_id, title, content, author_id, community_id," +
		" `status`, create_time from post where post_id = ?"
	row := db.QueryRow(sqlStr, id)
	p := &models.Post{}
	err := row.Scan(&p.ID, &p.Title, &p.Content, &p.AuthorID, &p.CommunityID, &p.Status, &p.CreateTime)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// GetPostList 查询帖子列表
func GetPostList(page, pageSize int) ([]*models.Post, error) {
	sqlStr := "select post_id, title, content, author_id, community_id," +
		"`status`, create_time from post   limit ?, ?"
	rows, err := db.Query(sqlStr, (page-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		p := &models.Post{}
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.AuthorID, &p.CommunityID, &p.Status, &p.CreateTime)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)

	}
	return posts, nil
}
