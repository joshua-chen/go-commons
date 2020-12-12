package application

import (
	stdContext "context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	_ "github.com/iris-contrib/swagger/v12"
	_ "github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/joshua-chen/go-commons/config"
	"github.com/joshua-chen/go-commons/middleware"
	_ "github.com/joshua-chen/go-commons/middleware/auth"
	_ "github.com/joshua-chen/go-commons/middleware/cors"
	"github.com/joshua-chen/go-commons/middleware/logger"
	_ "github.com/joshua-chen/go-commons/middleware/recover"
	"github.com/joshua-chen/go-commons/mvc"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	utilspath "github.com/joshua-chen/go-commons/utils/path"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	recover_middleware "github.com/kataras/iris/v12/middleware/recover"
	_ "github.com/kataras/iris/v12/sessions"

)

var closeLogFile, closeLogFile4golog func() error

func Run(appFunc func(app *iris.Application)) {
	app := newApp()
	configation(app)
	after(app)
	appFunc(app)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX 或 Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill等同于syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("shutdown...")
			timeout := 5 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}()
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

	app.Use(logger.NewConsoleLogger())
	loggerHandler, close := logger.NewRequestLogger()
	closeLogFile = close
	app.Use(loggerHandler)
	loggerHandler4golog, close4golog := logger.NewRequestLoggerForGolog()
	closeLogFile4golog = close4golog
	app.Use(loggerHandler4golog)
	app.Use(recover_middleware.New())
	app.Use(middleware.Instance().Recover.New())
	app.Use(middleware.Instance().Auth.New())
	app.Use(middleware.Instance().Cors.New()) // cors

	// yaag api 为文档生成器
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Iris",
		DocPath:   "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New())

	config.RegisterView(app)
	config.HandleStatic(app)
	config.HandleUpload(app)
	mvc.ConfigureMain(app)
	/*sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.RequestURI)
	})
	sillyConvertedToIon := iris.FromStd(sillyHTTPHandler)
	app.Use(sillyConvertedToIon)
	*/
	
	return app
}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	path := utilspath.GetFullPath("./config/iris.yml")
	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.YAML(path)))
	//
	golog.Info("[app.Configure]==>  ok")

	app.ConfigureHost(func(host *iris.Supervisor) { // <- 重要
		//您可以使用某些主机的方法控制流或延迟某些内容：
		// host.RegisterOnError
		// host.RegisterOnServe
		host.RegisterOnShutdown(func() {
			app.Logger().Infof("Application shutdown on signal")
			closeLogFile()
			closeLogFile4golog()
		})
	})

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(response.NewNotFoundResult())
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(response.NewErrorResult(iris.StatusInternalServerError * 100))
	})
}

func after(app *iris.Application) {
	// 主持后置
	app.Done(func(ctx iris.Context) {
		//golog.Debug("后置............")
		ctx.Application().Logger().Debug("后置............")
	})
}
