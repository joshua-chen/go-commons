/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 15:38:33
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 16:00:47
 */
package auth

import (
	_ "github.com/joshua-chen/go-commons/config"
	"github.com/joshua-chen/go-commons/exception"
	"github.com/joshua-chen/go-commons/middleware/jwt"
	"github.com/joshua-chen/go-commons/middleware/perm"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	"github.com/joshua-chen/go-commons/mvc/context/response/msg"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/context"

)

func Filter(ctx context.Context) bool {

	if !ctx.IsAjax() {
		return true
	}
	user, ok := jwt.ParseToken(ctx)
	if !ok {
		return false
	}

	yes := Enforce(user.ID, ctx.Path(), ctx.Method())
	if !yes {
		response.Unauthorized(ctx, msg.PermissionsLess, nil)
		ctx.StopExecution()
		return false
	}
	golog.Debug("HasPrivilege===> ", yes)
	return true
}

func Enforce(uid int64, path string, method string) bool {
	yes, err := perm.HasPrivilege(uid, path, method)

	if err != nil {
		exception.Fatal(err)
	}

	return yes
}
