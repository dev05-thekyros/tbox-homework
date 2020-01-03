package common

import "testing"

var validatePhoneTestCases = []struct {
	phoneNumber string // input,
	expected    bool   // expected result
}{
	{"0931317941", true},
	{"+84931317941", false}, // Only accept 1 type
	{"abcalhjadsldaksl", false},
	{"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111", false},
	{"0931313131313131313131313131313", false},
	{"#$@#$@#AFDSFSDFDSF@#$@$#@$@#$@#4", false},
	{"phonenumber", false},
}

func TestValidatePhone(t *testing.T) {
	for _, tt := range validatePhoneTestCases {
		actual := ValidatePhone(tt.phoneNumber)
		if actual != tt.expected {
			t.Errorf("Fib(%s): expected %v, actual %v ", tt.phoneNumber, tt.expected, actual)
		}
	}
}
