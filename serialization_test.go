package main

import (
	"encoding/json"
	"testing"

	yamlV2 "gopkg.in/yaml.v2"
	yamlV3 "gopkg.in/yaml.v3"
)

func TestSerialization(t *testing.T) {
	testcases := []struct {
		name           string
		input          interface{}
		expectedJSON   string
		expectedYAMLV2 string
		expectedYAMLV3 string
	}{
		// OK
		{
			name:           "a string",
			input:          "a string",
			expectedJSON:   "\"a string\"",
			expectedYAMLV2: "a string\n",
			expectedYAMLV3: "a string\n",
		},
		// OK
		{
			name:           "an empty string list",
			input:          []string{},
			expectedJSON:   "[]",
			expectedYAMLV2: "[]\n",
			expectedYAMLV3: "[]\n",
		},
		// OK
		{
			name:           "a string array",
			input:          []string{"first string", "second string"},
			expectedJSON:   "[\"first string\",\"second string\"]",
			expectedYAMLV2: "- first string\n- second string\n",
			expectedYAMLV3: "- first string\n- second string\n",
		},

		// How we can separate nil and [] ?
		// We can't! nil and the empty array are exported as empty array.
		{
			name:           "ObjectWithStringArray without data",
			input:          ObjectWithStringArray{},
			expectedJSON:   "{\"data\":null}",
			expectedYAMLV2: "data: []\n",
			expectedYAMLV3: "data: []\n",
		},
		{
			name: "ObjectWithStringArray with data: nil",
			input: ObjectWithStringArray{
				Data: nil,
			},
			expectedJSON:   "{\"data\":null}",
			expectedYAMLV2: "data: []\n",
			expectedYAMLV3: "data: []\n",
		},
		{
			name: "ObjectWithStringArray with data: []string{}",
			input: ObjectWithStringArray{
				Data: []string{},
			},
			expectedJSON:   "{\"data\":[]}",
			expectedYAMLV2: "data: []\n",
			expectedYAMLV3: "data: []\n",
		},
		{
			name: "ObjectWithStringArray with data in array",
			input: ObjectWithStringArray{
				Data: []string{"first string", "second string"},
			},
			expectedJSON:   "{\"data\":[\"first string\",\"second string\"]}",
			expectedYAMLV2: "data:\n- first string\n- second string\n",
			expectedYAMLV3: "data:\n    - first string\n    - second string\n",
		},

		// How we can separate nil and [] ?
		// We can! nil is exported as null and an empty array as empty array.
		{
			name:           "ObjectWithPointerToStringArray without data",
			input:          ObjectWithPointerToStringArray{},
			expectedJSON:   "{\"data\":null}",
			expectedYAMLV2: "data: null\n",
			expectedYAMLV3: "data: null\n",
		},
		{
			name: "ObjectWithPointerToStringArray with data: nil",
			input: ObjectWithPointerToStringArray{
				Data: nil,
			},
			expectedJSON:   "{\"data\":null}",
			expectedYAMLV2: "data: null\n",
			expectedYAMLV3: "data: null\n",
		},
		{
			name: "ObjectWithPointerToStringArray with data: []string{}",
			input: ObjectWithPointerToStringArray{
				Data: &[]string{},
			},
			expectedJSON:   "{\"data\":[]}",
			expectedYAMLV2: "data: []\n",
			expectedYAMLV3: "data: []\n",
		},
		{
			name: "ObjectWithPointerToStringArray with data in array",
			input: ObjectWithPointerToStringArray{
				Data: &[]string{"first string", "second string"},
			},
			expectedJSON:   "{\"data\":[\"first string\",\"second string\"]}",
			expectedYAMLV2: "data:\n- first string\n- second string\n",
			expectedYAMLV3: "data:\n    - first string\n    - second string\n",
		},

		// How we can separate nil and [] ?
		// We can't! nil and empty arrays are not exported.
		{
			name:           "ObjectWithOmitEmptyStringArray without data",
			input:          ObjectWithOmitEmptyStringArray{},
			expectedJSON:   "{}",
			expectedYAMLV2: "{}\n",
			expectedYAMLV3: "{}\n",
		},
		{
			name: "ObjectWithOmitEmptyStringArray with data: nil",
			input: ObjectWithOmitEmptyStringArray{
				Data: nil,
			},
			expectedJSON:   "{}",
			expectedYAMLV2: "{}\n",
			expectedYAMLV3: "{}\n",
		},
		{
			name: "ObjectWithOmitEmptyStringArray with data: []string{}",
			input: ObjectWithOmitEmptyStringArray{
				Data: []string{},
			},
			expectedJSON:   "{}",
			expectedYAMLV2: "{}\n",
			expectedYAMLV3: "{}\n",
		},
		{
			name: "ObjectWithOmitEmptyStringArray with data in array",
			input: ObjectWithOmitEmptyStringArray{
				Data: []string{"first string", "second string"},
			},
			expectedJSON:   "{\"data\":[\"first string\",\"second string\"]}",
			expectedYAMLV2: "data:\n- first string\n- second string\n",
			expectedYAMLV3: "data:\n    - first string\n    - second string\n",
		},

		// How we can separate nil and [] ?
		// We can! nil is not exported, and an empty array as empty array.
		{
			name:           "ObjectWithOmitEmptyPointerToStringArray without data",
			input:          ObjectWithOmitEmptyPointerToStringArray{},
			expectedJSON:   "{}",
			expectedYAMLV2: "{}\n",
			expectedYAMLV3: "{}\n",
		},
		{
			name: "ObjectWithOmitEmptyPointerToStringArray with data: nil",
			input: ObjectWithOmitEmptyPointerToStringArray{
				Data: nil,
			},
			expectedJSON:   "{}",
			expectedYAMLV2: "{}\n",
			expectedYAMLV3: "{}\n",
		},
		{
			name: "ObjectWithOmitEmptyPointerToStringArray with data: []string{}",
			input: ObjectWithOmitEmptyPointerToStringArray{
				Data: &[]string{},
			},
			expectedJSON:   "{\"data\":[]}",
			expectedYAMLV2: "data: []\n",
			expectedYAMLV3: "data: []\n",
		},
		{
			name: "ObjectWithOmitEmptyPointerToStringArray with data in array",
			input: ObjectWithOmitEmptyPointerToStringArray{
				Data: &[]string{"first string", "second string"},
			},
			expectedJSON:   "{\"data\":[\"first string\",\"second string\"]}",
			expectedYAMLV2: "data:\n- first string\n- second string\n",
			expectedYAMLV3: "data:\n    - first string\n    - second string\n",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			actualJSON, err := json.Marshal(testcase.input)
			actualYAMLV2, err := yamlV2.Marshal(testcase.input)
			actualYAMLV3, err := yamlV3.Marshal(testcase.input)

			if err != nil {
				t.Error(err)
			}
			if string(actualJSON) != testcase.expectedJSON {
				t.Errorf("Unexpected json\nactual:\n%v\nexpected:\n%v\n", string(actualJSON), testcase.expectedJSON)
			}
			if string(actualYAMLV2) != testcase.expectedYAMLV2 {
				t.Errorf("Unexpected yaml v2\nactual:\n%v\nexpected:\n%v\n", string(actualYAMLV2), testcase.expectedYAMLV2)
			}
			if string(actualYAMLV3) != testcase.expectedYAMLV3 {
				t.Errorf("Unexpected yaml v3\nactual:\n%v\nexpected:\n%v\n", string(actualYAMLV3), testcase.expectedYAMLV3)
			}
		})
	}
}
