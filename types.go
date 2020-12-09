package main

// ObjectWithStringArray is used in tests
type ObjectWithStringArray struct {
	Data []string `json:"data" yaml:"data"`
}

// ObjectWithOmitEmptyStringArray is used in tests
type ObjectWithOmitEmptyStringArray struct {
	Data []string `json:"data,omitempty" yaml:"data,omitempty"`
}
