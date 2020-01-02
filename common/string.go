package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

// getRandNum returns a random number of size four
func GetRandNumberToString() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

func TokenGenerator() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Validate phone number
// Valid phone number must
// 		1. Length = 10
//		2. Only contains number
// 		3. Have valid prefix number
func ValidatePhone(phone string) bool {
	if len(phone) != 10 {
		return false
	}

	re := regexp.MustCompile(`^0\d{9}$`)
	isContainOnlyNumber := re.MatchString(phone)
	if !isContainOnlyNumber {
		return false
	}

	validPhonePrefix := []string{"090", "093", "089", "091", "094", "088", "096", "097", "098", "086", "092",
		"070", "079", "077", "076", "078", "083", "084", "085", "081", "082", "032", "033", "034", "035",
		"036", "037", "038", "039", "056", "058", "099", "059", "095", "052"}
	prefix := phone[:3]
	return StringInArray(prefix, validPhonePrefix)
}
