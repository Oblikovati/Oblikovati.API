// SPDX-License-Identifier: Apache-2.0
package contract

// AssemblyLoading exposes whether an assembly uses a saved express presentation.
// Example: `if assembly.ExpressMode() { err = assembly.LoadFull() }`.
type AssemblyLoading interface {
	// ExpressMode reports that component geometry has not been fully loaded.
	ExpressMode() bool
	// LoadFull loads the components; on error the express state remains active.
	LoadFull() error
}
