package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/hkdf"
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
	return s[8:24]
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

func DecryptHkdf(secret []byte, salt []byte, info []byte){
	hash := sha256.New

	// Non-secret salt, optional (can be nil).
	// Recommended: hash-length random value.
	//salt := make([]byte, hash().Size())
	//if _, err := rand.Read(salt); err != nil {
	//	panic(err)
	//}

	// Non-secret context info, optional (can be nil).
	//info := []byte("hkdf example")

	// Generate three 128-bit derived keys.
	hkdfObj := hkdf.New(hash, secret, salt, info)

	var keys [][]byte
	for i := 0; i < 3; i++ {
		key := make([]byte, 16)
		if _, err := io.ReadFull(hkdfObj, key); err != nil {
			panic(err)
		}
		keys = append(keys, key)
	}

	for i := range keys {
		fmt.Printf("Key #%d: %v\n", i+1, !bytes.Equal(keys[i], make([]byte, 16)))
	}
}