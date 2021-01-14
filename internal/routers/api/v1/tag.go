package v1

import (
	"blog-server/internal/service"
	"blog-server/pkg/app"
	"github.com/gin-gonic/gin"
)

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
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	//valid, errs := true, "ert" //TODO
	//if !valid {
	//	global.Logger.WithCaller()
	//	response.ToErrorResponse()
	//	return
	//}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})

	if err != nil {

	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {

	}
	response.ToResponseList(tags, totalRows)
	return

}
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{
		Name:     "",
		CreateBy: "",
		State:    0,
	}
}
func (t Tag) Update(c *gin.Context) {

}
func (t Tag) Delete(c *gin.Context) {

}
