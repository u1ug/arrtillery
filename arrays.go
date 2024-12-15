package arrtillery

import (
	"errors"
	"math/rand"
	"time"
)

// ForEach applies function to the slice
func ForEach[T any](s []T, fn func(item T) T) {
	for i, v := range s {
		s[i] = fn(v)
	}
}

// Filter returns slice where the items are filtered through the given function
func Filter[T any](s []T, fn func(item T) bool) []T {
	filtered := make([]T, len(s))
	idx := 0
	for _, v := range s {
		if fn(v) {
			filtered[idx] = v
			idx++
		}
	}
	return filtered[:idx]
}

// Some returns true if at least one slice element matches the given criteria
func Some[T any](s []T, fn func(item T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

// Every returns true if all slice elements match the given criteria
func Every[T any](s []T, fn func(item T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return false
		}
	}
	return true
}

// SafeSlice slices the input slice from start to end indices safely.
// If the end index is out of range, it uses the last index of the slice.
func SafeSlice[T any](s []T, start, end int) []T {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start > end {
		start = end
	}
	return s[start:end]
}

// RandomUnique picks N unique random elements from a slice.
// It panics if N is greater than the length of the slice.
func RandomUnique[T any](slice []T, n int) ([]T, error) {
	if n > len(slice) {
		return nil, errors.New("n cannot be greater than the length of the slice")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice[:n], nil
}

// Deduplicate removes duplicates from a slice of comparable elements.
func Deduplicate[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	var result []T
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// CountDuplicates counts the number of elements that appear more than once in the slice.
func CountDuplicates[T comparable](slice []T) int {
	counts := make(map[T]int)
	duplicates := 0
	for _, value := range slice {
		counts[value]++
	}
	for _, count := range counts {
		if count > 1 {
			duplicates++
		}
	}
	return duplicates
}

// Shuffle shuffles a slice of any type.
func Shuffle[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(slice) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
