/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-19 09:29:27
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-19 09:29:40
 */
package image

import (
	"bytes"
	"encoding/base64"
 	"image/jpeg"
	"image/png"
 	"os"
	"path/filepath"
 	"strings"

	"github.com/joshua-chen/go-commons/utils"

)

//
func SaveAsFile(imageBase64 string, storagePath string) (bool, error) {

	coI := strings.Index(imageBase64, ",")
	rawImage := imageBase64

	// Encoded Image DataUrl //
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))

	res := bytes.NewReader(unbased)
	path, _ := os.Getwd()

	if !strings.HasPrefix(storagePath, "/") {
		storagePath = "/" + storagePath
	}
	// Path to store the image //
	newpath := filepath.Join(path, storagePath)
	os.MkdirAll(newpath, os.ModePerm)
	uid, _ := utils.GetUUID()
 	imageType := strings.TrimSuffix(imageBase64[5:coI], ";base64")

	var f *os.File

	switch imageType {
	case "image/png":
		pngI, err := png.Decode(res)
		if err != nil {
			return false, err
		}

		f, _ = os.OpenFile(newpath+"/"+uid+".png", os.O_WRONLY|os.O_CREATE, 0777)
		png.Encode(f, pngI)
 
		break
	case "image/jpeg":
		jpgI, err := jpeg.Decode(res)
		if err != nil {
			return false, err
		}
		f, _ = os.OpenFile(newpath+"/"+uid+".jpg", os.O_WRONLY|os.O_CREATE, 0777)
		jpeg.Encode(f, jpgI, &jpeg.Options{Quality: 75})
 
		break
	}

	// if image is png this function will create the image from dataurl string //

	defer func() {
		f.Close()
	}()
	return true, nil
}
