# Golidate

Golidate is a small Go validation library built around plain functions, explicit
results, and human-readable messages. It is useful when you want validation to be
easy to compose, inspect, translate, test, and present without adopting a large
framework or a struct-tag DSL.

The library gives you:

- Composable validation rules that are just Go functions.
- Structured results with codes, attributes, values, metadata, and children.
- Nested validation for values that implement `golidate.Validator`.
- Deterministic map validation output for stable tests and API responses.
- Translation dictionaries that turn stable codes into user-facing messages.
- Formatters for final presentation details such as capitalization.

## Installation

```sh
go get github.com/studiolambda/golidate
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"

	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/format"
	"github.com/studiolambda/golidate/rule"
	"github.com/studiolambda/golidate/translate/language"
)

func main() {
	results := golidate.Validate(
		context.Background(),
		golidate.Value("erik42").Name("username").Rules(
			rule.Alpha(),
		),
	)

	messages := results.
		Failed().
		Translate(language.English).
		Messages(format.Capitalize(), format.Punctuate())

	for _, message := range messages {
		fmt.Println(message)
	}

	// Output:
	// The username field must only contain letters.
}
```

## Core Concepts

### Rule

A `golidate.Rule` is a function:

```go
type Rule func(value any) Result
```

Rules inspect one value and return one `Result`. Rule constructors live in the
`rule` package. You can also write your own rule with `golidate.Pass`,
`golidate.Fail`, or `golidate.Uncertain`.

Rules can be wrapped:

```go
rule.MinLen(8).
	Code("password_length").
	Message("password_length").
	With("hint", "use at least eight characters")
```

Use `Code` for stable machine-readable identifiers, `Message` for untranslated
fallback text, and metadata for values that translation templates need.

### Result

A `golidate.Result` describes one validation outcome. It stores:

- `Attribute`: the field, index, key, or logical value name.
- `Value`: the original value checked by the rule.
- `Code`: the stable validation code.
- `Message`: a message or message key.
- `Metadata`: extra values used by translators or callers.

Results are copied by builder-style methods, so this is safe and predictable:

```go
result := golidate.Fail("abc", "min_len").
	Name("password").
	With("min", 8)
```

### Results

`golidate.Results` is a slice of `Result` values. It has helpers for common
presentation and decision points:

```go
if results.PassesAll() {
	// Continue.
}

failed := results.Failed()
messages := failed.Translate(language.English).Messages()
grouped := failed.Group().Messages()
```

An empty `Results` passes `PassesAll` because no failure exists.

### Pending

`golidate.Pending` is the builder created by `golidate.Value` or
`golidate.Self`:

```go
golidate.Value("Ada").Name("name").Rules(rule.Alpha(), rule.MinLen(2))
```

`Value` enables recursive validation for values that implement
`golidate.Validator`. `Self` applies the supplied rules to the value itself even
when it has a `Validate` method.

### Grouped

`Grouped` is a map from attribute names to their results. It is useful when you
want form-style or JSON API errors keyed by field name:

```go
messagesByField := results.Failed().Translate(language.English).Group().Messages()
```

### Validator

Any value with this method can validate itself:

```go
type Validator interface {
	Validate(ctx context.Context) golidate.Results
}
```

A validator normally composes smaller pending validations:

```go
type User struct {
	Email string
	Name  string
}

func (user User) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(user.Email).Name("email").Rules(rule.Email()),
		golidate.Value(user.Name).Name("name").Rules(rule.Required(rule.MinLen(2))),
	)
}
```

Then validate the whole value:

```go
results := golidate.Value(User{Email: "bad", Name: ""}).Validate(context.Background())
```

### Formatter

A `Formatter` transforms final strings only. It does not change codes, metadata,
or pass/fail state.

Built-in formatters:

- `format.Capitalize()` uppercases the first rune.
- `format.Uncapitalize()` lowercases the first rune.
- `format.Punctuate()` appends a period when missing.

### Translator, Dictionary, and Entries

A `golidate.Dictionary` maps result codes to translation entries. An entry
receives the merged dictionary and the result, then returns an updated result.

```go
dictionary := golidate.Dictionary{
	"username_required": translate.Simple(":attribute is required"),
}
```

Use the built-in English dictionary for the core rule catalog:

```go
messages := results.Failed().Translate(language.English).Messages()
```

Dictionaries are layered in order. Later dictionaries override earlier ones:

```go
messages := results.Failed().Translate(
	language.English,
	golidate.Dictionary{
		"email": translate.Simple("please enter a real email address for :attribute"),
	},
).Messages()
```

`translate.Simple` supports these placeholders:

- `:attribute`: the result attribute.
- `:code`: the result code.
- `:message`: the current result message.
- `:value`: the checked value formatted with `%+v`.
- `@key`: metadata by key, such as `@min`, `@max`, or `@values`.

Metadata placeholders are replaced in deterministic longest-first order, so
overlapping names such as `@min` and `@minimum` do not corrupt each other.

## Rule Catalog

The `rule` package contains the core rules. Every rule returns a `Result` with a
stable code. Most codes match the constructor name in snake case.

### Accepted

```go
rule.Accepted()
```

Passes for the strings `"true"`, `"1"`, `"on"`, and `"yes"`. It uses strict
membership, so the Go boolean `true` does not pass.

### Declined

```go
rule.Declined()
```

Passes for the strings `"false"`, `"0"`, `"off"`, and `"no"`.

### Boolean

```go
rule.Boolean()
```

Passes for string values commonly used in forms: `"true"`, `"false"`, `"1"`,
`"0"`, `"on"`, `"off"`, `"yes"`, and `"no"`. It does not accept Go `bool`
values.

### Alpha and Text Shape Rules

```go
rule.Alpha()
rule.AlphaNumeric()
rule.AlphaDash()
rule.AlphaExtended()
rule.Ascii()
```

These rules are regular-expression based and format values before matching.
Empty strings pass these shape checks because their regular expressions allow
zero characters. Use `rule.Required(...)` or `rule.MinLen(1)` when blank values
should fail.

### Email, Domain, URL, and MAC

```go
rule.Email()
rule.Domain()
rule.Url()
rule.Mac()
```

These rules are pragmatic regex checks. They are useful for user input feedback,
but they do not perform DNS lookup, network validation, full RFC mailbox
validation, or exhaustive URL parsing.

### Regex

```go
rule.Regex(regexp.MustCompile(`^[A-Z]{2}-\d+$`))
```

Formats the value with `fmt.Sprintf("%+v", value)` and matches the expression.
The expression string is stored as metadata key `regex`.

### HasPrefix and HasSuffix

```go
rule.HasPrefix("INV-")
rule.HasSuffix(".csv")
```

These rules format the value and then use `strings.HasPrefix` or
`strings.HasSuffix`.

### Lowercase and Uppercase

```go
rule.Lowercase()
rule.Uppercase()
```

These rules inspect letters in the formatted value. Digits, spaces, punctuation,
and symbols are ignored.

### Min and Max

```go
rule.Min(18)
rule.Max(120)
```

Numeric rules accept signed integers, unsigned integers, and floats. Values are
converted to `float64` for comparison, so decimal values work with integer
limits:

```go
rule.Min(10)(10.5).Passes() // true
rule.Max(10)(10.5).Passes() // false
```

Nil and non-numeric values fail safely.

### Len, MinLen, and MaxLen

```go
rule.Len(2)
rule.MinLen(8)
rule.MaxLen(255)
```

Length rules support arrays, channels, maps, slices, and strings. Nil values and
unsupported types fail safely without panic recovery.

### In and InTyped

```go
rule.In("small", "medium", "large")
rule.InTyped[int](1, 2, 3)
```

Membership is intentionally strict.

`In` requires the same dynamic type and `reflect.DeepEqual`:

```go
rule.In(int64(1))(int(1)).Passes() // false
rule.In("1")(1).Passes()          // false
```

`InTyped[T]` first type-asserts the value to `T`, then uses comparable equality.
Use it when you know the exact type and want compile-time safety for the allowed
values.

### Equal

```go
rule.Equal(expected)
```

Uses `reflect.DeepEqual`, so it follows Go's equality semantics for slices,
maps, structs, pointers, and nils.

### Type, TypeOf, and Convertible

```go
rule.Type[string]()
rule.TypeOf(time.Time{})
rule.Convertible[int64]()
```

`Type` and `TypeOf` require exact dynamic type equality. `Convertible` accepts
values whose type can be converted to the target type according to Go reflection.
Nil actual or expected types fail safely.

### Nil and Zero

```go
rule.Nil()
rule.Zero()
```

`Nil` passes for nil interfaces and nil-capable values whose reflected value is
nil. Non-nil values fail.

`Zero` passes when a value is the zero value of its concrete type. A nil
interface fails because there is no concrete type to inspect.

### Required and Optional

```go
rule.Required(rule.Email())
rule.Optional(rule.Email())
```

`Required` passes when the value is not nil and the nested rule passes.
`Optional` passes when the value is nil or the nested rule passes.

### And, Or, and Not

```go
rule.And(rule.MinLen(8), rule.HasPrefix("app_"))
rule.Or(rule.Email(), rule.Domain())
rule.Not(rule.In("admin", "root"))
```

Composite rules store nested operation results in metadata. The English
dictionary uses that metadata to build messages such as joined `and also` or
`or else` text and inverted `must` wording.

### SliceValues

```go
rule.SliceValues[[]string, string](rule.MinLen(2))
```

Applies child rules to every element in a slice of the exact generic type. Child
attributes use numeric indexes and are prefixed by the parent result when
expanded.

### MapKeys and MapValues

```go
rule.MapKeys[map[string]int, string, int](rule.Alpha())
rule.MapValues[map[string]int, string, int](rule.Min(1))
```

Applies child rules to every key or value in a map of the exact generic type.
Keys are formatted and sorted before validation, which keeps result ordering
deterministic despite Go's randomized map iteration.

## Nested Validation

`golidate.Value` recursively validates values that implement `Validator`.

```go
type Profile struct {
	DisplayName string
}

func (profile Profile) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(profile.DisplayName).Name("display_name").Rules(rule.MinLen(2)),
	)
}

type Account struct {
	Profile Profile
}

func (account Account) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(account.Profile).Name("profile"),
	)
}
```

Nested results are prefixed:

```text
profile.display_name
```

Slices and arrays containing validators use numeric indexes:

```text
users.0.email
users.1.email
```

Maps containing validators use formatted keys in deterministic sorted order:

```text
members.admin.email
members.guest.email
```

## Pointer Receiver Validation

Values with pointer-receiver `Validate` methods work when the value passed to
`Value` is a pointer:

```go
type Token struct {
	Value string
}

func (token *Token) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(token.Value).Name("value").Rules(rule.MinLen(10)),
	)
}

results := golidate.Value(&Token{Value: "short"}).Name("token").Validate(ctx)
```

Golidate keeps validator pointers intact so pointer methods can be discovered.
For non-validator pointers, it dereferences one level before applying direct
rules.

Use `golidate.Self` when a value implements `Validator` but you want to validate
the value itself with explicit rules instead of recursing into `Validate`.

## Nil and Reflection Edge Behavior

Rules that inspect reflection values are designed to fail safely on nil or
unsupported values.

- Numeric rules return failure for nil and non-numeric values.
- Length rules return failure for nil and unsupported kinds.
- Type rules return failure when actual or expected types are nil.
- `Nil` treats nil interfaces and nil-capable nil values as passing.
- `Zero` fails for a nil interface because no concrete zero value can be known.

This makes validation safe for optional request data and partially populated
structures without relying on panics for normal control flow.

## Strict Membership and Equality

Membership checks intentionally avoid broad coercion. This prevents surprising
passes such as a string form value matching a numeric allowed value.

```go
rule.In(1)(1).Passes()       // true
rule.In(int64(1))(1).Passes() // false
rule.In("1")(1).Passes()     // false
```

Use `InTyped[T]` when all allowed values are one comparable type. Use `Equal`
when you want Go's `reflect.DeepEqual` behavior for complex values.

## Translation and Formatting Examples

```go
results := golidate.Validate(
	ctx,
	golidate.Value("bad-email").Name("email").Rules(rule.Email()),
)

messages := results.Failed().Translate(language.English).Messages()
// []string{"the email field must be a valid email address"}
```

Add final display polish with formatters:

```go
messages := results.
	Failed().
	Translate(language.English).
	Messages(format.Capitalize(), format.Punctuate())
// []string{"The email field must be a valid email address."}
```

Override selected messages:

```go
custom := golidate.Dictionary{
	"min_len": translate.Simple("please make :attribute at least @min characters"),
}

messages := results.Failed().Translate(language.English, custom).Messages()
```

## Testing Examples

Validation results are plain Go data, so tests can assert pass state, messages,
codes, attributes, and metadata directly.

```go
func TestUsernameMessage(t *testing.T) {
	results := golidate.Validate(
		context.Background(),
		golidate.Value("erik42").Name("username").Rules(rule.Alpha()),
	)

	messages := results.Failed().Translate(language.English).Messages()

	if len(messages) != 1 {
		t.Fatalf("expected one message, got %d", len(messages))
	}

	if messages[0] != "the username field must only contain letters" {
		t.Fatalf("unexpected message: %q", messages[0])
	}
}
```

For deterministic map behavior, assert exact order when needed:

```go
results := golidate.Value(map[string]int{"b": 0, "a": 0}).Name("scores").Rules(
	rule.MapValues[map[string]int, string, int](rule.Min(1)),
).Validate(context.Background())

messages := results.Failed().Translate(language.English).Messages()
_ = messages // Child results are emitted in formatted key order.
```

## Writing Clear Validation Messages

Good validation messages help a person fix their input quickly.

- Name the field in the user's language, not the internal variable name, when
  possible.
- Say what the value must be, not just that it is invalid.
- Include limits and allowed values through metadata placeholders.
- Avoid exposing implementation details such as regular expressions unless the
  user can act on them.
- Prefer one specific message over several vague messages.
- Use `Code` for stable program behavior and translations for human wording.

Examples:

```go
rule.MinLen(8).With("hint", "use at least eight characters")
```

```go
golidate.Dictionary{
	"password_length": translate.Simple("the :attribute field must be at least @min characters"),
}
```

## Composing Validations

Start with simple rules and compose only when the resulting message remains
clear.

Use multiple pending validations when you want separate messages:

```go
golidate.Validate(
	ctx,
	golidate.Value(password).Name("password").Rules(rule.MinLen(8)),
	golidate.Value(password).Name("password").Rules(rule.HasPrefix("app_")),
)
```

Use `And`, `Or`, or `Not` when the logic is naturally one rule:

```go
rule.Or(rule.Email(), rule.Domain())
```

Use custom rules when:

- The validation depends on your domain language.
- The check needs application data or context.
- A composed message would be confusing.
- You need metadata tailored to your translation templates.

Custom rule example:

```go
func ReservedName(reserved map[string]struct{}) golidate.Rule {
	return func(value any) golidate.Result {
		name, ok := value.(string)
		result := golidate.Uncertain(value, "reserved_name")

		if !ok {
			return result.Fail()
		}

		if _, found := reserved[name]; found {
			return result.Fail()
		}

		return result.Pass()
	}
}
```

## Development

Run the standard Go checks before committing changes:

```sh
go test ./...
go vet ./...
go test -race ./...
```

Keep the repository focused on the Go library. Do not add Node, npm, or docs-site
tooling for library documentation.
