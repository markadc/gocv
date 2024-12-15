package rand

import (
	"fmt"
	"testing"
)

func TestRandomNumber(t *testing.T) {
	fmt.Printf("随机整数")
	for _ = range 10 {
		randomInt := RandomNumber(1, 10)
		fmt.Printf(" %d", randomInt)
	}
	fmt.Println()

}

func TestRandomFloat(t *testing.T) {
	fmt.Printf("随机浮点数")
	for _ = range 10 {
		randomFloat := RandomFloat()
		fmt.Printf(" %.2f ", randomFloat)
	}
	fmt.Println()
}
