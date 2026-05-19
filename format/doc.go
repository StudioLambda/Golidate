// Package format provides display-oriented message formatters.
//
// Formatters are applied after validation and optional translation. They do not
// change result codes, metadata, or pass state; they only transform strings
// returned by Results.Messages or Grouped.Messages.
package format
