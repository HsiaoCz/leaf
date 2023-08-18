package settings

import (
	"math/rand"
	"strconv"
	"time"
)

func VerificationCode() string {
	code := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
	return strconv.Itoa(code)
}
