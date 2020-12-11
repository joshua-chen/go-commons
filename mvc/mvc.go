/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 18:29:49
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 18:30:00
 */
package mvc

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"

)

type HomeController struct {
}

// GetRegister 处理 GET: http://localhost:8080/home.
func (c *HomeController) GetHome() mvc.Result {

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"Title": "Home"},
	}
}
func (c *HomeController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)
	//甚至添加基于此控制器路由的全局中间件，
	//在这个例子中是根“/”：
	// b.Router().Use(myMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide 服务
// 请求方法:   GET
// 请求资源路径: http://localhost:8080/custom_path
func (c *HomeController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}

//AfterActivation，所有依赖项都被设置,因此访问它们是只读
func (c *HomeController) AfterActivation(a mvc.AfterActivation) {}

// Configure
func Configure(party router.Party, configurators ...func(*mvc.Application)) *mvc.Application {
	return mvc.Configure(party, configurators...)
}

// ConfigureHome
func ConfigureHome(app *iris.Application) {
	//mvc.Configure(app.Party("/
	mvcApp := mvc.New(app.Party("/"))
	mvcApp.Handle(new(HomeController))
}

func homeMVC(app *mvc.Application) {
	//当然，你可以在MVC应用程序中使用普通的中间件。
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	//把依赖注入，controller(s)绑定
	//可以是一个接受iris.Context并返回单个值的函数（动态绑定）
	//或静态结构值（service）。
	app.Register(
	// sessions.New(sessions.Config{}).Start,
	//  &prefixedLogger{prefix: "DEV"},
	)
	// GET: http://localhost:8080/basic
	// GET: http://localhost:8080/basic/custom
	app.Handle(new(HomeController))
	//所有依赖项被绑定在父 *mvc.Application

}
