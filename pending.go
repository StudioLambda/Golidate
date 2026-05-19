package golidate

import (
	"context"
	"fmt"
	"reflect"
	"sort"
)

// Pending describes a value and the rules that will be applied to it.
//
// Pending is a small immutable builder. Methods return a modified copy so a
// base pending value can be reused safely with different names or rule sets.
type Pending struct {
	// value stores the value that rules and recursive validation inspect.
	value any
	// attribute stores the result attribute assigned to direct rule results.
	attribute string
	// rules stores direct rules evaluated before recursive child validation.
	rules []Rule
	// self disables Validator recursion for the top-level value when true.
	self bool
}

// pending creates a Pending value and normalizes non-validator pointers.
//
// Values that implement Validator are kept intact so pointer-receiver validators
// can be discovered. Other non-nil pointers are dereferenced once so rules check
// the pointed-to value rather than the pointer itself.
func pending(value any, self bool) Pending {
	if _, ok := value.(Validator); ok {
		return Pending{
			value:     value,
			attribute: "",
			rules:     make([]Rule, 0),
			self:      self,
		}
	}

	reflected := reflect.ValueOf(value)
	if reflected.Kind() == reflect.Ptr && !reflected.IsNil() {
		return Pending{
			value:     reflected.Elem().Interface(),
			attribute: "",
			rules:     make([]Rule, 0),
			self:      self,
		}
	}

	return Pending{
		value:     value,
		attribute: "",
		rules:     make([]Rule, 0),
		self:      self,
	}
}

// Value starts validation for a value and recursively validates children.
//
// If the value implements Validator, Validate delegates to that implementation.
// Slices, arrays, and maps are also scanned for elements that implement
// Validator. Map element validation is deterministic because keys are sorted by
// their formatted names before traversal.
func Value(value any) Pending {
	return pending(value, false)
}

// Self starts validation for a value without treating it as a Validator.
//
// Self is useful when the value itself has a Validate method but the caller
// wants to apply explicit rules to the value instead of recursing into it.
func Self(value any) Pending {
	return pending(value, true)
}

// Rules replaces the rules that will be applied to the pending value.
//
// Rules are evaluated in the order supplied. Existing rules on the copied
// Pending are discarded.
func (pending Pending) Rules(rules ...Rule) Pending {
	pending.rules = rules

	return pending
}

// AppendRules appends rules that will be applied to the pending value.
//
// Existing rules keep their order and new rules are evaluated afterward.
func (pending Pending) AppendRules(rules ...Rule) Pending {
	pending.rules = append(pending.rules, rules...)

	return pending
}

// Name assigns the attribute name used by generated results.
//
// Direct rule results receive this attribute. Nested collection and validator
// results use it as a prefix for child attribute names.
func (pending Pending) Name(attribute string) Pending {
	pending.attribute = attribute

	return pending
}

// attributeOfIndex formats the attribute for an indexed child value.
//
// Root slices and arrays use bare numeric attributes, while named parents use a
// dotted suffix such as items.0.
func (pending Pending) attributeOfIndex(i int) string {
	if pending.attribute == "" {
		return fmt.Sprintf("%d", i)
	}

	return fmt.Sprintf("%s.%d", pending.attribute, i)
}

// attributeOfKey formats the attribute for a mapped child value.
//
// Root maps use the formatted key alone, while named parents use a dotted suffix
// such as settings.enabled.
func (pending Pending) attributeOfKey(s string) string {
	if pending.attribute == "" {
		return s
	}

	return fmt.Sprintf("%s.%s", pending.attribute, s)
}

// recursiveValidate validates nested Validator values when Value is used.
//
// The top-level Validator shortcut is skipped when the Pending was created by
// Self. Nil and non-collection values simply produce no recursive results.
func (pending Pending) recursiveValidate(ctx context.Context) Results {
	if validatable, ok := pending.value.(Validator); ok && !pending.self {
		return validatable.Validate(ctx).Prefixed(pending.attribute)
	}

	results := make(Results, 0)
	reflected := reflect.ValueOf(pending.value)

	switch reflected.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < reflected.Len(); i++ {
			if validatable, ok := reflected.Index(i).Interface().(Validator); ok {
				name := pending.attributeOfIndex(i)
				results = append(results, validatable.Validate(ctx).Prefixed(name)...)
			}
		}
	case reflect.Map:
		keys := sortedMapKeys(reflected)
		for _, key := range keys {
			if validatable, ok := reflected.MapIndex(key.value).Interface().(Validator); ok {
				name := pending.attributeOfKey(key.name)
				results = append(results, validatable.Validate(ctx).Prefixed(name)...)
			}
		}
	}

	return results
}

// mapKey stores a reflect map key with its deterministic display name.
type mapKey struct {
	// name is the formatted key text used for sorting and attributes.
	name string
	// value is the original reflect key used to read the map value.
	value reflect.Value
}

// sortedMapKeys returns map keys ordered by formatted key text.
//
// Go deliberately randomizes map iteration. Sorting here keeps nested
// validation output stable for tests, logs, API responses, and users.
func sortedMapKeys(reflected reflect.Value) []mapKey {
	keys := reflected.MapKeys()
	sorted := make([]mapKey, len(keys))

	for i, key := range keys {
		sorted[i] = mapKey{
			name:  fmt.Sprintf("%v", key.Interface()),
			value: key,
		}
	}

	sort.Slice(sorted, func(i int, j int) bool {
		return sorted[i].name < sorted[j].name
	})

	return sorted
}

// Validate applies the pending rules and recursive child validation.
//
// Direct rule results are expanded before recursive validator results are
// appended. Expansion keeps parent results and nested child results in one flat
// ordered slice.
func (pending Pending) Validate(ctx context.Context) Results {
	results := make(Results, 0, len(pending.rules))

	for _, rule := range pending.rules {
		result := rule(pending.value)

		if pending.attribute != "" {
			result = result.Name(pending.attribute)
		}

		expanded := result.Results(pending.attribute)

		results = append(results, expanded...)
	}

	return append(results, pending.recursiveValidate(ctx)...)
}
