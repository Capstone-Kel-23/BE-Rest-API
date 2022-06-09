package utils

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/validation/utils"
	"math/rand"
	"strconv"
)

func GenerateNumberInvoice() string {
	x2 := rand.NewSource(100)
	y2 := rand.New(x2)

	randomString := utils.GenerateRandomString(5)

	return strconv.Itoa(y2.Int()) + randomString
}
