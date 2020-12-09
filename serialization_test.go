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
		{
			name:           "a simple string",
			input:          "a string",
			expectedJSON:   "\"a string\"",
			expectedYAMLV2: "a string\n",
			expectedYAMLV3: "a string\n",
		},
		{
			name:           "an empty string list",
			input:          []string{},
			expectedJSON:   "[]",
			expectedYAMLV2: "[]\n",
			expectedYAMLV3: "[]\n",
		},
		{
			name:           "a string list",
			input:          []string{"first string", "second string"},
			expectedJSON:   "[\"first string\",\"second string\"]",
			expectedYAMLV2: "- first string\n- second string\n",
			expectedYAMLV3: "- first string\n- second string\n",
		},
		{
			name:           "object without any data in array",
			input:          ObjectWithStringArray{},
			expectedJSON:   "{\"data\":null}",
			expectedYAMLV2: "data: []\n",
			expectedYAMLV3: "data: []\n",
		},
		{
			name: "object with data in array",
			input: ObjectWithStringArray{
				Data: []string{"first string", "second string"},
			},
			expectedJSON:   "{\"data\":[\"first string\",\"second string\"]}",
			expectedYAMLV2: "data:\n- first string\n- second string\n",
			expectedYAMLV3: "data:\n    - first string\n    - second string\n",
		},
		{
			name:           "object with omit empty string without data",
			input:          ObjectWithOmitEmptyStringArray{},
			expectedJSON:   "{}",
			expectedYAMLV2: "{}\n",
			expectedYAMLV3: "{}\n",
		},
		{
			name: "object with omit empty string with data",
			input: ObjectWithOmitEmptyStringArray{
				Data: []string{"first string", "second string"},
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
