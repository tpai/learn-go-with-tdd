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
	t.Run("with data types but map", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Tony", 30},
				[]string{"Tony"},
			},
			{
				"struct with one string field",
				struct {
					Name string
				}{"Tony"},
				[]string{"Tony"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Tony", "Taipei"},
				[]string{"Tony", "Taipei"},
			},
			{
				"nested fields",
				Person{
					"Tony",
					Profile{30, "Taipei"},
				},
				[]string{"Tony", "Taipei"},
			},
			{
				"pointer to things",
				&Person{
					"Tony",
					Profile{30, "Taipei"},
				},
				[]string{"Tony", "Taipei"},
			},
			{
				"slices",
				[]Profile{
					Profile{25, "Taichung"},
					Profile{30, "Taipei"},
				},
				[]string{"Taichung", "Taipei"},
			},
			{
				"arrays",
				[2]Profile{
					Profile{25, "Taichung"},
					Profile{30, "Taipei"},
				},
				[]string{"Taichung", "Taipei"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})
				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("expected %v but got %v", test.ExpectedCalls, got)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, got []string, expected string) {
	contains := false
	for _, x := range got {
		if x == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contains '%s' but it didn't", got, expected)
	}
}
