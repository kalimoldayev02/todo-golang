package utils

import (
	"crypto/sha1"
	"fmt"
	"todo-golang/pkg/config"
)

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(config.GetInstance().Salt)))
}
