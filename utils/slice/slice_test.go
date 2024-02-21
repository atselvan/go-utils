package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryExists(t *testing.T) {

	slice := []string{"one", "two", "three"}

	t.Run("entry exists", func(t *testing.T) {
		assert.True(t, EntryExists(slice, "one"))
	})

	t.Run("entry does not exists", func(t *testing.T) {
		assert.False(t, EntryExists(slice, "four"))
	})
}

func ExampleEntryExists() {
	slice := []string{"one", "two", "three"}
	fmt.Println(EntryExists(slice, "two"))
	// Output: true
}

func TestGetSliceEntryIndex(t *testing.T) {

	slice := []string{"one", "two", "three"}

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, 0, GetSliceEntryIndex(slice, "one"))
	})

	t.Run("entry does not exist", func(t *testing.T) {
		assert.Equal(t, -1, GetSliceEntryIndex(slice, "four"))
	})
}

func ExampleGetSliceEntryIndex() {
	slice := []string{"one", "two", "three"}
	fmt.Println(GetSliceEntryIndex(slice, "one"))
	// Output: 0
}

func TestRemoveEntryFromSlice(t *testing.T) {

	t.Run("remove entry that exists", func(t *testing.T) {
		slice := []string{"one", "two", "three"}
		assert.Equal(t, []string{"one", "three"}, RemoveEntryFromSlice(slice, "two"))
	})

	t.Run("remove entry that does not exist", func(t *testing.T) {
		slice := []string{"one", "two", "three"}
		assert.Equal(t, slice, RemoveEntryFromSlice(slice, "four"))
	})

	t.Run("remove entry that exists more than once", func(t *testing.T) {
		slice := []string{"one", "two", "three", "two"}
		assert.Equal(t, []string{"one", "three", "two"}, RemoveEntryFromSlice(slice, "two"))
	})
}

func ExampleRemoveEntryFromSlice() {
	slice := []string{"one", "two", "three"}
	fmt.Println(RemoveEntryFromSlice(slice, "two"))
	// Output: [one three]
}

func TestRemoveDuplicateEntries(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}
	expectedResult := []string{"one", "two", "three", "four"}

	assert.Equal(t, expectedResult, RemoveDuplicateEntries(slice))
}

func ExampleRemoveDuplicateEntries() {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}
	fmt.Println(RemoveDuplicateEntries(slice))
	// output: [one two three four]
}

func TestCountDuplicateEntries(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}

	expectedResult := make(map[string]int)
	expectedResult["one"] = 1
	expectedResult["two"] = 2
	expectedResult["three"] = 3
	expectedResult["four"] = 4

	result := CountDuplicateEntries(slice)

	assert.Equal(t, expectedResult, result)
}

func ExampleCountDuplicateEntries() {
	slice := []string{"one", "two", "two", "three", "three", "three"}
	fmt.Println(CountDuplicateEntries(slice))
	// Output: map[one:1 three:3 two:2]
}

func TestDuplicateEntryExists(t *testing.T) {

	t.Run("duplicate entries exists", func(t *testing.T) {
		slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}
		assert.True(t, DuplicateEntryExists(slice))
	})

	t.Run("duplicate entries do not exists", func(t *testing.T) {
		slice := []string{"one", "two", "three", "four"}
		assert.False(t, DuplicateEntryExists(slice))
	})
}

func ExampleDuplicateEntryExists() {
	slice := []string{"one", "two", "two"}
	fmt.Println(DuplicateEntryExists(slice))
	// Output: true
}

func TestCompareStringSlice(t *testing.T) {
	slice1 := []string{"one", "two", "three"}
	slice2 := []string{"two", "one", "three"}
	slice3 := []string{"one", "two", "two"}

	t.Run("slices match", func(t *testing.T) {
		assert.True(t, CompareStringSlice(slice1, slice2))
	})

	t.Run("slices do not match", func(t *testing.T) {
		assert.False(t, CompareStringSlice(slice1, slice3))
	})
}

func ExampleCompareStringSlice() {
	slice1 := []string{"one", "two", "three"}
	slice2 := []string{"two", "one", "three"}

	fmt.Println(CompareStringSlice(slice1, slice2))
	// Output: true
}
