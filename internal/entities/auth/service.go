package auth

import (
	"math/rand"
	"strconv"
	"strings"
)

func RandomCode(length int) string {
	var code = []string{}

	for range length {
		n := rand.Intn(10)
		code = append(code, strconv.Itoa(n))
	}

	return strings.Join(code, "")
}
