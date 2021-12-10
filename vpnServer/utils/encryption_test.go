package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

//go test -v utils/*

func TestEnAesgcm(t *testing.T) {
	t.Log("aes256加密")
	origData := []byte("这是我的原始数据")
	key := []byte("01234567891234560123456789123456")
	iv := []byte("0123456789123456")

	log.Println("------------------ CFB模式 --------------------")
	encrypted := AesEncryptCFB(origData, key, iv)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCFB(encrypted, key, iv)

	t.Log(fmt.Sprintf("经过加密再解密后的数据 [%v]", string(decrypted)))
}

//func TestMd5(t *testing.T) {
//	t.Log("md5加密")
//	res := GetMD5Encode("abcdefg")
//	t.Log(fmt.Sprintf("加密后数据为 %v", res))
//}