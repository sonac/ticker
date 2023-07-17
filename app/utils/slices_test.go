package utils

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		input    []string
		expected []string
	}{
		{[]string{"one", "two", "three", "one", "two", "four"}, []string{"four", "one", "three", "two"}},
		{[]string{"apple", "orange", "apple", "apple", "banana"}, []string{"apple", "banana", "orange"}},
		{[]string{}, []string{}},
	}

	for _, test := range tests {
		RemoveDuplicates(&test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Error("Test failed: input and expected do not match")
		}
	}
}
