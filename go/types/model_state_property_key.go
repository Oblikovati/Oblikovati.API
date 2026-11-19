// SPDX-License-Identifier: Apache-2.0

package types

// ModelStatePropertyKey names a document property without depending on a display label.
// Example: `types.ModelStatePropertyKey{Set: "Summary Information", Name: "Category"}`.
type ModelStatePropertyKey struct {
	Set  string `json:"set"`
	Name string `json:"name"`
}
