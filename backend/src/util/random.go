package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomFloat(min, max float64) float64 {
	randomInRange := min + rand.Float64()*(max-min)
	return randomInRange
}

const alphabet = "abcdefghijklmnopqrstuvxyz"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	count := len(alphabet)

	for i := 0; i < n; i++ {
		j := alphabet[rand.Intn(count)]
		sb.WriteByte(j)
	}
	return sb.String()
}

// return a string to be the owner name
func RandomOwner() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	min := int64(0)
	max := int64(10000)
	randomValue := RandomInt(min, max)
	return randomValue
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
