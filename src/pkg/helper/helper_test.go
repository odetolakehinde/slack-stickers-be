package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringToIntPointer(t *testing.T) {
	result := 4
	tests := []struct {
		name   string
		value  string
		result int
	}{
		{
			name:   "Convert string to int pointer correctly",
			value:  "4",
			result: result,
		},
		{
			name:   "Failure to convert string to int pointer",
			value:  "NONE",
			result: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, *StringToIntPointer(test.value))
		})
	}
}

func Test_StringToBool(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		result bool
	}{
		{
			name:   "Convert string to bool correctly",
			value:  "true",
			result: true,
		},
		{
			name:   "Failure to convert string to bool",
			value:  "NONE",
			result: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, StringToBool(test.value))
		})
	}
}

func Test_StringToBoolPointer(t *testing.T) {
	result1 := true
	result2 := false

	tests := []struct {
		name   string
		value  string
		result *bool
	}{
		{
			name:   "Convert string to bool pointer correctly",
			value:  "true",
			result: &result1,
		},
		{
			name:   "Failure to convert string to bool pointer",
			value:  "NONE",
			result: &result2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, StringToBoolPointer(test.value))
		})
	}
}
