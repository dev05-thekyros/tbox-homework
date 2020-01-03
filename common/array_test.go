package common

import "testing"

var stringInArrayTestCases = []struct {
	value    string   // input
	values   []string //input
	expected bool     // expected result
}{
	{"hung", []string{"hung", "dep", "trai"}, true},
	{"hung", []string{"123hung", "dep", "trai"}, false},
	{"hung", []string{"j", "xau", "trai"}, false},
}

func TestStringInArray(t *testing.T) {
	for _, tt := range stringInArrayTestCases {
		actual := StringInArray(tt.value, tt.values)
		if actual != tt.expected {
			t.Errorf("StringInArray(%s, %v): expected %v, actual %v ", tt.value, tt.values, tt.expected, actual)
		}
	}
}
