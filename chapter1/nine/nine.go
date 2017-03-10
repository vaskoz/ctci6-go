package nine

import "strings"

func StringRotation(s1, s2 string) bool {
	t := s2 + s2
	return IsSubstring(t, s1)
}

func IsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}
