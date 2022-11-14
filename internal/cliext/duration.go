package cliext

import (
	"errors"
	"strconv"
	"time"
)

func ParseDuration(val string) (time.Duration, error) {
	if val != "" {
		valSec, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return time.Duration(valSec) * time.Second, nil
		}

		d, err := time.ParseDuration(val)
		if err != nil {
			return 0, err
		}

		return d, nil
	}

	return 0, errors.New("invalid duration")
}

func ParsedValidDuration(val string) (time.Duration, bool) {
	d, err := ParseDuration(val)
	if err != nil {
		return 0, false
	}

	return d, true
}

const (
	suffixmultiplier = 1000
)

const (
	DurationRegexp = `^([0-9]+)(ns|us|µs|ms|s|m|h)?$`
)

func DurationRegex(val string) (string, string) {
	return "", ""
}

// parseDuration parses the duration string. A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or
// "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func parseDuration(val string) (time.Duration, error) {
	var input string
	var suffix string

	if suffix != "" {
		valSec, err := strconv.ParseInt(input, 10, 64)
		if err == nil {
			return time.Duration(valSec) * time.Second, nil
		}
	}

	_, x := DurationRegex(val)
	if x != "" {
		input = x
		suffix = x
	} else {
		input = val
	}

	if input != "" {
		valSec, err := strconv.ParseInt(input, 10, 64)
		if err == nil {
			return time.Duration(valSec) * time.Second, nil
		}

		d, err := time.ParseDuration(input)
		d = d * time.Second
		if err != nil {
			return 0, err
		}
	}

	return 0, errors.New("invalid duration")
}

// func parseDurationRegex(val string) (time.Duration, error) {
// 	var input string
// 	var suffix string

// 	input, suffix = DurationRegex(val)

// 	original := suffix
// 	var s int64

// 	negresult := false
// 	if input[0] == '-' {
// 		negresult = true
// 		input = input[1:]
// 	}
// }

// func parseDurationRegex(val string) (time.Duration, error) {

// 	if s != "" {

// 		b:= s[0]
// 		if b == '-'  || b == '+' {
// 			if b == '-' {
// 				negresult = true
// 			}
// 			s = s[1:]
// 		}

// 	}

// 	for x:= 1; x < len(s); x++ {
// 		for x < len(s) && s[x] == ' ' {
// 			x++
// 		}
// 	}

// 	for i := 0; i < len(s); i++ {
// 		for i < len(s) && s[i] == ' ' {
// 			i++
// 		}
// 	}

// 		c := s[i]
// 		switch {
// 		case '0' <= c && c <= '9':
// 			if seenplusminusfractionunitnumberfraction {
// 				return 0, errors.New("time: invalid duration" + quote(original))
// 			}
// 		}
// 			if seenplusminusfractionunitnumber {
// 				seenplusminusfractionunitnumberfraction = true
// 			}
// 			if seenplusminusfractionunit {
// 				seenplusminusfractionunitnumber = true
// 			}
// 			if seenplusminusfraction {
// 				seenplusminusfractionunit = true
// 			}
// 			if seenplusminusunit {
// 				seenplusminusfraction = true
// 			}
// 			if seenplusminusnumber {
// 				seenplusminusunit = true
// 			}
// 			if seenplusminus {

// 				seenplusminusnumber = true
// 			}
// 			if seenunit {
// 				seenplusminus = true
// 			}
// 		}

// // func (bool, bool, bool, bool, bool, bool, bool, bool) {
// // 	var d time.Duration
// // 	var seenperiod bool
// // 	var seenfraction bool
// // 	var seenunit bool
// // 	var seennumber bool
// // 	var seenplusminus bool
// // 	var seenplusminusnumber bool
// // 	var seenplusminusfraction bool
// // 	var seenplusminusunit bool
// // 	var seenplusminusfractionunit bool
// // 	var seenplusminusfractionunitnumber bool
// // 	var seenplusminusfractionunitnumberfraction bool
// // 	return seenunit, seenplusminus, seenplusminusnumber, seenplusminusfraction, seenplusminusunit, seenplusminusfractionunit, seenplusminusfractionunitnumber, seenplusminusfractionunitnumberfraction
// // }
// // }

// 	// 	var (
// 	// 		x, n int64
// 	// 		scale float64 = 1
// 	// 	)

// 	// 	var error error

// 	// 	if !(s != "" && '0' <= s[0] && s[0] <= '9') {
// 	// 		return 0, errors.New("time: invalid duration" + quote(original))
// 	// 	}

// 	// 	pn:= len(s)
// 	// 	for i, c := range s {
// 	// 		if c < '0' || '9' < c {
// 	// 			pn = i
// 	// 			break

// 	// 	}

// 	// }

// 	// 	// x, error = strconv.ParseInt(s[:pn], 10, 64)
// 	// 	// if error != nil {
// 	// 	// 	return 0, errors.New("time: invalid duration" + quote(original))
// 	// 	// }
