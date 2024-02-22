package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	const FROM int32 = 1000000000
	const TO int32 = 1000100000
	var limit int32 = int32(math.Sqrt(float64(TO)))
	var a [TO - FROM]int8

	var start = time.Now().UnixMilli()

	for n := FROM; n < TO; n++ {
		for i := 2; i < int(limit); i++ {
			if int(n)%i == 0 {
				a[n-FROM] = 1
				break
			}
		}
	}

	var elapsed = time.Now().UnixMilli() - start

	for n := FROM; n < TO; n++ {
		if a[n-FROM] == 0 {
			fmt.Println(n)
		}
	}

	fmt.Println("time = ", elapsed)
}
