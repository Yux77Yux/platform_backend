package jwt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// AES 加密密钥（生产环境请存到环境变量）
var aesKey = []byte("0123456789abcdef0123456789abcdef") // 32字节，AES-256

// 加密 RefreshToken
func EncryptRefreshToken(plainText string) (string, error) {
	return encryptAES(plainText, aesKey)
}

// 解密并验证 RefreshToken
func VerifyRefreshToken(encryptedToken string) (string, error) {
	return decryptAES(encryptedToken, aesKey)
}

// **AES CBC 加密**
func encryptAES(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 生成 IV（初始化向量）
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// PKCS7 填充
	paddedText := PKCS7Padding([]byte(plainText), aes.BlockSize)

	// 加密
	cipherText := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, paddedText)

	// 返回 IV + 密文（Base64 编码）
	return base64.StdEncoding.EncodeToString(append(iv, cipherText...)), nil
}

// **AES CBC 解密**
func decryptAES(encText string, key []byte) (string, error) {
	encBytes, err := base64.StdEncoding.DecodeString(encText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 获取 IV 和密文
	if len(encBytes) < aes.BlockSize {
		return "", errors.New("密文长度不正确")
	}
	iv := encBytes[:aes.BlockSize]
	ciphertext := encBytes[aes.BlockSize:]

	// 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 去掉填充
	plainText, err := PKCS7UnPadding(ciphertext)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// **PKCS7 填充**
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// **PKCS7 取消填充**
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("解密数据为空")
	}
	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("填充格式错误")
	}
	return data[:length-padding], nil
}
