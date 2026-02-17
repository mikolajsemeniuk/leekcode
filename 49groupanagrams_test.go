package leekcode

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

// https://leetcode.com/problems/group-anagrams
//
// 49. Group Anagrams
//
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Example 1:
//
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Example 2:
//
// Input: strs = [""]
// Output: [[""]]
//
// Example 3:
//
// Input: strs = ["a"]
// Output: [["a"]]
//
// Constraints:
//
// 1 <= strs.length <= 10^4
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.
func groupAnagrams(strs []string) [][]string {
	var mod int = 1e9 + 7

	fn := func(str string) int {
		if len(str) == 0 {
			return 0
		}

		sum := 0
		product := 1
		for i := 0; i < len(str); i++ {
			sum += int(str[i])
			product = (product * int(str[i])) % mod
		}

		return (sum * product) % mod
	}

	m := map[int][]string{}
	for _, elem := range strs {
		key := fn(elem)
		m[key] = append(m[key], elem)
	}

	var res [][]string
	for _, value := range m {
		res = append(res, value)
	}

	return res
}

func groupAnagramsSlower(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	anagrams := map[string][]string{}

	for _, v := range strs {
		runes := []rune(v)
		slices.Sort(runes)
		key := string(runes)
		anagrams[key] = append(anagrams[key], v)
	}

	out := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		out = append(out, v)
	}

	return out
}

// Helper function to check if two slices contain the same groups (order-independent)
func sameGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort each group and convert to comparable format
	sortGroups := func(groups [][]string) []string {
		sorted := make([]string, len(groups))
		for i, group := range groups {
			sortedGroup := make([]string, len(group))
			copy(sortedGroup, group)
			sort.Strings(sortedGroup)
			// Join to create a unique string representation
			result := ""
			for _, s := range sortedGroup {
				result += s + ","
			}
			sorted[i] = result
		}
		sort.Strings(sorted)
		return sorted
	}

	sortedA := sortGroups(a)
	sortedB := sortGroups(b)

	return reflect.DeepEqual(sortedA, sortedB)
}

func TestGroupAnagrams(t *testing.T) {
	t.Run("example 1: [eat,tea,tan,ate,nat,bat]", func(t *testing.T) {
		strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		result := groupAnagrams(strs)

		// Expected groups (order doesn't matter)
		expected := [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("example 2: empty string", func(t *testing.T) {
		strs := []string{""}
		result := groupAnagrams(strs)
		expected := [][]string{{""}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("example 3: single character", func(t *testing.T) {
		strs := []string{"a"}
		result := groupAnagrams(strs)
		expected := [][]string{{"a"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("no anagrams", func(t *testing.T) {
		strs := []string{"abc", "def", "ghi"}
		result := groupAnagrams(strs)
		expected := [][]string{{"abc"}, {"def"}, {"ghi"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("all same anagrams", func(t *testing.T) {
		strs := []string{"abc", "bca", "cab", "acb"}
		result := groupAnagrams(strs)
		expected := [][]string{{"abc", "bca", "cab", "acb"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("two groups of anagrams", func(t *testing.T) {
		strs := []string{"abc", "bca", "xyz", "zyx"}
		result := groupAnagrams(strs)
		expected := [][]string{{"abc", "bca"}, {"xyz", "zyx"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("single letter strings", func(t *testing.T) {
		strs := []string{"a", "b", "a", "c", "b"}
		result := groupAnagrams(strs)
		expected := [][]string{{"a", "a"}, {"b", "b"}, {"c"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("long strings", func(t *testing.T) {
		strs := []string{"listen", "silent", "enlist", "hello", "world"}
		result := groupAnagrams(strs)
		expected := [][]string{{"listen", "silent", "enlist"}, {"hello"}, {"world"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("strings with repeated letters", func(t *testing.T) {
		strs := []string{"aab", "aba", "baa", "aaa"}
		result := groupAnagrams(strs)
		expected := [][]string{{"aab", "aba", "baa"}, {"aaa"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("mixed length strings", func(t *testing.T) {
		strs := []string{"a", "aa", "aaa", "a"}
		result := groupAnagrams(strs)
		expected := [][]string{{"a", "a"}, {"aa"}, {"aaa"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("palindrome anagrams", func(t *testing.T) {
		strs := []string{"racecar", "carrace", "civic", "vicic"}
		result := groupAnagrams(strs)
		expected := [][]string{{"racecar", "carrace"}, {"civic", "vicic"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})

	t.Run("empty and non-empty mix", func(t *testing.T) {
		strs := []string{"", "a", "", "b"}
		result := groupAnagrams(strs)
		expected := [][]string{{"", ""}, {"a"}, {"b"}}

		if !sameGroups(result, expected) {
			t.Errorf("Expected groups %v, got %v", expected, result)
		}
	})
}
