/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-26 09:22:43
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 18:16:44
 */
 package path

 import (
 	 "fmt"
	 "os"
	 "path/filepath"
	 "strings"

 	 "github.com/kataras/golog"

 )
  
 
 func IsFullPath(path string) bool {
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
 
 func GetFullPath(path string) string {
	 if !IsFullPath(path) {
		 path = strings.ReplaceAll(path, `./`, "")
		 path = strings.ReplaceAll(path, `.\`, "")
		 if path!=""&&path[0:1] != `\` {
			 path = `\` + path
		 }
		 path = GetCurrentDir() + path
		 path = strings.ReplaceAll(path, `\`, "/")
 
	 }
	 return path
 }
 
 // @Title  PathExisted
 // @Description  路径是否存在
 // @Author  joshua  ${DATE} ${TIME}
 // @Update  joshua  ${DATE} ${TIME}
 func PathExisted(path string) bool {
	 existed := true
	 if _, err := os.Stat(path); os.IsNotExist(err) {
		 existed = false
	 }
	 return existed
 }
 
 
 func IsDir(name string) bool {
	 if info, err := os.Stat(name); err == nil {
		 return info.IsDir()
	 }
	 return false
 }
 
 func MakeDir(dir string) error {
	 if !PathExisted(dir) {
		 if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			 fmt.Println("MakeDir failed:", err)
			 return err
		 }
	 }
	 return nil
 }
  
 