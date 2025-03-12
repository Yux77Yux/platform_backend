package tools

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/argon2"

	"fmt"
)

// 验证
func VerifyPassword(passwordInStore, password string) (bool, error) {
	saltHex, hashHex, err := SplitHash(passwordInStore)
	if err != nil {
		return false, err
	}

	salt, err := hex.DecodeString(saltHex)
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	storedHash, err := hex.DecodeString(hashHex)
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	// 使用提取的盐值重新计算密码的哈希
	computedHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// 比较存储的哈希和计算的哈希是否一致
	return subtle.ConstantTimeCompare(storedHash, computedHash) == 1, nil
}

func SplitHash(hashWithSalt string) (string, string, error) {
	parts := strings.Split(hashWithSalt, ":")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid hash format")
	}
	return parts[0], parts[1], nil
}

// 加密
func HashPassword(password string) (string, error) {
	// 生成盐值
	salt, err := GenerateSalt()
	if err != nil {
		return "", err
	}

	// 使用 Argon2ID 进行哈希处理
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// 返回盐值和哈希结果，格式化为十六进制字符串
	return fmt.Sprintf("%s:%s", hex.EncodeToString(salt), hex.EncodeToString(hash)), nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16) // 16字节长度的盐值
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}
