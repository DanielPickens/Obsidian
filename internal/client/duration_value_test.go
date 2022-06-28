package cliext

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDurationValue(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{name: "no_transform", text: "1s", expected: "1s"},
		{name: "mixed_case", text: "1s", expected: "1s"},
		{name: "lower_case", text: "1s", expected: "1s"},
		{name: "multiple_occurances", text: "1s1s", expected: "1s1s"},

}tests := []struct {
		name      string
		val       string
		expected  time.Duration
		expectErr bool
	}{
		{name: "parse empty", val: "", expectErr: true},
		{name: "parse invalid", val: "oops", expectErr: true},
		{name: "parse invalid format", val: "1hour", expectErr: true},
		{name: "parse seconds", val: "15s", expected: 15 * time.Second},
		{name: "parse default seconds", val: "5", expected: 5 * time.Second},
		{name: "parse minutes", val: "5m", expected: 5 * time.Minute},
		{name: "parse mixed", val: "1m10s", expected: 70 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ParseDuration(tt.val)

			if tt.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}


func TestDurationString(t *testing.T) {
	tests := []struct {
		name      string
		val       time.Duration
		expected  string
	}{
		{name: "seconds", val: 15 * time.Second, expected: "15s"},
		{name: "minutes", val: 5 * time.Minute, expected: "5m"},
		{name: "mixed", val: 70 * time.Second, expected: "1m10s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.val.String()
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestDurationValue_Parse(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected time.Duration
	}{
		{name: "no_transform", text: "1s", expected: 1 * time.Second},
		{name: "mixed_case", text: "1s", expected: 1 * time.Second},
		{name: "lower_case", text: "1s", expected: 1 * time.Second},
		{name: "multiple_occurances", text: "1s1s", expected: 1 * time.Second + 1 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ParseDuration(tt.text)
			require.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}
