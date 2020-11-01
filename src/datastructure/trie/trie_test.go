package trie

import (
	"testing"
)

func TestInsertEmptyString(t *testing.T) {
	trie := Trie{}
	str := ""
	trie.Insert(str)
	actual, ok := trie.first.chars['\000']
	if ok {
		t.Errorf("expected empty but found %v", actual)
	}
}

func TestInsertSingleCharString(t *testing.T) {
	trie := Trie{}
	str := "a"
	trie.Insert(str)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
}

func TestInsertDoubleCharString(t *testing.T) {
	trie := Trie{}
	str := "ab"
	trie.Insert(str)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != false {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['b']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
}

func TestInsertTripleCharString(t *testing.T) {
	trie := Trie{}
	str := "abc"
	trie.Insert(str)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != false {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['b']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != false {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['c']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
}

func TestInsertMultipleStrings(t *testing.T) {
	trie := Trie{}
	str1 := "a"
	str2 := "ab"
	str3 := "abc"
	trie.Insert(str1)
	trie.Insert(str2)
	trie.Insert(str3)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['b']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['c']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected false  but found %t", actual.endOfWord)
	}
}

func TestInsertMultipleStringsPrefix(t *testing.T) {
	trie := Trie{}
	str1 := "ab"
	str2 := "a"
	trie.Insert(str1)
	trie.Insert(str2)
	actual, ok := trie.first.chars['a']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
	actual, ok = actual.chars['b']
	if !ok {
		t.Errorf("expected true  but found %t", ok)
	}
	if actual.endOfWord != true {
		t.Errorf("expected true  but found %t", actual.endOfWord)
	}
}

func TestSearchEmptyString(t *testing.T) {
	str := ""
	trie := Trie{}
	actual := trie.Contains(str, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharString(t *testing.T) {
	str := "a"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains(str, false)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharString(t *testing.T) {
	str := "ab"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains(str, false)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharNonExistentString(t *testing.T) {
	trie := Trie{}
	actual := trie.Contains("aa", false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchTripleCharString(t *testing.T) {
	str := "aaaa"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains(str, false)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchMultiplePrefixedStrings(t *testing.T) {
	str1 := "a"
	trie := Trie{}
	trie.Insert(str1)
	actual := trie.Contains("aa", false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchMultipleNonExistentStrings(t *testing.T) {
	str1 := "a"
	str2 := "aaa"
	trie := Trie{}
	trie.Insert(str1)
	trie.Insert(str2)
	actual := trie.Contains(str1, false)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains(str2, false)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains("aa", false)
	expected = false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains("aaaa", false)
	expected = false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchEmptyPrefix(t *testing.T) {
	str := ""
	trie := Trie{}
	actual := trie.Contains(str, true)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharPrefix(t *testing.T) {
	str := "ab"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains("a", true)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchSingleCharPrefixNonExistent(t *testing.T) {
	str := "ab"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains("b", true)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharPrefix(t *testing.T) {
	str := "abcd"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains("ab", true)
	expected := true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestSearchDoubleCharPrefixNonExistent(t *testing.T) {
	str := "abcd"
	trie := Trie{}
	trie.Insert(str)
	actual := trie.Contains("ac", true)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestDeleteEmptyString(t *testing.T) {
	str := ""
	trie := Trie{}
	trie.Insert(str)
	trie.Delete(str)
	actual := trie.Contains(str, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
}

func TestDeleteSingleCharString(t *testing.T) {
	str := "a"
	trie := Trie{}
	trie.Insert(str)
	trie.Delete(str)
	actual := trie.Contains(str, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	_, ok := trie.first.chars['a']
	if ok {
		t.Errorf("expected empty but found %t", ok)
	}
}

func TestDeleteDoubleCharString(t *testing.T) {
	str := "ab"
	trie := Trie{}
	trie.Insert(str)
	trie.Delete(str)
	actual := trie.Contains(str, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	_, ok := trie.first.chars['a']
	if ok {
		t.Errorf("expected empty but found %t", ok)
	}
}

func TestDeleteTripleCharString(t *testing.T) {
	str := "abc"
	trie := Trie{}
	trie.Insert(str)
	trie.Delete(str)
	actual := trie.Contains(str, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	_, ok := trie.first.chars['a']
	if ok {
		t.Errorf("expected empty but found %t", ok)
	}
}

func TestDeleteMultiplePrefixedStrings(t *testing.T) {
	str1 := "ab"
	str2 := "a"
	trie := Trie{}
	trie.Insert(str1)
	trie.Insert(str2)
	trie.Delete(str1)
	actual := trie.Contains(str1, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains(str2, false)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	trie.Delete(str2)
	_, ok := trie.first.chars['a']
	expected = false
	if ok {
		t.Errorf("expected %t but found %t", expected, ok)
	}
}

func TestDeleteMultiplePrefixedStringsReversed(t *testing.T) {
	str1 := "ab"
	str2 := "a"
	trie := Trie{}
	trie.Insert(str1)
	trie.Insert(str2)
	trie.Delete(str2)
	actual := trie.Contains(str2, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains(str1, false)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	trie.Delete(str1)
	_, ok := trie.first.chars['a']
	expected = false
	if ok {
		t.Errorf("expected %t but found %t", expected, ok)
	}
}

func TestDeleteMultiplePrefixedThreeCharStrings(t *testing.T) {
	str1 := "abc"
	str2 := "ab"
	str3 := "a"
	trie := Trie{}
	trie.Insert(str1)
	trie.Insert(str2)
	trie.Insert(str3)
	trie.Delete(str1)
	actual := trie.Contains(str1, false)
	expected := false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains(str2, false)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains(str3, false)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	trie.Delete(str2)
	actual = trie.Contains(str2, false)
	expected = false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	actual = trie.Contains(str3, false)
	expected = true
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	trie.Delete(str3)
	actual = trie.Contains(str3, false)
	expected = false
	if actual != expected {
		t.Errorf("expected %t but found %t", expected, actual)
	}
	_, ok := trie.first.chars['a']
	expected = false
	if ok {
		t.Errorf("expected %t but found %t", expected, ok)
	}
}
