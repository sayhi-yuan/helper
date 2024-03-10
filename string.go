package helper

import (
	"math/rand"
	"time"
	"unicode"
)

// 获取字符串中的汉子个数
func StrHanLen(content string) int {
	count := 0
	for _, word := range content {
		if unicode.Is(unicode.Han, word) {
			count++
		}
	}

	return count
}

// 获取字符串长度,中文算一个字符
func StrLen(content string) int {
	return len(content) - StrHanLen(content)
}

// StrRandom 随机生成字符串
func StrRandom(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seedRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = str[seedRand.Intn(len(str))]
	}

	return string(b)
}
