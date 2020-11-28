package application

import (
	"os"
	_ "time"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	_ "github.com/iris-contrib/swagger/v12"
	_ "github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/joshua-chen/go-commons/utils"
	"github.com/joshua-chen/go-commons/config"
	"github.com/joshua-chen/go-commons/middleware"
	_ "github.com/joshua-chen/go-commons/middleware/auth"
	_ "github.com/joshua-chen/go-commons/middleware/cors"
	_ "github.com/joshua-chen/go-commons/middleware/recover"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	recover_middleware "github.com/kataras/iris/v12/middleware/recover"
	_ "github.com/kataras/iris/v12/mvc"
	_ "github.com/kataras/iris/v12/sessions"

)

func Run(appFunc func(app *iris.Application)) {
	app := newApp()
	configation(app)
	after(app)
	appFunc(app)
	app.Run(
		iris.Addr(":"+config.AppConfig.Port), //在端口8080进行监听
		iris.WithCharset("UTF-8"),
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)

}

//构建App
func newApp() *iris.Application {
	app := iris.New()
	app.AllowMethods(iris.MethodOptions)
	//设置日志级别  开发阶段为debug

	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover_middleware.New())
	app.Use(logger.New())
	app.Use(middleware.Instance().Recover.New())
	app.Use(middleware.Instance().Auth.New())
	app.Use(middleware.Instance().Cors.New()) // cors

	// yaag api 为文档生成器
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New())

	exists, _ := utils.PathExisted("views")
	if exists {
		app.RegisterView(iris.HTML("./views", ".html"))
		golog.Info("[RegisterView]==> ./views, ok")
	}

	/*sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.RequestURI)
	})
	sillyConvertedToIon := iris.FromStd(sillyHTTPHandler)
	app.Use(sillyConvertedToIon)
	*/

	/*


			//注册视图文件
			app.RegisterView(iris.HTML("./static", ".html"))
			app.Get("/", func(context context.Context) {
				context.View("index.html")
			})



		authConfig := basicauth.Config{
			Users:   map[string]string{"wangshubo": "wangshubo", "superWang": "superWang"},
			Realm:   "Authorization Required",
			Expires: time.Duration(30) * time.Minute,
		}

		authentication := basicauth.New(authConfig)
	*/

	/*app.Get("/", func(ctx context.Context) { ctx.Redirect("/admin") })

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", h)
		// http://localhost:8080/admin/profile
		needAuth.Get("/profile", h)

		// http://localhost:8080/admin/settings
		needAuth.Get("/settings", h)
	}
	*/

	return app
}

func handleStatic(app *iris.Application) {
	//注册静态资源
	staticPath := config.AppConfig.StaticPath
	app.HandleDir(staticPath[0], staticPath[1])
	//app.HandleDir("/manage/static", staticPath[1])
	app.HandleDir(staticPath[0]+"/images", staticPath[1]+"/images")
}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	path:= utils.GetAbsolutePath("./config/iris.yml");
	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.YAML(path)))
	//
	golog.Info("[app.Configure]==>  ok")

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(response.NewNotFoundResult())
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(response.NewErrorResult(iris.StatusInternalServerError))
	})
}

func after(app *iris.Application) {
	// 主持后置
	app.Done(func(ctx iris.Context) {
		//golog.Debug("后置............")
		ctx.Application().Logger().Debug("后置............")
	})
}
