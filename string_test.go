package pdfedit

import (
	"testing"
)

var stringValueTests = []struct {
	name     string
	string   *String
	expected string
}{
	{
		name:     "NilString",
		string:   nil,
		expected: "",
	},
	{
		name:     "ValueSet",
		string:   &String{value: "Hello world"},
		expected: "Hello world",
	},
}

func TestStringValue(t *testing.T) {
	for _, test := range stringValueTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.string.Value()

			if result != test.expected {
				t.Errorf("Expected `%s`. Got `%s`", test.expected, result)
			}
		})
	}
}

var stringSetValueTests = []struct {
	name     string
	string   *String
	value    string
	expected string
}{
	{
		name:     "NilString",
		string:   nil,
		value:    "Hello world",
		expected: "",
	},
	{
		name:     "ChangeValue",
		string:   &String{value: "Hello word"},
		value:    "Foo bar",
		expected: "Foo bar",
	},
}

func TestStringSetValue(t *testing.T) {
	for _, test := range stringSetValueTests {
		t.Run(test.name, func(t *testing.T) {
			test.string.SetValue(test.value)
			result := test.string.Value()

			if result != test.expected {
				t.Errorf("Expected `%s`. Got `%s`", test.expected, result)
			}
		})
	}
}

var isHexadecimalTests = []struct {
	name     string
	string   *String
	expected bool
}{
	{
		name:     "NilString",
		string:   nil,
		expected: false,
	},
	{
		name:     "IsHexadecimalFalse",
		string:   &String{isHexadecimal: false},
		expected: false,
	},
	{
		name:     "IsHexadecimalTrue",
		string:   &String{isHexadecimal: true},
		expected: true,
	},
}

func TestIsHexadecimal(t *testing.T) {
	for _, test := range isHexadecimalTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.string.IsHexadecimal()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
		})
	}
}

var setIsHexadecimalTests = []struct {
	name     string
	string   *String
	value    bool
	expected bool
}{
	{
		name:     "NilString",
		string:   nil,
		value:    true,
		expected: false,
	},
	{
		name:     "ChangeValue",
		string:   &String{isHexadecimal: false},
		value:    true,
		expected: true,
	},
}

func TestSetIsHexadecimal(t *testing.T) {
	for _, test := range setIsHexadecimalTests {
		t.Run(test.name, func(t *testing.T) {
			test.string.SetIsHexadecimal(test.value)
			result := test.string.IsHexadecimal()

			if result != test.expected {
				t.Errorf("Expected `%v`. Got `%v`", test.expected, result)
			}
		})
	}
}
