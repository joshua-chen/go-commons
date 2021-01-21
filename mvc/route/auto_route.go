/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-22 15:48:01
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 10:52:46
 */
package route

import (
	"os"
	//"strings"
	"bufio"

	//"github.com/joshua-chen/go-commons/config"
	"github.com/joshua-chen/go-commons/utils/file"
	"github.com/kataras/iris/v12"
	"github.com/kataras/golog"

)

func PartyAuto(app *iris.Application,  fn func(router iris.Party)) {

	dir, _ := os.Getwd()
	files, _ := file.ListDir(dir, "")
	for _, filename := range files{
		if filename != "" {
			file, err := os.Open(filename)
			if err != nil {
				golog.Fatal(err)
			}
			defer file.Close()
		
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				//fmt.Println(scanner.Text())

				golog.Info(scanner.Text())
			}
		
			if err := scanner.Err(); err != nil {
				golog.Fatal(err)
			}
		}
	}
	//api := app.Party(config.AppConfig.APIPrefix.Base)
	//api.PartyFunc(path, fn)
}
