package golidate

// Metadata stores additional information about a validation result.
type Metadata map[string]any

// OnRename updates dependent result state after a result is renamed.
type OnRename func(result Result)

// Result describes the outcome of one validation rule.
type Result struct {
	ok               bool        `json:"-"`
	conditions       []Condition `json:"-"`
	onRename         OnRename    `json:"-"`
	children         Results     `json:"-"`
	prefixedChildren Results     `json:"-"`
	Attribute        string      `json:"attribute"`
	Value            any         `json:"value"`
	Code             string      `json:"code"`
	Message          string      `json:"message"`
	Metadata         Metadata    `json:"metadata"`
}

// Results expands a result and its children into a flat result list.
func (result Result) Results(name string) Results {
	results := make(Results, 1)

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

// WithChild appends a child result that shares the parent attribute name.
func (result Result) WithChild(child Result) Result {
	result.children = append(result.children, child)

	return result
}

// WithPrefixedChild appends a child result prefixed by the parent attribute.
func (result Result) WithPrefixedChild(child Result) Result {
	result.prefixedChildren = append(result.prefixedChildren, child)

	return result
}

// Name sets the result attribute name.
func (result Result) Name(attribute string) Result {
	result.Attribute = attribute

	if result.onRename != nil {
		result.onRename(result)
	}

	return result
}

// OnRenameMany returns a callback that renames result metadata lists.
func OnRenameMany(key string) OnRename {
	return func(result Result) {
		if operations, ok := result.Metadata[key].(Results); ok {
			for i := range operations {
				operations[i] = operations[i].Name(result.Attribute)
			}
		}
	}
}

// OnRenameSingle returns a callback that renames result metadata values.
func OnRenameSingle(key string) OnRename {
	return func(result Result) {
		if operation, ok := result.Metadata[key].(Result); ok {
			result.Metadata[key] = operation.Name(result.Attribute)
		}
	}
}

// OnRename sets a callback invoked when the result is renamed.
func (result Result) OnRename(callback OnRename) Result {
	result.onRename = callback

	return result
}

// With adds metadata to the result.
func (result Result) With(key string, value any) Result {
	return result.WithMetadataMerged(
		Metadata{
			key: value,
		},
	)
}

// WithMetadata replaces result metadata.
func (result Result) WithMetadata(metadata Metadata) Result {
	result.Metadata = metadata

	return result
}

// WithMetadataMerged merges metadata into the result.
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
func (result Result) Pass() Result {
	result.ok = true

	return result
}

// Fail marks the result as failing.
func (result Result) Fail() Result {
	result.ok = false

	return result
}

// Passes reports whether the direct result state passes.
func (result Result) Passes() bool {
	for _, condition := range result.conditions {
		if !condition(result.Value) {
			return true
		}
	}

	return result.ok
}

// PassesChilds reports whether all child results pass.
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

// PassesAll reports whether the result and all child results pass.
func (result Result) PassesAll() bool {
	return result.Passes() && result.PassesChilds()
}

// Conditions sets applicability conditions on the result.
func (result Result) Conditions(conditions ...Condition) Result {
	result.conditions = conditions

	return result
}

// When sets a boolean applicability condition on the result.
func (result Result) When(condition bool) Result {
	result.conditions = []Condition{
		func(any) bool {
			return condition
		},
	}

	return result
}

// Translate returns the result translated by the provided dictionaries.
func (result Result) Translate(dictionaries ...Dictionary) Result {
	dictionary := make(Dictionary)

	for _, dict := range dictionaries {
		for key, value := range dict {
			dictionary[key] = value
		}
	}

	if entry, ok := dictionary[result.Code]; ok {
		return entry(dictionary, result)
	}

	return result
}

// Uncertain creates a result with an unset pass or fail state.
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
func Fail(value any, code string) Result {
	return Uncertain(value, code).Fail()
}

// Pass creates a passing result.
func Pass(value any, code string) Result {
	return Uncertain(value, code).Pass()
}
