// SPDX-License-Identifier: Apache-2.0

package client

import "fmt"

// boundsCheckedAt returns items[index], or the zero value of T and an
// out-of-range error naming kind (e.g. "NameValueMap", "ObjectCollection
// insert") when index falls outside [0, len(items)).
func boundsCheckedAt[T any](items []T, index int, kind string) (T, error) {
	if index < 0 || index >= len(items) {
		var zero T
		return zero, fmt.Errorf("client: %s index %d out of range [0,%d)", kind, index, len(items))
	}
	return items[index], nil
}

// indexOfFunc returns the position of the first item matching match, or -1.
func indexOfFunc[T any](items []T, match func(T) bool) int {
	for i, v := range items {
		if match(v) {
			return i
		}
	}
	return -1
}
