package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
)

func MomentRouter(app *iris.Application) {

	momentRouter := app.Party("/api/moment")
	{
		//获取文章列表
		momentRouter.Get("/", controller.GetMomentsV2)
		//获取文章列表
		momentRouter.Get("/{id:uint64}", middleware.JWT, controller.GetMoment)
		//新建文章
		momentRouter.Post("", middleware.JWT, controller.AddMoment)
		//更新指定文章
		momentRouter.Put("/{id:uint64}", middleware.JWT, controller.EditMoment)
		//删除指定文章
		momentRouter.Delete("/{id:uint64}", middleware.JWT, controller.DeleteMoment)
	}
}
