package random

import (
	"math/rand"
	"time"
)

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 10桁のランダムな整数を生成する
func GenerateRandomInt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9000000000) + 1000000000
}

// n桁のランダムな文字列を生成する
func MakeRandomStringId(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}
