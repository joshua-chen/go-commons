/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 21:47:23
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:39:03
 */
package exception

import (
	"errors"
	"sync"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

)

type Exception struct {
	Code    int
	Message string
	Err     error
}

var (
	instance *Exception
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Exception {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Exception{
				Code:    iris.StatusInternalServerError * 100,
				Message: "内部服务器错误"}
		}
	}
	return instance
}
func Fatal(err error, code ...int) {
	Instance().Fatal(err, code...)
}
func FatalS(err string, code ...int) {
	Instance().FatalS(err, code...)
}
func (e *Exception) Fatal(err error, code ...int) {
	if len(code) > 0 {
		e.Code = code[0]
	}
	msg := err.Error()
	e.Err = err
	e.Message = msg
	golog.Errorf("Fatal[%d]: %s", e.Code, msg)
	panic(err)
}
func (e *Exception) FatalS(errMsg string, code ...int) {
	if len(code) > 0 {
		e.Code = code[0]
	}
	e.Message = errMsg
	golog.Errorf("Fatal[%d]: %s ", e.Code, errMsg)
	err := errors.New(errMsg)
	e.Err = err
	panic(err)
}
func (e *Exception) Error(err error) {

}
func (e *Exception) ErrorS(err string) {

}
