/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-25 17:37:30
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 20:34:29
 */

package response

import (
	"fmt"
	_ "strings"

	_ "github.com/jmespath/go-jmespath"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/context"
	_ "github.com/kataras/iris/v12/hero"

)

func DefaultResult(data interface{}) Result {

	//var result = new(models.ResponseResult)
	//result := models.NewResponseResult(data, "200")
	return NewResult(data, 200)
	//return result
}
func BoolResult(data bool) Result {

	//var result = new(models.ResponseResult)
	//result := models.NewResponseResult(data, "200")
	return NewBoolResult(data, 200)
	//return result
}
func NewBoolResult(data bool, c int, m ...string) Result {
	r := Result{Data: iris.Map{}, Code: c, Success: true}

 	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}
func NewResult(data interface{}, c int, m ...string) Result {
	r := Result{Data: data, Code: c, Success: false}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Success = true
		r.Msg = "SUCCESS"
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}

func NewUnauthorizedResult(msg string, data ...interface{}) Result {
	result := Result{Code: iris.StatusUnauthorized, Msg: msg, Success: false}
	if len(data) > 0 {
		result.Data = data[0]
	}
	return result
}
func NewSuccessResult(data interface{}, c int, msg ...string) Result {
	result := Result{Data: data, Code: c, Success: true}
	if len(msg) > 0 {
		result.Msg = msg[0]
	}
	return result
}
func NewNotFoundResult(msg ...string) Result {
	result := Result{Code: iris.StatusNotFound, Msg: "not found", Data: iris.Map{}}
	if len(msg) > 0 {
		result.Msg = msg[0]
	}
	return result
}

func NewErrorResult(code int, msg ...string) Result {
	result := Result{Code: code, Msg: "server interal error", Data: iris.Map{}}
	if len(msg) > 0 {
		result.Msg = msg[0]
	}
	return result
}

func Fail(ctx iris.Context, statusCode int, format string, a ...interface{}) {
	err := HttpError{
		Code:   statusCode,
		Reason: fmt.Sprintf(format, a...),
	}
	//记录所有> = 500内部错误。
	if statusCode >= 500 {
		ctx.Application().Logger().Error(err)
	}
	ctx.StatusCode(statusCode)
	ctx.JSON(err)
	//没有下一个处理程序将运行。
	ctx.StopExecution()
}

func Ok(ctx iris.Context, data interface{}, msg ...string) {
	ctx.StatusCode(iris.StatusOK)
	result := NewSuccessResult(data, iris.StatusOK, msg...)
	ctx.JSON(result)
}

// 401 error define
func Unauthorized(ctx iris.Context, msg string, data interface{}) {
	result := NewUnauthorizedResult(msg, data)
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(result)
}

// common error define
func Error(ctx iris.Context, statusCode int, msg ...string) {
	result := NewErrorResult(statusCode, msg...)

	result.Code = statusCode

	err := HttpError{
		Code:   statusCode,
		Reason: result.Msg,
	}
	//记录所有> = 500内部错误。
	if statusCode >= 500 {
		ctx.Application().Logger().Error(err)
	}

	ctx.StatusCode(statusCode)
	ctx.JSON(result)
}
