package utils

import (
	"crypto/rand"
	"math/big"
)

// utils/otp.go (continued)
func GenerateOTP() string {
	const otpChars = "0123456789"
	const otpLength = 6

	otp := make([]byte, otpLength)
	for i := range otp {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(otpChars))))
		otp[i] = otpChars[randomIndex.Int64()]
	}
	return string(otp)
}
