package schema

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/mail"
	"net/url"
	"regexp"
)

// Valid string formats and configuration.
const (
	stringURI         = "uri"
	stringEmail       = "email"
	stringUUID        = "uuid"
	stringBinary      = "binary"
	stringUUIDVersion = 4
)

func checkStringConstraints(v, pattern string, min, max int, t string) (string, error) {
	fmt.Println(pattern)
	fmt.Println(min, len(v))
	fmt.Println(max)
	if min != 0 && len(v) <= min {
		return v, fmt.Errorf("constraint check error: %s:%v < minimum:%v", t, v, min)
	}

	if max != 0 && len(v) >= max {
		return v, fmt.Errorf("constraint check error: %s:%v > maximum:%v", t, v, max)
	}

	if pattern != "" {
		match, err := regexp.MatchString(pattern, v)
		if false == match || err != nil {
			return v, fmt.Errorf("constraint check error: %s:%v don't fit pattern : %v ", t, v, pattern)
		}
	}

	return v, nil
}

func decodeString(format, value string, c Constraints) (string, error) {

	_, err := checkStringConstraints(value, c.Pattern, c.MinLength, c.MaxLength, StringType)

	if err != nil {
		return value, err
	}

	switch format {
	case stringURI:
		_, err := url.ParseRequestURI(value)
		return value, err
	case stringEmail:
		_, err := mail.ParseAddress(value)
		return value, err
	case stringUUID:
		v, err := uuid.FromString(value)
		if v.Version() != stringUUIDVersion {
			return value, fmt.Errorf("invalid UUID version - got:%d want:%d", v.Version(), stringUUIDVersion)
		}
		return value, err
	}
	// NOTE: Returning the value for unknown format is in par with the python library.

	return value, nil

}
func castString(format, value string) (string, error) {
	switch format {
	case stringURI:
		_, err := url.ParseRequestURI(value)
		return value, err
	case stringEmail:
		_, err := mail.ParseAddress(value)
		return value, err
	case stringUUID:
		v, err := uuid.FromString(value)
		if v.Version() != stringUUIDVersion {
			return value, fmt.Errorf("invalid UUID version - got:%d want:%d", v.Version(), stringUUIDVersion)
		}
		return value, err
	}
	// NOTE: Returning the value for unknown format is in par with the python library.
	return value, nil
}
