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

func TestShortestPalindromeFourCharDistinct(t *testing.T) {
	actual := ShortestPalindrome("abab")
	expected := "babab"
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

func TestShortestPalindromeFiveCharPalindrome(t *testing.T) {
	actual := ShortestPalindrome("eeeee")
	expected := "eeeee"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeFiveCharOneDistinct(t *testing.T) {
	actual := ShortestPalindrome("bbaba")
	expected := "ababbaba"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeSevenCharOneDistinct(t *testing.T) {
	actual := ShortestPalindrome("babbaba")
	expected := "ababbaba"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeFiveCharPalindromeOneDistinct(t *testing.T) {
	actual := ShortestPalindrome("eeeed")
	expected := "deeeed"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeFiveCharPalindromeTwoDistinct(t *testing.T) {
	actual := ShortestPalindrome("edeed")
	expected := "deedeed"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeFiveCharPalindromeThreeDistinct(t *testing.T) {
	actual := ShortestPalindrome("edced")
	expected := "decdedced"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeFiveCharPalindromeFourDistinct(t *testing.T) {
	actual := ShortestPalindrome("edcfd")
	expected := "dfcdedcfd"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeFiveCharPalindromeFiveDistinct(t *testing.T) {
	actual := ShortestPalindrome("eacfd")
	expected := "dfcaeacfd"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeEightChars(t *testing.T) {
	actual := ShortestPalindrome("aacecaaa")
	expected := "aaacecaaa"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}

func TestShortestPalindromeMultiChars(t *testing.T) {
	actual := ShortestPalindrome("aaabddsssaabccaaabbccbddsss")
	expected := "sssddbccbbaaaccbaasssddbaaabddsssaabccaaabbccbddsss"
	if actual != expected {
		t.Errorf("expected %s but found %s", expected, actual)
	}
}
