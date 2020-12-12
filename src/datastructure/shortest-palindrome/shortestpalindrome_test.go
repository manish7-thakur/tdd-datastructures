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
func TestShortestPalindromeTwoCharSame(t *testing.T) {
	actual := ShortestPalindrome("aa")
	expected := "aa"
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

func TestShortestPalindromeThreeRepeatedCharOneDistinct(t *testing.T) {
	actual := ShortestPalindrome("bbd")
	expected := "dbbd"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeThreeCharPalindrome(t *testing.T) {
	actual := ShortestPalindrome("bbb")
	expected := "bbb"
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

func TestShortestPalindromeFourCharPalindrome(t *testing.T) {
	actual := ShortestPalindrome("cccc")
	expected := "cccc"
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

func TestShortestPalindromeFourRepeatedCharOneDistinct(t *testing.T) {
	actual := ShortestPalindrome("bbbd")
	expected := "dbbbd"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func ShortestPalindrome(s string) string {
	strlen := len(s)
	if strlen <= 1 {
		return s
	}
	mid := strlen / 2
	i := 0
	j := 0
	if strlen%2 == 0 {
		i = mid - 1
		j = mid
	} else {
		i = mid
		j = mid
	}

	for ; i >= 0 && j < strlen; {
		if s[i] == s[j] {
			i--
			j++
		} else {
			for ; i != j && s[i] != s[j]; {
				j--
			}
		}
	}
	prefix := ""
	for ; j < strlen; j++ {
		prefix = string(s[j]) + prefix
	}
	return prefix + s
}
