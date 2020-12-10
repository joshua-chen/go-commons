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
	"os"
	"path/filepath"
	"strings"
	_ "strings"

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

func IsAbsolutePath(path string) bool {
	name := filepath.VolumeName(path)
	return name != ""
}

func GetRunDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		golog.Fatal(err)
	}
	//golog.Info("[dir]==============================> ", dir)

	return dir

}

func GetCurrentDir() string {
	dir, _ := os.Getwd()
	golog.Info("[current dir]==============================> ", dir)
	return dir
}

func GetAbsolutePath(path string) string {
	if !IsAbsolutePath(path) {
		path = strings.ReplaceAll(path, `./`, "")
		path = strings.ReplaceAll(path, `.\`, "")
		if path[0:1] != `\` {
			path = `\` + path
		}
		path = GetCurrentDir() + path
		path = strings.ReplaceAll(path, `\`, "/")

	}
	return path
}

// @Title  PathExists
// @Description  路径是否存在
// @Author  joshua  ${DATE} ${TIME}
// @Update  joshua  ${DATE} ${TIME}
func PathExisted(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

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

 
