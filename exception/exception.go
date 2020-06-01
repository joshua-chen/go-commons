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

func (e *Exception) Fatal(err error) {
	golog.Errorf("Fatal: %s",err.Error())
	panic(err)
}
func (e *Exception) FatalS(err string) {
	golog.Errorf("Fatal: %s",err)
	panic(errors.New(err))
}
func (e *Exception) Error(err error) {

}
func (e *Exception) ErrorS(err string) {

}
