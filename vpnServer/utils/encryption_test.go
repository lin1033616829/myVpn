package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/hkdf"
	"io"
	"log"
	"math/rand"
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
	decrypted, err := AesDecryptCFB(encrypted, key, iv)
	if err != nil {
		t.Log(fmt.Sprintf("AesDecryptCFB err %v", err))
	}

	t.Log(fmt.Sprintf("经过加密再解密后的数据 [%v]", string(decrypted)))
}

func TestHkdf(t *testing.T) {
	t.Log("开始测试Hkdf-------------")
	// Underlying hash function for HMAC.
	hash := sha256.New

	// Cryptographically secure master secret.
	secret := []byte{0x00, 0x01, 0x02, 0x03} // i.e. NOT this.

	// Non-secret salt, optional (can be nil).
	// Recommended: hash-length random value.
	salt := make([]byte, hash().Size())
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	// Non-secret context info, optional (can be nil).
	info := []byte("hkdf example")

	// Generate three 128-bit derived keys.
	hkdf := hkdf.New(hash, secret, salt, info)

	var keys [][]byte
	for i := 0; i < 3; i++ {
		key := make([]byte, 16)
		if _, err := io.ReadFull(hkdf, key); err != nil {
			panic(err)
		}
		keys = append(keys, key)
	}

	fmt.Println(fmt.Sprintf("info %v", keys))

	for i := range keys {
		fmt.Printf("Key #%d: %v\n", i+1, !bytes.Equal(keys[i], make([]byte, 16)))
	}

}