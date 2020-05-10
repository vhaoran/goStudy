package test

import (
	"log"
	"math"
	"math/rand"
	"testing"
)

func Test_Rand_call(t *testing.T) {
	var start float64 = 10
	var sum float64 = 0
	n := 10
	for i := 0; i < 10; i++ {
		remain := n - i
		f := RandFloatN(start, remain)
		if remain == 1 {
			f = start
		} else {
			start -= f
		}

		sum += f
		log.Println("rand_test.go->", f)
	}

	log.Println("sum", sum)
}

func RandFloatN(max float64, n int) float64 {
	f := rand.Float64() * max / float64(n)

	for {
		if f < max/float64(n) {
			f += max / float64(n)

			f = math.Floor(f*100) / 100
		}

		if f > max/float64(n/2) {
			f += max - f

			f = math.Floor(float64(n))
			f = math.Floor(f*100) / 100
		}

		return f
	}

	return 0
}
