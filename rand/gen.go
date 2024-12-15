package rand

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 生成一个在 [min, max) 范围内的随机整数
func RandomNumber(min, max int) int {
	return r.Intn(max-min) + min
}

// 生成一个随机浮点数，范围在 [0.0, 1.0)
func RandomFloat() float64 {
	return r.Float64()
}
