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
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

)
 
var homeView = mvc.View{
    Name: "index.html",
    Data: iris.Map{"Title": "Home"},
}

type homeController struct {

}
// GetRegister 处理 GET: http://localhost:8080/user/register.
func (c *homeController) GetRegister() mvc.Result {
   
    return homeView
}
// Configure
func Configure(party router.Party, configurators ...func(*mvc.Application)) *mvc.Application{
 return	mvc.Configure(party, configurators...)
}
// ConfigureHome
func ConfigureHome(app *iris.Application){
	mvc.Configure(app.Party("/home"), homeMVC)
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
    app.Handle(new(homeController))
    //所有依赖项被绑定在父 *mvc.Application 
    
}