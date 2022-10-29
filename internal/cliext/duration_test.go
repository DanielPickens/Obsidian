package cliext

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDurationParse(t *testing.T) {
	tests := []struct {
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

func TestDurationMarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		val       time.Duration
		expected  string
	}{
		{name: "seconds", val: 15 * time.Second, expected: "15s"},
		{name: "minutes", val: 5 * time.Minute, expected: "5m"},
		{name: "mixed", val: 70 * time.Second, expected: "1m10s"},
	}

	if testing.Short() {
		tests = tests[:1]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.val.MarshalJSON()
			assert.Nil(t, err)
			assert.Equal(t, []byte(tt.expected), res)
		}
		)
	}
}

func TestDurationBool(t *testing.T) {
	tests := []struct {
		name      string
		val       time.Duration
		expected  bool
	}{
		{name: "zero", val: 0, expected: false},
		{name: "non-zero", val: 15 * time.Second, expected: true},
	}

	if testing.Short() {
		tests = tests[:1]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := tt.val.Bool()

			assert.Equal(t, tt.expected, res)
		}
		)
	}
}

func TestDurationUnmarshalJSON(t *testing.T) {
	tests := []struct {
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
			var res time.Duration

			err := res.UnmarshalJSON([]byte(tt.val))

			if tt.expectErr {

				assert.NotNil(t, err)
			} else {

				assert.Nil(t, err)

				assert.Equal(t, tt.expected, res)

			}

		}

		)
	}
}




func TestDurationParseError(t *testing.T) {
	_, err := ParseDuration("")
	assert.Error(t, err, "Expected non nil error")
}

func TestDurationPrimative(t *testing.T) {
	tests := []struct {
		name      string
		val       time.Duration
		expected  bool
	}{
		{name: "zero", val: 0, expected: false},
		{name: "non-zero", val: 15 * time.Second, expected: true},
	}

	if testing.Short() {
		tests = tests[:1]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := tt.val.Primative()

			assert.Equal(t, tt.expected, res)
		}
		)
	}
}

func TestDurationNonPrimativeDataType(t *testing.T) {
	tests := []struct {
		name      string
		val       time.Duration
		expected  bool
	}{
		{name: "zero", val: 0, expected: false},
		{name: "non-zero", val: 15 * time.Second, expected: true},
	}

	if testing.Short() {
		tests = tests[:1]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := tt.val.NonPrimativeDataType()

			assert.Equal(t, tt.expected, res)
		}
		)
	}
}

