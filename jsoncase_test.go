package main

import (
	"encoding/json"
	"testing"
)

func TestConvertKeysCase(t *testing.T) {
	inputJSON := `{
  "first": {
    "title": "Some title",
    "meta": {
      "author": "S",
      "entry": {
        "MoreMeta": {
          "ID": "1",
          "Name": "Json",
          "Date": "Today",
          "Age": 123,
          "valid": true
        }
      }
    }
  },
  "last": {
    "Title": "Another title"
  }
}`

	var jsonData map[string]any
	err := json.Unmarshal([]byte(inputJSON), &jsonData)
	if err != nil {
		t.Fatalf("Error unmarshalling input JSON: %v", err)
	}

	tests := []struct {
		caseFlag bool
		expected string
	}{
		{
			caseFlag: true,
			expected: `{"FIRST":{"META":{"AUTHOR":"S","ENTRY":{"MOREMETA":{"AGE":123,"DATE":"Today","ID":"1","NAME":"Json","VALID":true}}},"TITLE":"Some title"},"LAST":{"TITLE":"Another title"}}`,
		},
		{
			caseFlag: false,
			expected: `{"first":{"meta":{"author":"S","entry":{"moremeta":{"age":123,"date":"Today","id":"1","name":"Json","valid":true}}},"title":"Some title"},"last":{"title":"Another title"}}`,
		},
	}

	for _, test := range tests {
		modifiedData := convertKeysCase(jsonData, test.caseFlag)
		result, err := json.Marshal(modifiedData)
		if err != nil {
			t.Fatalf("Error marshalling modified JSON: %v", err)
		}

		if string(result) != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, string(result))
		}
	}
}
