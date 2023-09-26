package golidate

type Dictionary map[string]Entry

type Entry func(dictionary Dictionary, result Result) Result
