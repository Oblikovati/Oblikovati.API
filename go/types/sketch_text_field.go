// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

// SketchTextFieldKind identifies the source of a live Sketch Text value.
type SketchTextFieldKind string

const (
	SketchTextFieldModelParameter SketchTextFieldKind = "modelParameter"
	SketchTextFieldFileName       SketchTextFieldKind = "fileName"
)

// SketchTextFieldPlaceholder occupies one Unicode scalar in stored content.
const SketchTextFieldPlaceholder = "\uFFFC"

// SketchTextField binds the placeholder at At, a zero-based Unicode scalar offset.
// For example, At: 1 links the placeholder in "⌀\uFFFC". Formatting runs may style
// the placeholder and apply to the entire displayed value. Fractions cannot contain
// a linked placeholder. Only placeholders named by a field resolve to a value.
type SketchTextField struct {
	At        uint32              `json:"at"`
	Kind      SketchTextFieldKind `json:"kind"`
	Parameter string              `json:"parameter,omitempty"` // Current name from the document parameter API.
	Precision *uint8              `json:"precision,omitempty"` // Model parameters default to three decimal places.
}

// ValidateSketchTextFields checks the source, order, and locations of fields.
// Pass separately validated fraction ranges and nil content to defer positions to the host.
// For example, ValidateSketchTextFields(&text, fields, nil) checks fields in new content.
func ValidateSketchTextFields(content *string, fields []SketchTextField, fractions []SketchTextFraction) error {
	if len(fields) == 0 {
		return nil
	}
	scalars, err := sketchFieldScalars(content)
	if err != nil {
		return err
	}
	for index, field := range fields {
		if index > 0 && field.At <= fields[index-1].At {
			return fmt.Errorf("sketch text field %d at %d: expected a scalar offset greater than %d", index, field.At, fields[index-1].At)
		}
		if err := field.validate(scalars, fractions); err != nil {
			return err
		}
	}
	return nil
}

func sketchFieldScalars(content *string) (*[]rune, error) {
	if content == nil {
		return nil, nil
	}
	if !utf8.ValidString(*content) {
		return nil, fmt.Errorf("linked sketch text %q: expected valid UTF-8", *content)
	}
	scalars := []rune(*content)
	return &scalars, nil
}

func (field SketchTextField) validate(scalars *[]rune, fractions []SketchTextFraction) error {
	if err := field.validateSource(); err != nil {
		return err
	}
	if scalars != nil && (uint64(field.At) >= uint64(len(*scalars)) || string((*scalars)[field.At]) != SketchTextFieldPlaceholder) {
		return fmt.Errorf("sketch text field at %d in %q: expected an object-replacement scalar U+FFFC", field.At, string(*scalars))
	}
	index := sort.Search(len(fractions), func(i int) bool { return field.At < fractions[i].End })
	if index < len(fractions) && field.At >= fractions[index].Start {
		return fmt.Errorf("sketch text field at %d in stacked range %+v: expected a location outside all stacked fractions", field.At, fractions[index])
	}
	return nil
}

func (field SketchTextField) validateSource() error {
	switch field.Kind {
	case SketchTextFieldModelParameter:
		if field.Parameter == "" || strings.TrimSpace(field.Parameter) != field.Parameter {
			return fmt.Errorf("sketch text field at %d parameter %q: expected a nonempty parameter name without surrounding spaces", field.At, field.Parameter)
		}
		if field.Precision != nil && *field.Precision > 8 {
			return fmt.Errorf("sketch text field at %d precision %d: expected zero through eight decimal places", field.At, *field.Precision)
		}
	case SketchTextFieldFileName:
		if field.Parameter != "" {
			return fmt.Errorf("file name field at %d parameter %q: expected no parameter name or precision", field.At, field.Parameter)
		}
		if field.Precision != nil {
			return fmt.Errorf("file name field at %d precision %d: expected no parameter name or precision", field.At, *field.Precision)
		}
	default:
		return fmt.Errorf("sketch text field at %d kind %q: expected modelParameter or fileName", field.At, field.Kind)
	}
	return nil
}
