/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 21:47:23
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:39:03
 */
package validator

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	"github.com/kataras/golog"

)

//
type Validator struct {
	Code    int
	Message string
	Err     error
}

var (
	instance *Validator
	lock     *sync.Mutex = &sync.Mutex{}
)

//
func Singleton() *Validator {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Validator{
				Code:    response.StatusValidatorFailed,
				Message: "验证未通过"}
		}
	}
	return instance
}

//
func Instance() *Validator {
	return Singleton()
}

//
func Error(err error, code ...int) {
	Instance().Error(err, code...)
}

//
func ErrorS(err string, code ...int) {
	Instance().ErrorS(err, code...)
}
func (e *Validator) Error(err error, code ...int) {
	if len(code) > 0 {
		e.Code = code[0]
	}
	e.Err = err
	errs := err.(validator.ValidationErrors)
	trans, _ := GetTranslator()
	msg := removeStructName(errs.Translate(trans))
	e.Message = fmt.Sprintf("%v ", msg)
	golog.Errorf("Error[%d]: %s", e.Code, e.Message)
	panic(err)
}

//
func (e *Validator) ErrorS(errMsg string, code ...int) {
	if len(code) > 0 {
		e.Code = code[0]
	}
	e.Message = errMsg
	golog.Errorf("Error[%d]: %s ", e.Code, errMsg)
	err := errors.New(errMsg)
	e.Err = err
	panic(err)
}

//
func (e *Validator) New() *validator.Validate {

	trans, _ := GetTranslator() //获取需要的语言
	validate := validator.New()
	zhtrans.RegisterDefaultTranslations(validate, trans)
	return validate
}

func GetTranslator(language ...string) (trans ut.Translator, found bool) {
	zh := zh.New() //中文翻译器
	en := en.New() //英文翻译器
	// 第一个参数是必填，如果没有其他的语言设置，就用这第一个
	// 后面的参数是支持多语言环境（
	// uni := ut.New(en, en) 也是可以的
	// uni := ut.New(en, zh, tw)
	uni := ut.New(en, zh)
	lang := "zh"
	if len(language) > 0 {
		lang = language[0]
	}
	return uni.GetTranslator(lang)
}

func removeStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}
