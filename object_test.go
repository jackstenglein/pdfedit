package pdfedit

import (
	"reflect"
	"testing"
)

var nilBoolean *Boolean
var nilString *String

var typeTests = []struct {
	name         string
	object       Object
	expectedType string
	expectedErr  error
}{
	{
		name:         "Boolean",
		object:       &Boolean{},
		expectedType: "Boolean",
		expectedErr:  nil,
	},
	{
		name:         "String",
		object:       &String{},
		expectedType: "String",
		expectedErr:  nil,
	},
}

func TestType(t *testing.T) {
	for _, test := range typeTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.object.Type()

			if result != test.expectedType {
				t.Errorf("Expected `%s`. Got `%s`", test.expectedType, result)
			}
			if err != test.expectedErr {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedErr, err)
			}
		})
	}
}

var isIndirectTests = []struct {
	name           string
	object         Object
	expectedResult bool
	expectedErr    error
}{
	{
		name:           "NilBoolean",
		object:         nilBoolean,
		expectedResult: false,
		expectedErr:    nil,
	},
	{
		name:           "DirectBoolean",
		object:         &Boolean{objectNumber: 0},
		expectedResult: false,
		expectedErr:    nil,
	},
	{
		name:           "IndirectBoolean",
		object:         &Boolean{objectNumber: 1},
		expectedResult: true,
		expectedErr:    nil,
	},
	{
		name:           "NilString",
		object:         nilString,
		expectedResult: false,
		expectedErr:    nil,
	},
	{
		name:           "DirectString",
		object:         &String{objectNumber: 0},
		expectedResult: false,
		expectedErr:    nil,
	},
	{
		name:           "IndirectString",
		object:         &String{objectNumber: 1},
		expectedResult: true,
		expectedErr:    nil,
	},
}

func TestIsIndirect(t *testing.T) {
	for _, test := range isIndirectTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.object.IsIndirect()

			if result != test.expectedResult {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedResult, result)
			}
			if err != test.expectedErr {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedErr, err)
			}
		})
	}
}

var objectNumberTests = []struct {
	name           string
	object         Object
	expectedResult int
	expectedErr    error
}{
	{
		name:           "NilBoolean",
		object:         nilBoolean,
		expectedResult: 0,
		expectedErr:    nil,
	},
	{
		name:           "Boolean",
		object:         &Boolean{objectNumber: 3},
		expectedResult: 3,
		expectedErr:    nil,
	},
	{
		name:           "NilString",
		object:         nilString,
		expectedResult: 0,
		expectedErr:    nil,
	},
	{
		name:           "String",
		object:         &String{objectNumber: 5},
		expectedResult: 5,
		expectedErr:    nil,
	},
}

func TestObjectNumber(t *testing.T) {
	for _, test := range objectNumberTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.object.ObjectNumber()

			if result != test.expectedResult {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedResult, result)
			}
			if err != test.expectedErr {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedErr, err)
			}
		})
	}
}

var generationNumberTests = []struct {
	name           string
	object         Object
	expectedResult int
	expectedErr    error
}{
	{
		name:           "NilBoolean",
		object:         nilBoolean,
		expectedResult: -1,
		expectedErr:    nil,
	},
	{
		name:           "Boolean",
		object:         &Boolean{generationNumber: 3},
		expectedResult: 3,
	},
	{
		name:           "NilString",
		object:         nilString,
		expectedResult: -1,
		expectedErr:    nil,
	},
	{
		name:           "String",
		object:         &String{generationNumber: 5},
		expectedResult: 5,
	},
}

func TestGenerationNumber(t *testing.T) {
	for _, test := range generationNumberTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.object.GenerationNumber()

			if result != test.expectedResult {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedResult, result)
			}
			if err != nil {
				t.Errorf("Expected nil. Got `%v`", err)
			}
		})
	}
}

var childrenTests = []struct {
	name           string
	object         Object
	expectedResult []*Object
	expectedErr    error
}{
	{
		name:           "Boolean",
		object:         &Boolean{},
		expectedResult: nil,
	},
	{
		name:           "String",
		object:         &String{},
		expectedResult: nil,
	},
}

func TestChildren(t *testing.T) {
	for _, test := range childrenTests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.object.Children()

			if !reflect.DeepEqual(result, test.expectedResult) {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedResult, result)
			}
			if err != test.expectedErr {
				t.Errorf("Expected `%v`. Got `%v`", test.expectedErr, err)
			}
		})
	}
}
