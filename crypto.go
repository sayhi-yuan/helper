package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// ====================== Base64 ===========================
func Base64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) (string, error) {
	if res, err := base64.StdEncoding.DecodeString(data); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%x", res), nil
	}
}

func Base64Url(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func Base64UrlDecode(data string) (string, error) {
	if res, err := base64.URLEncoding.DecodeString(data); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%x", res), nil
	}
}

// ====================== Md5 ===========================
type md5Crypto struct {
	bitSize uint8 `desc:"返回的加密字符串的长度"`
}

type md5CryptoOption func(m *md5Crypto)

type Md5BitSize uint8

const (
	Md5BitSize16 Md5BitSize = 1
	Md5BitSize32 Md5BitSize = 2
)

func (m Md5BitSize) Index() uint8 {
	return uint8(m) * 16
}

func Md5() *md5Crypto {
	return &md5Crypto{}
}

func (m *md5Crypto) Encrypt(data string, options ...md5CryptoOption) string {
	for _, option := range options {
		option(m)
	}

	h := md5.New()
	h.Write([]byte(data))

	if m.bitSize == 16 {
		return fmt.Sprintf("%x", h.Sum(nil))[8:24]
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func Md5CryptoOptionBitSize(size Md5BitSize) md5CryptoOption {
	return func(m *md5Crypto) {
		m.bitSize = size.Index()
	}
}

// ====================== Md5 ===========================
func Sha1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))

	return fmt.Sprintf("%x", h.Sum(nil))
}
