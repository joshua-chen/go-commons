/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-27 15:10:24
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 16:34:01
 */ 

 package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"

)

 type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	//Password string `json:"password"`
	//User models.User `json:"user"`
	jwt.StandardClaims
}