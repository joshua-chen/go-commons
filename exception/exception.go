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

)

type Exception struct {
	Code    int
	Message string
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
			instance = &Exception{}
		}
	}
	return instance
}
func Fatal(err error, code ...int) {
	Instance().Fatal(err, code...)
}
func Fatalf(err string, code ...int) {
	Instance().Fatalf(err, code...)
}
func (e *Exception) Fatal(err error, code ...int) {
	if len(code) > 0 {
		e.Code = code[0]
	}
	msg := err.Error()
	e.Message = msg
	golog.Errorf("Fatal: %s", msg)
	panic(err)
}
func (e *Exception) Fatalf(err string, code ...int) {
	if len(code) > 0 {
		e.Code = code[0]
	}
	e.Message = err
	golog.Errorf("Fatal: %s", err)
	panic(errors.New(err))
}
func (e *Exception) Error(err error) {

}
func (e *Exception) ErrorS(err string) {

}
