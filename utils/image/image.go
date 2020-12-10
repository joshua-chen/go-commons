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
	_ "bytes"
	"encoding/base64"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/joshua-chen/go-commons/utils"

)

//
func SaveAsFile(imageBase64 string, storagePath string) (id string, name string, fullpath string, err error) {

	headerIndex := strings.Index(imageBase64, ",")
	rawImage := imageBase64[headerIndex+1:]

	// Encoded Image DataUrl //
	unbased, _ := base64.StdEncoding.DecodeString(rawImage)

	res := bytes.NewReader(unbased)
	path, _ := os.Getwd()
	newpath := storagePath
	if !strings.HasPrefix(storagePath, "./") {
		if !strings.HasPrefix(storagePath, "/") {
			storagePath = "/" + storagePath
		}
		newpath = filepath.Join(path, storagePath)
	}

	// Path to store the image //
	os.MkdirAll(newpath, os.ModePerm)
	id, _ = utils.GetUUID()
	imageTag := imageBase64[5:headerIndex]
	imageType := strings.TrimSuffix(imageTag, ";base64")

	var f *os.File

	switch imageType {
	case "image/png":
		pngI, err := png.Decode(res)
		if err != nil {
			return "", "", "", err
		}

		f, _ = os.OpenFile(newpath+"/"+id+".png", os.O_WRONLY|os.O_CREATE, 0777)
		png.Encode(f, pngI)
		break
	case "image/jpeg":
		jpgI, err := jpeg.Decode(res)
		if err != nil {
			return "", "", "", err
		}
		f, _ = os.OpenFile(newpath+"/"+id+".jpg", os.O_WRONLY|os.O_CREATE, 0777)
		jpeg.Encode(f, jpgI, &jpeg.Options{Quality: 75})
		break
	}

	defer f.Close()
	//f.Write(unbased)

	fullpath = f.Name()
	fileinfo, nil := f.Stat()
	name = fileinfo.Name()

	return id, name, fullpath, nil
}
