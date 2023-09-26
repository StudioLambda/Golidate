package golidate

import (
	"fmt"
	"reflect"
)

type Pending struct {
	value     any
	attribute string
	rules     []Rule
}

func Value(value any) Pending {
	return Pending{
		value:     value,
		attribute: "",
		rules:     make([]Rule, 0),
	}
}

func (pending Pending) Rules(rules ...Rule) Pending {
	pending.rules = rules

	return pending
}

func (pending Pending) AppendRules(rules ...Rule) Pending {
	pending.rules = append(pending.rules, rules...)

	return pending
}

func (pending Pending) Name(attribute string) Pending {
	pending.attribute = attribute

	return pending
}

func (pending Pending) attributeOfIndex(i int) string {
	if pending.attribute == "" {
		return fmt.Sprintf("%d", i)
	}

	return fmt.Sprintf("%s.%d", pending.attribute, i)
}

func (pending Pending) attributeOfKey(s string) string {
	if pending.attribute == "" {
		return s
	}

	return fmt.Sprintf("%s.%s", pending.attribute, s)
}

func (pending Pending) recursiveValidate() Results {
	if validatable, ok := pending.value.(Validater); ok {
		return validatable.Validate().Prefixed(pending.attribute)
	}

	results := make(Results, 0)

	reflected := reflect.ValueOf(pending.value)

	switch reflected.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < reflected.Len(); i++ {
			if validatable, ok := reflected.Index(i).Interface().(Validater); ok {
				name := pending.attributeOfIndex(i)
				results = append(results, validatable.Validate().Prefixed(name)...)
			}
		}
	case reflect.Map:
		for _, key := range reflected.MapKeys() {
			if validatable, ok := reflected.MapIndex(key).Interface().(Validater); ok {
				name := pending.attributeOfKey(key.String())
				results = append(results, validatable.Validate().Prefixed(name)...)
			}
		}
	}

	return results
}

func (pending Pending) Validate() Results {
	results := make(Results, 0, len(pending.rules))

	for _, rule := range pending.rules {
		result := rule(pending.value)

		if pending.attribute != "" {
			result = result.Name(pending.attribute)
		}

		expanded := result.Results(pending.attribute)

		results = append(results, expanded...)
	}

	return append(results, pending.recursiveValidate()...)
}
