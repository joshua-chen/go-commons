/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-25 16:24:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 15:16:08
 */
package aes

import (
	"bytes"
	aesutil "crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/joshua-chen/go-commons/utils/security"
)

/**
AES加密解密
使用CBC模式+PKCS7 填充方式实现AES的加密和解密
*/
const key = "ARRSWdczx13213EDDSWQ!!@W"

// 校验密码
func CheckPWD(password, enPassword string) bool {
	de := AESDecrypt(enPassword)
	if password == de {
		return true
	}
	return false
}

// -----------------------------------------------------------
// ----------------------- 解密 ------------------------------
// -----------------------------------------------------------
// 先base64转码，再解密
func AESDecrypt(baseStr string) string {
	crypted, err := base64.StdEncoding.DecodeString(baseStr)
	if err != nil {
		fmt.Println("base64 encoding 错误")
	}

	block, _ := aesutil.NewCipher([]byte(key))
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, []byte(key)[:blockSize])
	originPWD := make([]byte, len(crypted))
	blockMode.CryptBlocks(originPWD, crypted)
	originPWD = pkcs7_unPadding(originPWD)
	return string(originPWD)
}
func AESDecrypt2(cipherkey, ciphertext []byte) ([]byte, error) {
	block, err := aesutil.NewCipher(cipherkey)
	if err != nil {
		return nil, err
	}
	src := make([]byte, hex.DecodedLen(len(ciphertext)))
	_, err = hex.Decode(src, ciphertext)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	r := make([]byte, len(src))
	dst := r
	for len(src) > 0 {
		block.Decrypt(dst, src)
		src = src[bs:]
		dst = dst[bs:]
	}
	return security.RemovePad(r)
}

// 补码
func pkcs7_unPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}

// -----------------------------------------------------------
// ----------------------- 加密 ------------------------------
// -----------------------------------------------------------
// 加密后再base64编码成string
func AESEncrypt(originPWD []byte) string {
	//获取block块
	block, _ := aesutil.NewCipher([]byte(key))
	//补码
	originPWD = pkcs7_padding(originPWD, block.BlockSize())
	//加密模式，
	blockMode := cipher.NewCBCEncrypter(block, []byte(key)[:block.BlockSize()])
	//创建明文长度的数组
	crypted := make([]byte, len(originPWD))
	//加密明文
	blockMode.CryptBlocks(crypted, originPWD)

	return base64.StdEncoding.EncodeToString(crypted)
}

// AESEncrypt encrypts a piece of data.
// The cipherkey argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func AESEncrypt2(cipherkey, src []byte) []byte {
	block, err := aesutil.NewCipher(cipherkey)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	src = security.PadData(src, bs)
	r := make([]byte, len(src))
	dst := r
	for len(src) > 0 {
		block.Encrypt(dst, src)
		src = src[bs:]
		dst = dst[bs:]
	}
	dst = make([]byte, hex.EncodedLen(len(r)))
	hex.Encode(dst, r)
	return dst
}

// 补码
func pkcs7_padding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padtext...)
}
