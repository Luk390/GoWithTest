package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		caseName string
		input    interface{}
		expected []string
	}{
		{
			caseName: "struct with one string field",
			input: struct{
				Name string
			}{"Luke"},
			expected: []string{"Luke"},
		},
		{
			caseName: "struct with two string fields",
			input: struct {
				Name string
				City string
			}{
				Name: "Luke",
				City: "Bristol",
			},
			expected: []string{"Luke", "Bristol"},
		},
		{
			caseName: "struct with one non string field",
			input: struct {
				Name string
				Age  int
			}{
				Name: "Luke",
				Age:  32,
			},
			expected: []string{"Luke"},
		},
		{
			caseName: "struct with nested fields",
			input: Person{
				Name:    "Luke",
				Profile: Profile{32, "Bristol"},
			},
			expected: []string{"Luke", "Bristol"},
		},
		{
			caseName: "pointers to things",
			input: &Person{
				Name:    "Luke",
				Profile: Profile{32, "Bristol"},
			},
			expected: []string{"Luke", "Bristol"},
		},
		{
			caseName: "slices",
			input: []Profile{
				{32, "Bristol"},
				{25, "Worcester"},
			},
			expected: []string{"Bristol", "Worcester"},
		},
		{
			caseName: "arrays",
			input: [2]Profile{
				{32, "Bristol"},
				{25, "Worcester"},
			},
			expected: []string{"Bristol", "Worcester"},
		},
		{
			caseName: "maps",
			input: map[string]string{
				"Luke": "Bristol",
			},
			expected: []string{"Bristol"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.caseName, func(t *testing.T) {
			var got []string
			want := testCase.expected
			walk(testCase.input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"Luke": "Bristol",
			"Kasia": "La la Land",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bristol")
		assertContains(t, got, "La la Land")
	})
}
func assertContains(t testing.TB, got []string, want string) {
	t.Helper()

	contains := false
	for _, v := range got {
		if v == want {
			contains = true
		} 
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", got, want)
	}
}
