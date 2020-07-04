package pdfedit

import (
	"testing"
)

var booleanValueTests = []struct {
	name     string
	boolean  *Boolean
	expected bool
}{
	{
		name:     "NilBoolean",
		boolean:  nil,
		expected: false,
	},
	{
		name:     "ValueTrue",
		boolean:  &Boolean{value: true},
		expected: true,
	},
	{
		name:     "ValueFalse",
		boolean:  &Boolean{value: false},
		expected: false,
	},
}

func TestBooleanValue(t *testing.T) {
	for _, test := range booleanValueTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.boolean.Value()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
		})
	}
}

var booleanSetValueTests = []struct {
	name     string
	boolean  *Boolean
	value    bool
	expected bool
}{
	{
		name:     "NilBoolean",
		boolean:  nil,
		value:    true,
		expected: false,
	},
	{
		name:     "ValueTrue",
		boolean:  &Boolean{value: false},
		value:    true,
		expected: true,
	},
	{
		name:     "ValueFalse",
		boolean:  &Boolean{value: true},
		value:    false,
		expected: false,
	},
}

func TestBooleanSetValue(t *testing.T) {
	for _, test := range booleanSetValueTests {
		t.Run(test.name, func(t *testing.T) {
			test.boolean.SetValue(test.value)
			result := test.boolean.Value()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
		})
	}
}
