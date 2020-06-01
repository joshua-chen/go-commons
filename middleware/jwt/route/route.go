/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-22 15:36:44
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 16:52:59
 */ 
package route

import (
	_ "commons/middleware/jwt"
	_ "commons/mvc"
	_ "commons/mvc/models"
	_ "fmt"
	_ "log"
	_ "time"

	_ "github.com/dgrijalva/jwt-go/request"
	_ "github.com/iris-contrib/middleware/cors"
	_ "github.com/kataras/iris/v12"

)

 
