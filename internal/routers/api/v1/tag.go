package v1

import (
	"blog-server/global"
	"blog-server/internal/service"
	"blog-server/pkg/app"
	"blog-server/pkg/convert"
	"blog-server/pkg/errcode"
	"fmt"
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
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid:%v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})

	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag:%v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
	}
	//测试异常
	var tmp map[int]int

	tmp[1] = 2
	fmt.Println("测试异常：", tmp)

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList:%v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
	}
	response.ToResponseList(tags, totalRows)
	fmt.Println("处理业务")
	return

}
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid:%v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(errcode.Success)
	return

}
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid:%v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag:%v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(errcode.Success)
	return
}
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid:%v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	svc := service.New(c)
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Debug(c, "svc.DeleteTag: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(errcode.Success)
}
