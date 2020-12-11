/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-26 09:22:43
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 18:16:44
 */
package utils

import (
	"github.com/google/uuid"
	"github.com/kataras/golog"

)

// s 中是否以 prefix 开始
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// s 中是否以 suffix 结尾
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

//
func GetUUID() (string, error) {
	u2, err := uuid.NewUUID()
	if err != nil {
		golog.Warnf("Something went wrong: %s", err)
		return "", err
	}
	return u2.String(), nil
}
