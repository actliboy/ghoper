package router

//go:generate qtc -dir=../template
import (
	"context"
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/iris-contrib/middleware/prometheus"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	. "github.com/kataras/iris/middleware/recover"
	"hoper/client/controller"
	"hoper/client/controller/common/logging"
	"hoper/client/controller/hnsq"
	"hoper/client/controller/hwebsocket"
	"hoper/client/controller/upload"
	"hoper/client/middleware"
	"net/http"
	"runtime/debug"
	"time"
)

func init() {
	raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
}

func IrisRouter() *iris.Application {
	app := iris.New()

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		//关闭所有主机
		app.Shutdown(ctx)
	})
	app.StaticWeb("/api/static", "../static")

	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		router(w, r)
	})
	//api文档
	/*	yaag.Init(&yaag.Config{
			On:       true,                 //是否开启自动生成API文档功能
			DocTitle: "Iris",
			DocPath:  "../static/api/apidoc.html",        //生成API文档名称存放路径
			BaseUrls: map[string]string{"Production": "", "Staging": ""},
		})
		//注册中间件
		app.Use(irisyaag.New())*/
	//https://rpm.newrelic.com/accounts/2269290/applications
	/*	config := newrelic.Config("hoper", "199e00247f278548fe92d6c81aeaadac0fc52b4b")
		m, err := newrelic.New(config)
		if err != nil {
			app.Logger().Fatal(err)
		}
		app.Use(m.ServeHTTP)*/
	app.Use(New())

	prometheus := prometheus.New("hoper")
	app.Use(prometheus.ServeHTTP)
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		//错误代码处理程序不与其他路由共享相同的中间件，所以单独执行错误
		prometheus.ServeHTTP(ctx)
		ctx.Writef("Not Found")
	})

	/*middleware必须要写ctx.next(),且写在路由前，路由后的midddleware在请求之前的路由时不生效
	  iris.FromStd()将其他Handler转为iris的Handler
	*/
	globalLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "../i18n/locale_en-US.ini",
			"zh-CN": "../i18n/locale_zh-CN.ini"}})
	app.Use(globalLocale)

	app.Logger().Printer.SetOutput(logging.F)

	UserRouter(app)

	WS(app)

	TPLRouter(app)

	ArticleRouter(app)

	MomentRouter(app)

	GraphqlRouter(app)

	Smart(app)

	app.Post("/api/upload/{classify:string}", iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.Upload(ctx)
	})
	app.Post("/api/upload_multiple/{classify:string}", iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.UploadMultiple(ctx)
	})

	//获取标签
	app.Get("/api/tag", controller.GetTags)

	app.Post("/api/comment/:classify", middleware.JWT, controller.AddComment)

	app.Get("/api/push", controller.Push)

	app.Get("/api/chat/getChat", hwebsocket.GetChat)

	app.Post("/api/nsq", hnsq.Start)

	return app
}
