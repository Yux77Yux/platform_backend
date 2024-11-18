package repository

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/argon2"
)

// // 验证密码是否匹配
// func verifyPassword(storedHash, password string) (bool, error) {
// 	// 分离存储的盐值和哈希值
// 	parts := strings.Split(storedHash, ":")
// 	if len(parts) != 2 {
// 		return false, errors.New("invalid stored hash format")
// 	}

// 	// 解码盐值和哈希值
// 	salt, err := hex.DecodeString(parts[0])
// 	if err != nil {
// 		return false, fmt.Errorf("failed to decode salt: %w", err)
// 	}
// 	expectedHash, err := hex.DecodeString(parts[1])
// 	if err != nil {
// 		return false, fmt.Errorf("failed to decode hash: %w", err)
// 	}

// 	// 使用相同的参数重新计算哈希
// 	computedHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

// 	// 比较哈希值
// 	return compareHashes(expectedHash, computedHash), nil
// }

// func compareHashes(a, b []byte) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	result := 0
// 	for i := range a {
// 		result |= int(a[i] ^ b[i])
// 	}
// 	return result == 0
// }

func hashPassword(password string) (string, error) {
	// 生成盐值
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	// 使用 Argon2ID 进行哈希处理
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// 返回盐值和哈希结果，格式化为十六进制字符串
	return fmt.Sprintf("%s:%s", hex.EncodeToString(salt), hex.EncodeToString(hash)), nil
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16) // 推荐16字节长度的盐值
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}

// 简单解密函数
func decryptWithTimestamp(encryptedData string) (string, error) {
	const maxAgeSeconds int64 = 70 // 默认有效时间为 70 秒

	keys := [][]byte{
		[]byte("c3a9ffe7f1385244f66a94070e3cc96c490cc5af13bc403707376b0ad8b36013"),
		[]byte("11b6c266ac3551db12acc80acca1051547f428e4936de8fa6e83e4acf84ed542"),
		[]byte("3d9954af0b271632abfda8288ca4ea55f24e3d3568345070810433e403820ad7"),
	}

	// Base64 解码
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	// 获取 IV、密文和 HMAC 校验码
	blockSize := aes.BlockSize
	if len(ciphertext) < blockSize+sha256.Size {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize : len(ciphertext)-sha256.Size]
	receivedHMAC := ciphertext[len(ciphertext):]

	// 逐个尝试密钥进行解密
	for _, key := range keys {
		// 计算 HMAC 校验码
		hmac := hmac.New(sha256.New, key) // 使用当前密钥
		hmac.Write(ciphertext)
		expectedHMAC := hmac.Sum(nil)

		// 使用 subtle.ConstantTimeCompare 来进行 HMAC 校验码的比较，避免泄露信息
		if subtle.ConstantTimeCompare(receivedHMAC, expectedHMAC) == 0 {
			// 如果 HMAC 校验失败，则尝试下一个密钥
			continue
		}

		// 创建 AES 解密块
		block, err := aes.NewCipher(key)
		if err != nil {
			return "", err
		}

		// 创建 CBC 解密模式
		mode := cipher.NewCBCDecrypter(block, iv)

		// 解密数据
		plaintext := make([]byte, len(ciphertext))
		mode.CryptBlocks(plaintext, ciphertext)

		// 去除 PKCS7 填充
		padding := int(plaintext[len(plaintext)-1])
		plaintext = plaintext[:len(plaintext)-padding]

		// 分离时间戳和原始数据
		data := string(plaintext)
		parts := bytes.SplitN([]byte(data), []byte(":"), 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid ciphertext format")
		}

		// 校验时间戳
		timestamp, err := strconv.ParseInt(string(parts[0]), 10, 64)
		if err != nil {
			return "", err
		}
		if time.Now().Unix()-timestamp > maxAgeSeconds {
			return "", fmt.Errorf("ciphertext expired")
		}

		// 返回解密后的数据
		return string(parts[1]), nil
	}

	// 如果所有密钥都尝试失败
	return "", errors.New("decryption failed with all keys")
}
