package main

import (
	"log"
)

func main() {
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 20; j++ {
			if i == 4 && j == 2 {
				break
			}
			log.Printf("A(%d,%d) == %d, in cache: %d", i, j, AckermannWithCache(i, j), len(ackermannCache))
		}
	}
}

var ackermannCache = make(map[[2]int]int)

func AckermannWithCache(m, n int) int {
	if m == 0 {
		return n + 1
	}

	if n == 0 {
		return AckermannWithCache(m-1, 1)
	}

	// Check cache
	key := [2]int{m, n}
	if val, ok := ackermannCache[key]; ok {
		return val
	}

	// Calculate
	val := AckermannWithCache(m-1, AckermannWithCache(m, n-1))
	ackermannCache[key] = val
	return val
}

func Ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	}

	if n == 0 {
		return Ackermann(m-1, 1)
	}

	return Ackermann(m-1, Ackermann(m, n-1))
}
