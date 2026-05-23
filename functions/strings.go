package functions

import (
	"strconv"
	"strings"
)

// TrimSequenceNumber removes a " - <number>" suffix from a string.
// main use case: getting asset / device name from nameWithSequence variable
func TrimSequenceNumber(s string) string {
	// 1. Find the last index of the separator " - "
	lastSep := strings.LastIndex(s, " - ")

	// 2. If the separator doesn't exist, return the original string.
	if lastSep == -1 {
		return s
	}

	// 3. Get the part *after* the last separator.
	suffix := s[lastSep+3:]

	// 4. Check if that suffix is purely a number.
	_, err := strconv.Atoi(suffix)

	// 5. If it is a number (no error), return the string *before* the separator.
	if err == nil {
		return s[:lastSep]
	}

	// 6. If it's not a number (e.g., "Name - Location"), return the original string.
	return s
}
