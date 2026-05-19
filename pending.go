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
	self      bool
}

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
func Value(value any) Pending {
	return pending(value, false)
}

// Self starts validation for a value without treating it as a Validator.
func Self(value any) Pending {
	return pending(value, true)
}

// Rules replaces the rules that will be applied to the pending value.
func (pending Pending) Rules(rules ...Rule) Pending {
	pending.rules = rules

	return pending
}

// AppendRules appends rules that will be applied to the pending value.
func (pending Pending) AppendRules(rules ...Rule) Pending {
	pending.rules = append(pending.rules, rules...)

	return pending
}

// Name assigns the attribute name used by generated results.
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
		for _, key := range reflected.MapKeys() {
			if validatable, ok := reflected.MapIndex(key).Interface().(Validator); ok {
				name := pending.attributeOfKey(fmt.Sprintf("%v", key.Interface()))
				results = append(results, validatable.Validate(ctx).Prefixed(name)...)
			}
		}
	}

	return results
}

// Validate applies the pending rules and recursive child validation.
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
