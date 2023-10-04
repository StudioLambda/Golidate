package golidate

import (
	"context"
	"fmt"
	"reflect"
)

type Pending struct {
	value     any
	attribute string
	rules     []Rule
}

func Value(value any) Pending {
	reflected := reflect.ValueOf(value)
	if reflected.Kind() == reflect.Ptr && !reflected.IsNil() {
		return Pending{
			value:     reflected.Elem().Interface(),
			attribute: "",
			rules:     make([]Rule, 0),
		}
	}

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

func (pending Pending) recursiveValidate(ctx context.Context) Results {
	if validatable, ok := pending.value.(Validator); ok {
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
		for _, key := range reflected.MapKeys() {
			if validatable, ok := reflected.MapIndex(key).Interface().(Validator); ok {
				name := pending.attributeOfKey(key.String())
				results = append(results, validatable.Validate(ctx).Prefixed(name)...)
			}
		}
	}

	return results
}

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
