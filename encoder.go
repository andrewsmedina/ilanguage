package encoder

import "strings"

// Encode translates the data to 'i' language
func Encode(data string) string {
	data = strings.Replace(data, "a", "i", -1)
	data = strings.Replace(data, "e", "i", -1)
	data = strings.Replace(data, "o", "i", -1)
	data = strings.Replace(data, "u", "i", -1)
	return data
}
