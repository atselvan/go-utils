package slice

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// EntryExists checks if an entry exists in a slice of strings
// The function returns a boolean value:
//
//	true if the entry exists or
//	false if the entry does not exist
func EntryExists(slice []string, entry string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == entry {
			return true
		}
	}
	return false
}

// GetSliceEntryIndex returns the index of an entry in a slice of strings
// The function returns an integer value of the index of the first occurrence of the slice entry
func GetSliceEntryIndex(slice []string, entry string) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == entry {
			return i
		}
	}
	return -1
}

// RemoveEntryFromSlice removes a entry from a slice of strings
// The function removed the first occurrence of the entry and then returns the updated slice back
// If the Entry does not exist then the function returns the same slice back
func RemoveEntryFromSlice(slice []string, entry string) []string {
	i := GetSliceEntryIndex(slice, entry)
	if i == -1 {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

// RemoveDuplicateEntries removes duplicate entries in a slice of strings
// The function returns a slice with unique string entries
func RemoveDuplicateEntries(stringSlice []string) []string {
	keys := make(map[string]bool)
	var result []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

// CountDuplicateEntries counts the number of times a entry repeats in a slice of strings
// The function returns a map with the unique entries in the slice as keys and
// the duplicate frequencies of the unique entries of the slice
func CountDuplicateEntries(list []string) map[string]int {

	duplicateCount := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicateCount map
		if _, exist := duplicateCount[item]; exist {
			duplicateCount[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicateCount[item] = 1 // else start counting from 1
		}
	}
	return duplicateCount
}

// DuplicateEntryExists checks if a slice has duplicate entries or not
// The function returns a boolean response
//
//	true : the slice contains duplicate entries
//	func : the slice does not contain duplicate entries
func DuplicateEntryExists(stringSlice []string) bool {

	duplicateCount := CountDuplicateEntries(stringSlice)

	for _, k := range duplicateCount {
		if k > 1 {
			return true
		}
	}

	return false
}

// CompareStringSlice compares two slices of strings.
// The function returns a boolean response
//
//	true : the slices match
//	func : the slices do not match
func CompareStringSlice(s1, s2 []string) bool {
	less := func(a, b string) bool { return a < b }
	return cmp.Diff(s1, s2, cmpopts.SortSlices(less)) == ""
}
