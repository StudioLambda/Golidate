package golidate

type Metadata map[string]any

type OnRename func(result Result)

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

func (result Result) WithChild(child Result) Result {
	result.children = append(result.children, child)

	return result
}

func (result Result) WithPrefixedChild(child Result) Result {
	result.prefixedChildren = append(result.prefixedChildren, child)

	return result
}

func (result Result) Name(attribute string) Result {
	result.Attribute = attribute

	if result.onRename != nil {
		result.onRename(result)
	}

	return result
}

func OnRenameMany(key string) OnRename {
	return func(result Result) {
		if operations, ok := result.Metadata[key].(Results); ok {
			for i := range operations {
				operations[i] = operations[i].Name(result.Attribute)
			}
		}
	}
}

func OnRenameSingle(key string) OnRename {
	return func(result Result) {
		if operation, ok := result.Metadata[key].(Result); ok {
			result.Metadata[key] = operation.Name(result.Attribute)
		}
	}
}

func (result Result) OnRename(callback OnRename) Result {
	result.onRename = callback

	return result
}

func (result Result) With(key string, value any) Result {
	return result.WithMetadataMerged(
		Metadata{
			key: value,
		},
	)
}

func (result Result) WithMetadata(metadata Metadata) Result {
	result.Metadata = metadata

	return result
}

func (result Result) WithMetadataMerged(metadata Metadata) Result {
	for key, value := range metadata {
		result.Metadata[key] = value
	}

	return result
}

func (result Result) Pass() Result {
	result.ok = true

	return result
}

func (result Result) Fail() Result {
	result.ok = false

	return result
}

func (result Result) Passes() bool {
	for _, condition := range result.conditions {
		if !condition(result.Value) {
			return true
		}
	}

	return result.ok
}

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

func (result Result) PassesAll() bool {
	return result.Passes() && result.PassesChilds()
}

func (result Result) Conditions(conditions ...Condition) Result {
	result.conditions = conditions

	return result
}

func (result Result) When(condition bool) Result {
	result.conditions = []Condition{
		func(any) bool {
			return condition
		},
	}

	return result
}

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

func Fail(value any, code string) Result {
	return Uncertain(value, code).Fail()
}

func Pass(value any, code string) Result {
	return Uncertain(value, code).Pass()
}
