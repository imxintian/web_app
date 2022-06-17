package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成postId
	p.ID = snowflake.GenID()
	// 2. 存储到数据库, 并返回err
	err = mysql.CreatePost(p)
	if err != nil {
		zap.L().Error("createPost error", zap.Error(err))
		return err
	}
	err = redis.CreatePostVote(p.ID)
	if err != nil {
		zap.L().Error("createPostVote error", zap.Error(err))
		return err
	}
	return nil
}

// GetPostById 获取帖子详情
func GetPostById(id int64) (*models.APiPostDetail, error) {
	// 查询并组合帖子详情数据接口
	post, err := mysql.GetPostById(id)
	if err != nil {
		zap.L().Error("getPostById error", zap.Error(err))
		return nil, err
	}
	// 根据作者id 查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("getUserById error", zap.Error(err))
		return nil, err
	}
	// 根据社区id 查询社区信息
	community, err := mysql.GetCommunityById(post.CommunityID)
	if err != nil {
		zap.L().Error("getCommunityById error", zap.Error(err))
		return nil, err
	}
	// 组合数据接口
	postDetail := &models.APiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return postDetail, nil

}

// GetPostList 获取帖子列表
func GetPostList(page, pageSize int) ([]*models.APiPostDetail, error) {
	// 查询帖子列表
	postList, err := mysql.GetPostList(page, pageSize)
	if err != nil {
		zap.L().Error("getPostList error", zap.Error(err))
		return nil, err
	}

	data := make([]*models.APiPostDetail, 0, len(postList))
	// 根据作者id查询作者信息
	for _, post := range postList {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("getUserById error", zap.Error(err))
			return nil, err
		}
		// 根据社区id查询社区信息
		community, err := mysql.GetCommunityById(post.CommunityID)
		if err != nil {
			zap.L().Error("getCommunityById error", zap.Error(err))
			return nil, err
		}
		// 组合数据接口
		postDetail := &models.APiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, nil
}
