package nine

import "strings"

// StringRotation checks if strings are rotations of each other.
// Concatenates a string to itself to check for rotation which doubles space.
func StringRotation(s1, s2 string) bool {
	t := s2 + s2
	return IsSubstring(t, s1)
}

func IsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}
