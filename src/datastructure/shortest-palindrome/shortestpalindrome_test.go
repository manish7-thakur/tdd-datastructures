package shortest_palindrome

import (
	"math"
	"testing"
)

func TestShortestPalindromeEmptyStr(t *testing.T) {
	actual := ShortestPalindrome("")
	expected := ""
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeSingleChar(t *testing.T) {
	actual := ShortestPalindrome("a")
	expected := "a"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeTwoChar(t *testing.T) {
	actual := ShortestPalindrome("ab")
	expected := "bab"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeThreeChar(t *testing.T) {
	actual := ShortestPalindrome("abb")
	expected := "bbabb"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeThreeCharDiff(t *testing.T) {
	actual := ShortestPalindrome("abc")
	expected := "cbabc"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeFourChar(t *testing.T) {
	actual := ShortestPalindrome("abcd")
	expected := "dcbabcd"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func ShortestPalindrome(s string) string {
	strlen := len(s)
	mid := int(math.Ceil(float64(strlen / 2.0)))
	i := mid
	j := mid
	for ; i > 0 && j < strlen; {
		if s[i-1] != s[j] {
			i--
		} else {

		}
	}
	prefix := ""
	i++
	for ; i < strlen; i++ {
		prefix = string(s[i]) + prefix
	}
	return prefix + s
}
