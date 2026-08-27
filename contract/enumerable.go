// SPDX-License-Identifier: Apache-2.0

package contract

// Enumerable is a 0-based, Count()+Item(i)-style read-only collection view — the
// shared shape behind every scalar collection contract in this package (ADR-0006,
// Oblikovati GPL repo: IEnumerable/1-based Item[i] → a 0-based generic collection).
// Item returns the zero value of T for an out-of-range index rather than panicking.
type Enumerable[T any] interface {
	// Count returns the number of elements in the collection.
	Count() int
	// Item returns the element at index i (0-based), or the zero value when out of range.
	Item(i int) T
}
