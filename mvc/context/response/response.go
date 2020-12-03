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

func Result(success bool, data interface{}, msg ...string) JsonResult {

	if success {
		return NewSuccessResult(data, msg...)
	}

	return NewFailResult(data, msg...)
}

func DefaultResult(data interface{}) JsonResult {

	return NewResult(data, StatusOK)
}
func BoolResult(data bool) JsonResult {

	return NewBoolResult(data, StatusOK)
}
func NewBoolResult(data bool, c int, m ...string) JsonResult {
	r := JsonResult{Data: iris.Map{}, Code: c, Success: true}

	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}
func NewResult(data interface{}, c int, m ...string) JsonResult {
	r := JsonResult{Data: data, Code: c, Success: false}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Success = true
		r.Msg = ""
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}

func NewUnauthorizedResult(msg string, data ...interface{}) JsonResult {
	result := JsonResult{Code: StatusUnauthorized, Msg: msg, Success: false}
	if len(data) > 0 {
		result.Data = data[0]
	}
	return result
}
func NewSuccessResult(data interface{}, msg ...string) JsonResult {
	result := JsonResult{Data: data, Code: StatusOK, Success: true}
	if len(msg) > 0 {
		result.Msg = msg[0]
	}
	return result
}
func NewFailResult(data interface{}, msg ...string) JsonResult {
	result := JsonResult{Data: data, Code: StatusExpectationFailed, Success: false}
	if len(msg) > 0 {
		result.Msg = msg[0]
	}
	return result
}
func NewNotFoundResult(msg ...string) JsonResult {
	result := JsonResult{Code: StatusNotFound, Msg: "not found"}
	if len(msg) > 0 {
		result.Msg = msg[0]
	}
	return result
}

func NewErrorResult(code int, errMsg ...string) JsonResult {
	result := JsonResult{Code: code, Msg: "server interal error"}
	if len(errMsg) > 0 {
		result.Msg = errMsg[0]
	}
	return result
}

func ContextFail(ctx iris.Context, statusCode int, format string, a ...interface{}) {
	err := HttpError{
		Code:   statusCode,
		Reason: fmt.Sprintf(format, a...),
	}
	//记录所有> = 500内部错误。
	if statusCode >= 500*100 {
		ctx.Application().Logger().Error(err)
	}
	ctx.StatusCode(statusCode)
	ctx.JSON(err)
	//没有下一个处理程序将运行。
	ctx.StopExecution()
}

// common error define
func ContextError(ctx iris.Context, msg ...string) {

	result := NewErrorResult(StatusInternalServerError, msg...)
	err := HttpError{
		Code:   result.Code,
		Reason: result.Msg,
	}
	//记录所有> = 500内部错误。
	if result.Code >= 500*100 {
		ctx.Application().Logger().Error(err)
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(result)
}

func Error(statusCode int, msg ...string) JsonResult {

	result := NewErrorResult(statusCode, msg...)

	return result
}

//
//
func ContextOk(ctx iris.Context, data interface{}, msg ...string) {
	ctx.StatusCode(iris.StatusOK)
	result := NewSuccessResult(data, msg...)
	ctx.JSON(result)
}

//
func Ok(data interface{}, msg ...string) JsonResult {
	result := NewSuccessResult(data, msg...)
	return result
}

// 401 error define
func Unauthorized(ctx iris.Context, msg string, data interface{}) {
	result := NewUnauthorizedResult(msg, data)
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(result)
}

//
//
func PaginationResult(rows interface{}, total int64) JsonResult {
	return NewSuccessResult(iris.Map{"rows": rows, "total": total})
	//return result
}

func OkPg(rows interface{}, total int64) JsonResult {
	return PaginationResult(rows, total)
}
