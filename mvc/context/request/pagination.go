/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-25 16:38:18
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 11:35:19
 */ 
package request

import (
	_"errors"

	"github.com/kataras/iris/v12"

)

// bootstraptable 分页参数
type Pagination struct {
	PageNum int //当前看的是第几页
	PageSize   int //每页显示多少条数据

	// 用于分页设置的参数
	Offset int
	Limit int

	SortName  string //用于指定的排序
	SortOrder string // desc或asc

	// 时间范围
	StartDate string
	EndDate   string

	ID int64 // 公用的特殊参数
}

func NewPagination(ctx iris.Context) (*Pagination) {
	pageNum := ctx.URLParamIntDefault("pageNum",-1)
	pageSize:= ctx.URLParamIntDefault("pageSize",-1)
	offset := ctx.URLParamIntDefault("offset",-1)
	limit := ctx.URLParamIntDefault("limit",-1)
	sortName := ctx.URLParamDefault("sortName","-1")
	sortOrder := ctx.URLParamDefault("sortOrder","-1")
	
	 
	var page Pagination
	if(offset != -1 && limit != -1){
		page = Pagination{
			SortName:   sortName,
			SortOrder:  sortOrder,
			Limit: limit,
			Offset: offset,
		}
		return &page
	}
	
	page = Pagination{
		PageNum: pageNum,
		PageSize: pageSize,
		SortName:   sortName,
		SortOrder:  sortOrder,
	}
	page.set()
	
	return &page
}

// 设置分页参数
func (p *Pagination) set() {
	if p.PageNum < 1 {
		p.PageNum = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 1
	}

	p.Offset = (p.PageNum - 1) * p.PageSize
	p.Limit = p.PageSize
}

 

