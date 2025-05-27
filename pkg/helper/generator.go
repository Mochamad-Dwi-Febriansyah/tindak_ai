package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateComplaintNumber() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	randomPart := make([]byte, 6)
	for i := range randomPart {
		randomPart[i] = charset[rand.Intn(len(charset))]
	}

	datePart := time.Now().Format("20060102") // Format: YYYYMMDD
	return fmt.Sprintf("CMP-%s-%s", datePart, string(randomPart))
}
