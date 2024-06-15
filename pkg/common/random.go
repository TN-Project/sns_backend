package common

import (
	"math/rand"
	"time"
)

// 10桁のランダムな整数を生成する
func GenerateRandomInt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(9000000000) + 1000000000
}
