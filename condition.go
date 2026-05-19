package golidate

// Condition controls whether a result should be considered applicable.
type Condition func(value any) bool
