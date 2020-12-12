package shortest_palindrome

import (
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

func TestShortestPalindromeFourRepeatedChar(t *testing.T) {
	actual := ShortestPalindrome("babd")
	expected := "dbabd"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func ShortestPalindrome(s string) string {
	strlen := len(s)
	mid := strlen / 2
	i := 0
	if mid-1 >= 0 {
		i = mid - 1
	} else {
		i = mid
	}
	j := i + 1
	for ; i > 0 && j < strlen; {
		if s[i-1] != s[j] {
			i--
			j--
		} else {
			i--
			j++
		}
	}
	prefix := ""
	for ; j < strlen; j++ {
		prefix = string(s[j]) + prefix
	}
	return prefix + s
}
