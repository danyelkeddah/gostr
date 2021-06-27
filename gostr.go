package gostr

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// Return the remainder of a string after the first occurrence of a given value.
func After(s string, search string) string {

	if search == "" {
		return s
	}

	position := strings.Index(s, search)

	if position == -1 {
		return s
	}
	return s[position+len(search):]
}

// Return the remainder of a string after the last occurrence of a given value.
func AfterLast(s string, search string) string {
	if search == "" {
		return s
	}

	position := strings.LastIndex(s, search)

	if position == -1 {
		return s
	}

	return s[position+len(search):]
}

// Get the portion of a string before the first occurrence of a given value.
func Before(s string, search string) string {
	if search == "" {
		return s
	}
	position := strings.Index(s, search)

	if position == -1 {
		return s
	}
	return s[0:position]
}

// Get the portion of a string before the last occurrence of a given value.
func BeforeLast(s string, search string) string {
	if search == "" {
		return s
	}

	str := strings.LastIndex(s, search)
	if str == -1 {
		return s
	}

	return s[0:str]
}

// Get the portion of a string between two given values.
func Between(s string, from string, to string) string {
	if from == "" || to == "" {
		return s
	}

	return BeforeLast(After(s, from), to)
}

// Determine if a given string contains any of the the given substrings.
func Contains(haystack string, needles []string) bool {
	for _, needle := range needles {
		if needle != "" && strings.Index(haystack, needle) != -1 {
			return true
		}
	}

	return false
}

// Determine if a given string contains all given substrings.
func ContainsAll(haystack string, needles []string) bool {
	for _, needle := range needles {
		if !Contains(haystack, []string{needle}) {
			return false
		}
	}

	return true
}

// Determine if a given string ends with any of the the given substrings.
func EndsWith(haystack string, needles []string) bool {
	for _, needle := range needles {
		if needle != "" &&
			strings.HasSuffix(haystack, needle) {
			return true
		}
	}

	return false
}

// Determine if a given string is a valid UUID.
func IsUUID(s string) bool {
	_, err := uuid.Parse(s)

	return err == nil
}

// Return the length of the given string.
func Length(s string) int {
	return len(s)
}

// Convert the given string to lower-case.
func Lower(s string) string {
	return strings.ToLower(s)
}

// Limit the number of words in a string.
func Words(s string, words uint, end string) string {
	if words <= 1 || s == "" {
		return s
	}

	exp := fmt.Sprintf(`^\s*(?:\S+\s*){1,%d}`, words)
	re := regexp.MustCompile(exp)
	t := re.FindString(s)
	return fmt.Sprintf(`%v%v`, strings.TrimSpace(t), end)
}

// Pad the right side of a string with another.
func PadRight(s string, length uint, pad string) string {
	if len(s) >= int(length) || len(pad) == 0 {
		return s
	}

	if (int(length)-len(s))%len(pad) == 0 {
		nPads := (int(length) - len(s)) / len(pad)

		return s + strings.Repeat(pad, nPads)
	}
	nPads := (int(length) - len(s)) / len(pad)
	rC := int(length) - len(s+strings.Repeat(pad, nPads))

	return s + strings.Repeat(pad, nPads) + pad[:rC]
}

// Pad the left side of a string with another.
func PadLeft(s string, length uint, pad string) string {
	if len(s) >= int(length) || len(pad) == 0 {
		return s
	}

	if (int(length)-len(s))%len(pad) == 0 {
		nPads := (int(length) - len(s)) / len(pad)

		return strings.Repeat(pad, nPads) + s
	}
	nPads := (int(length) - len(s)) / len(pad)
	rC := int(length) - len(s+strings.Repeat(pad, nPads))

	return strings.Repeat(pad, nPads) + pad[:rC] + s
}

// Generate a more truly "random" alpha-numeric string.
func Random(l uint) string {
	s := ""
	if l == 0 {
		l = 16
	}
	for len(s) < int(l) {
		size := int(l) - len(s)
		bytes := make([]byte, size)
		rand.Read(bytes)
		encoded := base64.RawStdEncoding.EncodeToString(bytes)
		r := strings.NewReplacer("/", "", "+", "", "=", "")
		s += r.Replace(encoded[0:size])
	}
	return s
}

// Repeat the given string.
func Repeat(s string, times uint) string {
	return strings.Repeat(s, int(times))
}

// Convert the given string to upper-case.
func Upper(s string) string {
	return strings.ToUpper(s)
}
