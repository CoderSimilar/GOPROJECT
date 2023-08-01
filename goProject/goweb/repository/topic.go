package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"time"

)

//Topic和Post都是JSON数据
type Topic struct {
	Id			int64 		//话题id
	userId		int64 		//用户id	
	Title		string 		//话题标题
	Content		string 		//话题内容
	createTime	time.Time 	//创建时间
}

//repository需要实现查询操作：QueryTopicById 根据话题信息查询话题内容

//使用索引，将数据行映射为内存Map，根据索引定位数据
//定义索引
var (
	topicIndexMap	map[int64]*Topic
	postIndexMap	map[int64][]*Post
)

//定义话题索引
func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(open)//从打开的文件里读取
	topicTmpMap := make(map[int64]*Topic)//创建map键值对，
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicIndexMap = topicTmpMap
	}

	return scanner.Err()

}

