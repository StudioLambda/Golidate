package golidate

// Condition controls whether a result should be considered applicable.
//
// A condition receives the same value that was checked by the rule. When any
// condition returns false, Result.Passes treats the result as passing because
// the rule was not relevant for that value.
type Condition func(value any) bool
