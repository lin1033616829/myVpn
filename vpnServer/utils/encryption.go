package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func GetMD5Encode(str string) string {
	data := []byte(str)
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	fmt.Println(fmt.Sprintf("加密后数据 s=[%v]", s))
	return s
}

func Get16MD5encode(str string) string {
	data := []byte(str)
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	fmt.Println(fmt.Sprintf("加密后数据 s=[%v]", s))
	return s
}

// AesEncryptCFB =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte, iv []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//encrypted = make([]byte, aes.BlockSize+len(origData))
	encrypted = make([]byte, len(origData))
	//iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted, origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte, iv []byte) (decrypted []byte, err error) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	//iv := encrypted[:aes.BlockSize]
	//encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}
