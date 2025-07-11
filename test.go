package main

import (
	"HOTEL-REGISTRY_API/helpers"
	"crypto/rand"
	"math/big"
)

func main() {

}

// Generate OTP
func GenerateOTP(pDebug *helpers.HelperStruct, length int) (string, error) {
	pDebug.Log(helpers.Statement, "GenerateOTP (+)")

	const digits = "0123456789"
	otp := ""
	for i := 0; i < length; i++ {
		randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", helpers.ErrReturn(err)
		}
		otp += string(digits[randomInt.Int64()])
	}
	pDebug.Log(helpers.Statement, "GenerateOTP (-)")
	return otp, nil
}
