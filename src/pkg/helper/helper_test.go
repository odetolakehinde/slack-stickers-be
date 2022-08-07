package helper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Test_StringToDatetime(t *testing.T) {
	timeValue1 := "2021-01-13 01:57:00"
	timeValue2 := "09/5/2021 4:46"
	timeValue3 := "9/05/2021 4:46"
	timeValue4 := "9/5/2021 4:46"
	timeValue5 := "9-5-2021 4:46 PM"
	timeValue6 := "09-05-2021 4:46"
	tests := []struct {
		name        string
		timeString  string
		shouldError bool
	}{
		{
			name:        "successful case 1",
			timeString:  timeValue1,
			shouldError: false,
		},
		{
			name:        "failure case 1",
			timeString:  timeValue2,
			shouldError: true,
		},
		{
			name:        "failure case 2",
			timeString:  timeValue3,
			shouldError: true,
		},
		{
			name:        "failure case 3",
			timeString:  timeValue4,
			shouldError: true,
		},
		{
			name:        "failure case 4",
			timeString:  timeValue5,
			shouldError: true,
		},
		{
			name:        "failure case 5",
			timeString:  timeValue6,
			shouldError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			retTime, err := ParseAPIDate(test.timeString)
			fmt.Println(retTime.String())
			if test.shouldError {
				require.Error(t, err)
			} else {
				require.Equal(t, test.timeString, strings.TrimSuffix(retTime.String(), " +0000 UTC"))
			}
		})
	}
}
