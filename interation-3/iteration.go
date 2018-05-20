// Package iteration terates over stuff.
package iteration

import "strings"

// Repeat repeats a character n times.
func Repeat(character string, n int) string {
	return strings.Repeat(character, n)
}
