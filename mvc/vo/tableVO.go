/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-28 17:56:13
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 17:58:46
 */ 
package vo
 type TableVO struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}
