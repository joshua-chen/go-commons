/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-26 18:20:47
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 10:35:35
 */
package context

import (
	"github.com/joshua-chen/go-commons/mvc/context/request"

	"github.com/kataras/iris/v12"
)

type HttpContext interface {
	iris.Context
}

type Handler func(ctx HttpContext)
type Filter func(ctx HttpContext) bool

func PagingParams() *request.Pagination {
	return &request.Pagination{}
}
