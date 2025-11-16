package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Util struct{}

func (Util) SliceUniqueString(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func (Util) SliceUniqueNumber(target []int64) (unique []int64) {
	m := map[int64]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

// 配列の中に特定の文字列が含まれているか
func (Util) ArrayStringContains(target []string, str string) bool {
	for _, v := range target {
		if v == str {
			return true
		}
	}
	return false
}

// 引数datetime文字列とタイムゾーン文字列を受け取りTime型に変換する関数
func (Util) ConvertDatetimeToLocationTime(datetime string, timezone string) (time.Time, error) {
	timeZone, _ := time.LoadLocation(timezone)
	locationTime, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, timeZone)
	if err != nil {
		log.Fatal(err.Error())
		return time.Time{}, err
	}
	return locationTime, nil
}

func (Util) TimeToUnixMilli(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

func (Util) LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("invalid private key data")
	}

	switch block.Type {
	case "RSA PRIVATE KEY": // PKCS#1
		return x509.ParsePKCS1PrivateKey(block.Bytes)

	case "PRIVATE KEY": // PKCS#8
		keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key, ok := keyInterface.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not RSA private key")
		}
		return key, nil

	default:
		return nil, fmt.Errorf("unsupported private key type: %s", block.Type)
	}
}

// DecryptRSA はRSAによる復号化をします
func (Util) DecryptRSA(priv *rsa.PrivateKey, ciphertext string) (string, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, decodedCiphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
