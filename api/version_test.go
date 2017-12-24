package api_test

import (
	"reflect"
	"testing"

	"github.com/christiangda/mango/api"
)

func TestNewVerion(t *testing.T) {
	var tests = []struct {
		input string
		want  *api.Version
	}{
		{"1.0.1", &api.Version{Major: 1, Minor: 0, Patch: 1}},
		{"1.1.1", &api.Version{Major: 1, Minor: 1, Patch: 1}},
		{"1.0.a", nil},
		{"1.a.0", nil},
		{"a.0.0", nil},
		{"1.0.0.0", nil},
		{"", nil},
	}
	for _, test := range tests {
		if got, _ := api.NewVersion(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("NewVersion(%q) = %v, and want is = %v", test.input, got, test.want)
		}
	}
}

func TestGetVersion(t *testing.T) {
	var tests = []struct {
		input string
		want  *api.Version
	}{
		{"1.0.1", &api.Version{Major: 1, Minor: 0, Patch: 1}},
		{"1.1.1", &api.Version{Major: 1, Minor: 1, Patch: 1}},
		{"1.0.a", nil},
		{"1.a.0", nil},
		{"a.0.0", nil},
		{"1.0.0.0", nil},
		{"", nil},
	}

	for _, test := range tests {

		version, _ := api.NewVersion(test.input)

		if got := version.GetVersion(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("GetVersion(%q) = %v, and want is = %v", test.input, got, test.want)
		}
	}
}

func TestToString(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"1.0.1", "1.0.1"},
		{"1.1.1", "1.1.1"},
		{"1.0.a", ""},
		{"1.a.0", ""},
		{"a.0.0", ""},
		{"1.0.0.0", ""},
		{"", ""},
	}

	for _, test := range tests {

		version, _ := api.NewVersion(test.input)

		if got := version.ToString(); !reflect.DeepEqual(got, test.want) {
			t.Errorf("GetVersion(%q) = %v, and want is = %v", test.input, got, test.want)
		}
	}
}
