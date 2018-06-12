package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected []rune
	}{
		{
			name:     "with newline",
			s:        "1\n2",
			expected: []rune{'1', '2'},
		},
		{
			name:     "with whitespace",
			s:        "1 2",
			expected: []rune{'1', '0', '2'},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Tokenize(tc.s)

			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("expected result: %c, saw: %c", tc.expected, result)
			}
		})
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		name     string
		pos      [][]int
		chars    []rune
		expected string
	}{
		{
			name:     "has enough chars",
			pos:      [][]int{{0, 1}, {1, 0}},
			chars:    []rune{'a', 'b'},
			expected: " a\nb \n",
		},
		{
			name:     "has not enough chars",
			pos:      [][]int{{0, 1}, {1, 0}},
			chars:    []rune{'a'},
			expected: " a\na \n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := String(tc.pos, tc.chars)

			if result != tc.expected {
				t.Errorf("expected result: %v, saw: %v", tc.expected, result)
			}
		})
	}
}

func TestPosition(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected [][]int
	}{
		{
			name:     "single line",
			s:        "a",
			expected: [][]int{{1}},
		},
		{
			name:     "multiline",
			s:        " first\nsecond",
			expected: [][]int{{0, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Position(tc.s)

			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("expected result: %v, saw: %v", tc.expected, result)
			}
		})
	}
}
