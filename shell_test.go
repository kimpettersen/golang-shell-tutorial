package main

import (
	"reflect"
	"testing"
)

func TestTokenizeStrings(t *testing.T) {

	var cases = []struct {
		in       string
		expected []string
	}{
		{
			"mkdir folder",
			[]string{"mkdir", "folder"},
		},
	}

	for _, tt := range cases {
		actual := Tokenize(tt.in)

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf(
				"Tokenize(%v): expected %v, actual %v",
				tt.in,
				tt.expected,
				actual,
			)
		}
	}
}
