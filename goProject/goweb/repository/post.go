package repository

import (
	"time"
)

type Post struct {
	Id 				int64 //回帖id
	topicId 		int64 //回帖所属话题id
	UserId			int64 //用户id
	Content 		string //回帖内容
	createTime 	time.Time //回帖创建时间
}

//repository需要实现查询操作：QueryPostsBytopicTD 根据话题id查询所有帖子列表


