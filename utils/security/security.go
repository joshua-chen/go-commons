/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-25 16:24:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 15:13:51
 */ 
package security

import (
	_"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"errors"

)

// Md5 returns the MD5 checksum string of the data.
func Md5(b []byte) string {
	checksum := md5.Sum(b)
	return hex.EncodeToString(checksum[:])
}


// AESDecrypt decrypts a piece of data.
// The cipherkey argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.


func PadData(d []byte, bs int) []byte {
	padedSize := ((len(d) / bs) + 1) * bs
	pad := padedSize - len(d)
	for i := len(d); i < padedSize; i++ {
		d = append(d, byte(pad))
	}
	return d
}

func RemovePad(r []byte) ([]byte, error) {
	l := len(r)
	if l == 0 {
		return []byte{}, errors.New("input []byte is empty")
	}
	last := int(r[l-1])
	pad := r[l-last : l]
	isPad := true
	for _, v := range pad {
		if int(v) != last {
			isPad = false
			break
		}
	}
	if !isPad {
		return r, errors.New("remove pad error")
	}
	return r[:l-last], nil
}
