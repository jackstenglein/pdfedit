package pdfedit

import (
	"testing"
)

var typeTests = []struct {
	name    string
	boolean *Boolean
}{
	{
		name:    "NilBoolean",
		boolean: nil,
	},
	{
		name:    "DefaultBoolean",
		boolean: &Boolean{},
	},
}

func TestType(t *testing.T) {
	for _, test := range typeTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.boolean.Type()

			if result != "Boolean" {
				t.Errorf("Expected `Boolean`. Got `%s`", result)
			}
			if err != nil {
				t.Errorf("Expected nil. Got `%v`", err)
			}
		})
	}
}

var isIndirectTests = []struct {
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
		name:     "ZeroObjectNumber",
		boolean:  &Boolean{objectNumber: 0},
		expected: false,
	},
	{
		name:     "PositiveObjectNumber",
		boolean:  &Boolean{objectNumber: 1},
		expected: true,
	},
}

func TestIsIndirect(t *testing.T) {
	for _, test := range isIndirectTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.boolean.IsIndirect()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
			if err != nil {
				t.Errorf("Expected nil. Got `%v`", err)
			}
		})
	}
}

var objectNumberTests = []struct {
	name     string
	boolean  *Boolean
	expected int
}{
	{
		name:     "NilBoolean",
		boolean:  nil,
		expected: 0,
	},
	{
		name:     "ObjectNumberSet",
		boolean:  &Boolean{objectNumber: 3},
		expected: 3,
	},
}

func TestObjectNumber(t *testing.T) {
	for _, test := range objectNumberTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.boolean.ObjectNumber()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
			if err != nil {
				t.Errorf("Expected nil. Got `%v`", err)
			}
		})
	}
}

var generationNumberTests = []struct {
	name     string
	boolean  *Boolean
	expected int
}{
	{
		name:     "NilBoolean",
		boolean:  nil,
		expected: -1,
	},
	{
		name:     "GenerationNumberSet",
		boolean:  &Boolean{generationNumber: 5},
		expected: 5,
	},
}

func TestGenerationNumber(t *testing.T) {
	for _, test := range generationNumberTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.boolean.GenerationNumber()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
			if err != nil {
				t.Errorf("Expected nil. Got `%v`", err)
			}
		})
	}
}

var childrenTests = []struct {
	name    string
	boolean *Boolean
}{
	{
		name:    "NilBoolean",
		boolean: nil,
	},
	{
		name:    "DefaultBoolean",
		boolean: &Boolean{},
	},
}

func TestChildren(t *testing.T) {
	for _, test := range childrenTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.boolean.Children()

			if result != nil {
				t.Errorf("Expected nil. Got `%v`", result)
			}
			if err != nil {
				t.Errorf("Expected nil. Got `%v`", err)
			}
		})
	}
}

var valueTests = []struct {
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

func TestValue(t *testing.T) {
	for _, test := range valueTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.boolean.Value()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
		})
	}
}

var setValueTests = []struct {
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

func TestSetValue(t *testing.T) {
	for _, test := range setValueTests {
		t.Run(test.name, func(t *testing.T) {
			test.boolean.SetValue(test.value)
			result := test.boolean.Value()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
		})
	}
}
