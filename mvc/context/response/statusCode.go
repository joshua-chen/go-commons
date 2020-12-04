/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:04:32
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 11:55:44
 */
package response

import "github.com/kataras/iris/v12"

//状态系数
const StatusCoefficient = 100 

const StatusInternalServerError = iris.StatusInternalServerError * StatusCoefficient

const StatusUnauthorized = iris.StatusUnauthorized * StatusCoefficient

const StatusOK = iris.StatusOK * StatusCoefficient

const StatusNotFound = iris.StatusNotFound * StatusCoefficient

const StatusExpectationFailed = iris.StatusExpectationFailed * StatusCoefficient

//验证状态码基数
const StatusValidateCodeBase = 600

const StatusValidateFailed = StatusValidateCodeBase * StatusCoefficient

const StatusTokenParamEmpty = StatusInternalServerError + 999
