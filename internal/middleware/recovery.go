package middleware

import (
	"blog-server/global"
	"blog-server/pkg/app"
	"blog-server/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort() //终止程序执行，直接从调用的地方跳出
			}
		}()

		/*
				c.Next()之前的操作是在Handler执行之前就执行
				c.Next()之后的操作是在Handler执行之后就执行

			之前的操作一般用来做验证处理，访问是否允许之类的，
			之后的操作一般是用来做总结处理的，比如格式化输出，相应结束时间，相应时长计算之类的
		*/
		c.Next()
		fmt.Println("后面的要不还要执行")
	}
}
