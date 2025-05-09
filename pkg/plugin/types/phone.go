package types

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql/driver"
	"encoding/base64"
	"errors"
	"io"
)

type Phone string

var encryptionKey = []byte("passphrasewhichneedstobe32bytes!")

// Scan implements the sql.Scanner interface
func (p *Phone) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var encrypted string
	switch v := value.(type) {
	case string:
		encrypted = v
	case []byte:
		encrypted = string(v)
	default:
		return errors.New("invalid type for MobileNumber")
	}

	// 解密逻辑
	decrypted, err := decrypt(encrypted)
	if err != nil {
		return err
	}

	*p = Phone(decrypted)
	return nil
}

// Value implements the driver.Valuer interface
func (p Phone) Value() (driver.Value, error) {
	// 加密逻辑
	return encrypt(string(p))
}

// 加密函数
func encrypt(plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// 解密函数
func decrypt(ciphertext string) (string, error) {
	if len(ciphertext) == 0 {
		return "", nil
	}
	decoded, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	if len(decoded) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := decoded[:aes.BlockSize]
	decrypted := decoded[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)

	return string(decrypted), nil
}
