package v1

import "github.com/gin-gonic/gin"

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

//@Summary 获取标签
//@Product json
//@Success 200 {string} string "成功"
func (t Tag) Get(c *gin.Context) {

}
func (t Tag) List(c *gin.Context) {

}
func (t Tag) Create(c *gin.Context) {

}
func (t Tag) Upate(c *gin.Context) {

}
func (t Tag) Delete(c *gin.Context) {

}
