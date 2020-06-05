/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:04:32
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 11:55:44
 */
package response

type HttpError struct {
    Code   int    `json:"code"`
    Reason string `json:"reason"`
}



type Result struct {
	Code    int      `json:"code"`
	Msg     string      `json:"msg"`
	Error    string      `json:"error"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

