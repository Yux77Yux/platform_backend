package repository

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/argon2"
)

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
	salt := make([]byte, 16) // 16字节长度的盐值
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}
