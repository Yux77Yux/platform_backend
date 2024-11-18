package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"time"
)

// 简单加密函数
func encryptWithTimestamp(plaintext string) (string, error) {
	keys := [][]byte{
		[]byte("c3a9ffe7f1385244f66a94070e3cc96c490cc5af13bc403707376b0ad8b36013"),
		[]byte("11b6c266ac3551db12acc80acca1051547f428e4936de8fa6e83e4acf84ed542"),
		[]byte("3d9954af0b271632abfda8288ca4ea55f24e3d3568345070810433e403820ad7"),
	}

	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
	if err != nil {
		log.Println("error generating random index:", err)
		return plaintext, err
	}

	key := keys[randomIndex.Int64()]

	// 添加时间戳
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	plaintext = timestamp + ":" + plaintext // 将时间戳和数据拼接

	// 创建 AES 加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 生成随机 IV（初始化向量）
	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	// 填充数据
	dataBytes := []byte(plaintext)
	padding := block.BlockSize() - len(dataBytes)%block.BlockSize()
	paddedData := append(dataBytes, make([]byte, padding)...)

	// 创建加密结果
	ciphertext := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedData)

	// 计算 HMAC 校验码
	hmac := hmac.New(sha256.New, key) // HMAC 使用相同的密钥
	hmac.Write(ciphertext)
	hash := hmac.Sum(nil)

	// 合并 IV、密文和 HMAC
	encrypted := append(iv, ciphertext...)
	encrypted = append(encrypted, hash...)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
