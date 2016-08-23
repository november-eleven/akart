package main

import (
	"reflect"
	"testing"
)

func TestEngine(t *testing.T) {

	e := NewEngine()

	var tests = []struct {
		search   string
		expected []string
	}{
		{"r", []string{"Reprobe", "River"}},
		{"p", []string{"Pandora", "Paypal", "Pg&e", "Phone"}},
		{"pr", []string{"Prank", "Press democrat", "Print", "Proactive"}},
		{"pro", []string{"Proactive", "Processor", "Procurable", "Progenex"}},
		{"prog", []string{"Progenex", "Progeria", "Progesterone", "Programming"}},
	}

	for _, test := range tests {

		r := e.Find(test.search)

		if r == nil {
			t.Fatal("Unexpected nil suggestions")
		}

		if !reflect.DeepEqual(r, test.expected) {
			t.Fatalf("Unexpected suggestions (%v): %v\n", test.expected, r)
		}

		t.Logf("%s: %v", test.search, r)

	}

}
