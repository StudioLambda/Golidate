package golidate

// Metadata stores additional information about a validation result.
//
// Rule constructors use metadata for values that translators need, such as
// limits, allowed values, expected types, and nested operation results. Metadata
// keys are strings by convention and are referenced in translation templates as
// @key placeholders.
type Metadata map[string]any

// OnRename updates dependent result state after a result is renamed.
//
// Composite rules use rename callbacks to keep nested results in metadata in
// sync when a parent result receives a more specific attribute name.
type OnRename func(result Result)

// Result describes the outcome of one validation rule.
//
// Result values are immutable by convention: methods return a modified copy so
// callers can compose changes without mutating the original value. Children are
// kept out of JSON output because Results expands them before presentation.
type Result struct {
	// ok stores the direct pass state for this result.
	ok bool `json:"-"`
	// conditions stores applicability checks evaluated by Passes.
	conditions []Condition `json:"-"`
	// onRename updates nested metadata when Name changes Attribute.
	onRename OnRename `json:"-"`
	// children stores child results that share the parent attribute name.
	children Results `json:"-"`
	// prefixedChildren stores child results prefixed by the parent attribute.
	prefixedChildren Results `json:"-"`
	// Attribute names the field, value, key, or index that was validated.
	Attribute string `json:"attribute"`
	// Value stores the original value passed to the rule.
	Value any `json:"value"`
	// Code stores the stable machine-readable validation code.
	Code string `json:"code"`
	// Message stores the human-readable or translation-ready message.
	Message string `json:"message"`
	// Metadata stores translation data and nested operation details.
	Metadata Metadata `json:"metadata"`
}

// Results expands a result and its children into a flat result list.
//
// The receiver is always first in the returned slice. Non-prefixed children can
// inherit the provided name, while prefixed children receive the receiver's
// attribute as a dotted prefix.
func (result Result) Results(name string) Results {
	results := make(Results, 1, result.resultCount())

	results[0] = result

	for _, child := range result.children {
		if name != "" {
			results = append(results, child.Name(name).Results(name)...)
		} else {
			results = append(results, child.Results(name)...)
		}
	}

	for _, child := range result.prefixedChildren.Prefixed(result.Attribute) {
		results = append(results, child.Results(name)...)
	}

	return results
}

// resultCount returns the number of direct and nested results.
//
// It is used to preallocate flattened result slices without walking the same
// nested result tree repeatedly while appending.
func (result Result) resultCount() int {
	count := 1

	for _, child := range result.children {
		count += child.resultCount()
	}

	for _, child := range result.prefixedChildren {
		count += child.resultCount()
	}

	return count
}

// WithChild appends a child result that shares the parent attribute name.
//
// Shared children are useful for composite rules whose child operations all
// describe the same attribute as the parent result.
func (result Result) WithChild(child Result) Result {
	result.children = append(result.children, child)

	return result
}

// WithPrefixedChild appends a child result prefixed by the parent attribute.
//
// Prefixed children are used by collection rules so nested failures become
// names such as users.0.email or settings.[enabled].
func (result Result) WithPrefixedChild(child Result) Result {
	result.prefixedChildren = append(result.prefixedChildren, child)

	return result
}

// Name sets the result attribute name.
//
// If an OnRename callback is registered, the callback is invoked after the
// attribute changes so nested operation metadata can be renamed consistently.
func (result Result) Name(attribute string) Result {
	result.Attribute = attribute

	if result.onRename != nil {
		result.onRename(result)
	}

	return result
}

// OnRenameMany returns a callback that renames result metadata lists.
//
// The metadata value at key must be a Results value. Other metadata shapes are
// ignored, which keeps rename callbacks safe for custom metadata.
func OnRenameMany(key string) OnRename {
	return func(result Result) {
		if operations, ok := result.Metadata[key].(Results); ok {
			copied := make(Results, len(operations))
			copy(copied, operations)

			for i := range copied {
				copied[i] = copied[i].Name(result.Attribute)
			}

			result.Metadata[key] = copied
		}
	}
}

// OnRenameSingle returns a callback that renames result metadata values.
//
// The metadata value at key must be a Result value. Other metadata shapes are
// ignored.
func OnRenameSingle(key string) OnRename {
	return func(result Result) {
		if operation, ok := result.Metadata[key].(Result); ok {
			result.Metadata[key] = operation.Name(result.Attribute)
		}
	}
}

// OnRename sets a callback invoked when the result is renamed.
//
// Only one callback is stored. Setting a new callback replaces any previous
// callback on the copied result.
func (result Result) OnRename(callback OnRename) Result {
	result.onRename = callback

	return result
}

// With adds one metadata value to the result.
//
// With initializes Metadata when needed and replaces any existing value stored
// under the same key.
func (result Result) With(key string, value any) Result {
	return result.WithMetadataMerged(
		Metadata{
			key: value,
		},
	)
}

// WithMetadata replaces result metadata.
//
// Replacement is exact, including allowing nil metadata. Translation entries
// that depend on metadata placeholders may leave those placeholders unchanged if
// required metadata is removed.
func (result Result) WithMetadata(metadata Metadata) Result {
	result.Metadata = metadata

	return result
}

// WithMetadataMerged merges metadata into the result.
//
// Existing keys not present in metadata are preserved. Matching keys are
// overwritten by the supplied values, giving callers deterministic control over
// placeholder data used by translators.
func (result Result) WithMetadataMerged(metadata Metadata) Result {
	if result.Metadata == nil {
		result.Metadata = make(Metadata)
	}

	for key, value := range metadata {
		result.Metadata[key] = value
	}

	return result
}

// Pass marks the result as passing.
//
// Conditions may still cause Passes to report true for an otherwise failing
// result, but Pass sets the direct state to successful.
func (result Result) Pass() Result {
	result.ok = true

	return result
}

// Fail marks the result as failing.
//
// A failing direct state contributes to Failed results unless an applicability
// condition marks the result as irrelevant.
func (result Result) Fail() Result {
	result.ok = false

	return result
}

// Passes reports whether the direct result state passes.
//
// If any condition returns false, the result is considered to pass because the
// rule did not apply to the value. Otherwise the stored direct pass state is
// returned.
func (result Result) Passes() bool {
	for _, condition := range result.conditions {
		if !condition(result.Value) {
			return true
		}
	}

	return result.ok
}

// PassesChilds reports whether all child results pass.
//
// The method checks direct child results attached through WithChild and
// WithPrefixedChild. It does not recursively call PassesAll on grandchildren;
// callers that need full flattened behavior should use Results first.
func (result Result) PassesChilds() bool {
	for _, child := range result.children {
		if !child.Passes() {
			return false
		}
	}

	for _, child := range result.prefixedChildren {
		if !child.Passes() {
			return false
		}
	}

	return true
}

// PassesAll reports whether the result and all direct child results pass.
//
// This is the predicate used by Results.Failed and Results.Passed for each
// stored result.
func (result Result) PassesAll() bool {
	return result.Passes() && result.PassesChilds()
}

// Conditions sets applicability conditions on the result.
//
// Conditions replace any existing conditions on the copied result.
func (result Result) Conditions(conditions ...Condition) Result {
	result.conditions = conditions

	return result
}

// When sets a boolean applicability condition on the result.
//
// When(false) makes Passes return true regardless of the direct pass state,
// which is useful for conditionally skipping a rule after it has been built.
func (result Result) When(condition bool) Result {
	result.conditions = []Condition{
		func(any) bool {
			return condition
		},
	}

	return result
}

// Translate returns the result translated by the provided dictionaries.
//
// Dictionaries are merged once in argument order. Later dictionaries override
// earlier dictionaries with the same code. Missing entries leave the result
// unchanged.
func (result Result) Translate(dictionaries ...Dictionary) Result {
	dictionary := mergeDictionaries(dictionaries...)

	return result.translate(dictionary)
}

// translate applies one already-merged dictionary to the result.
//
// The helper avoids repeatedly merging dictionaries while translating slices of
// results and nested operation metadata.
func (result Result) translate(dictionary Dictionary) Result {
	if entry, ok := dictionary[result.Code]; ok {
		return entry(dictionary, result)
	}

	return result
}

// mergeDictionaries combines dictionaries using later entries as overrides.
//
// The returned map is newly allocated so translating does not mutate caller
// dictionaries or rely on shared package-level state.
func mergeDictionaries(dictionaries ...Dictionary) Dictionary {
	dictionary := make(Dictionary)

	for _, dict := range dictionaries {
		for key, value := range dict {
			dictionary[key] = value
		}
	}

	return dictionary
}

// Uncertain creates a result with an unset pass or fail state.
//
// The initial state is failing until Pass is called. Attribute defaults to
// "attribute", Message defaults to code, and Metadata is initialized so rule
// constructors can immediately attach placeholder data.
func Uncertain(value any, code string) Result {
	return Result{
		ok:               false,
		conditions:       make([]Condition, 0),
		children:         make(Results, 0),
		prefixedChildren: make(Results, 0),
		Attribute:        "attribute",
		Value:            value,
		Code:             code,
		Message:          code,
		Metadata:         make(Metadata),
	}
}

// Fail creates a failing result.
//
// It is equivalent to Uncertain(value, code).Fail().
func Fail(value any, code string) Result {
	return Uncertain(value, code).Fail()
}

// Pass creates a passing result.
//
// It is equivalent to Uncertain(value, code).Pass().
func Pass(value any, code string) Result {
	return Uncertain(value, code).Pass()
}
