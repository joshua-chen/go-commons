/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-19 09:29:27
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-19 09:29:40
 */
package json

import (
	"encoding/json"
	"os"

)

func GetJson(path string, s interface{}) interface{} {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s)
	if err != nil {
		panic(err.Error())
	}
	return &s
}

func ToJson(obj interface{}) (string, error) {

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
