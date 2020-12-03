/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-26 09:22:43
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 18:16:44
 */
package commons

import (
	_ "strings"

)
 

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

