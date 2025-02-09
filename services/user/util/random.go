package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

func RandomTimestamp() time.Time {
	return time.Now().Add(time.Duration(RandomInt(-1000, 0)) * time.Hour)
}

func RandomPassword() string {
	return RandomString(10)
}

func RandomPhoneNumber() string {
	var letter = []rune("0123456789")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}