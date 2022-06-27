package utils

import (
	"math/rand"
	"time"
)

func GenerateClassCode() string {
	rand.Seed(time.Now().Unix())
	length := 6

	dec := []int{48, 65, 97}
	ran_str := make([]byte, length)

	for i := 0; i < length; i++ {
		d := dec[rand.Intn(3)]
		if d == dec[0] {
			ran_str[i] = byte(d + rand.Intn(10))
		} else {
			ran_str[i] = byte(d + rand.Intn(25))
		}
	}

	return string(ran_str)
}
