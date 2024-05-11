package main

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	cases := []struct {
		operations []struct {
			method string
			data   any
		}
		expected []any
	}{
		{
			operations: []struct {
				method string
				data   any
			}{
				{"pushToHead", "john"},
				{"getSize", nil},
				{"getNodeAtPosition", 0},
				{"pushToHead", "bane"},
				{"pushToHead", "lily"},
				{"getSize", nil},
				{"getNodeAtPosition", 2},
				{"getNodeAtPosition", 1},
				{"getNodeAtPosition", 0},
			},
			expected: []any{
				nil,
				1,
				"john",
				nil,
				nil,
				3,
				"john",
				"bane",
				"lily",
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"popFromHead", nil},
				{"getSize", nil},
				{"pushToHead", "bane"},
				{"pushToHead", "lily"},
				{"popFromHead", nil},
				{"getSize", nil},
				{"getNodeAtPosition", 0},
			},
			expected: []any{
				nil,
				0,
				nil,
				nil,
				"lily",
				1,
				"bane",
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"pushToTail", "john"},
				{"getSize", nil},
				{"getNodeAtPosition", 0},
				{"pushToTail", "bane"},
				{"pushToTail", "lily"},
				{"getSize", nil},
				{"getNodeAtPosition", 2},
				{"getNodeAtPosition", 1},
				{"getNodeAtPosition", 0},
			},
			expected: []any{
				nil,
				1,
				"john",
				nil,
				nil,
				3,
				"lily",
				"bane",
				"john",
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"popFromTail", nil},
				{"getSize", nil},
				{"pushToTail", "bane"},
				{"pushToTail", "lily"},
				{"popFromTail", nil},
				{"getSize", nil},
				{"getNodeAtPosition", 0},
			},
			expected: []any{
				nil,
				0,
				nil,
				nil,
				"lily",
				1,
				"bane",
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"pushToTail", "bane"},
				{"pushToTail", "lily"},
				{"getNodeAtPosition", 0},
				{"getNodeAtPosition", 1},
				{"getSize", nil},
				{"pushToPosition", []any{"john", 1}},
				{"getNodeAtPosition", 0},
				{"getNodeAtPosition", 1},
				{"getNodeAtPosition", 2},
				{"getSize", nil},
			},
			expected: []any{
				nil,
				nil,
				"bane",
				"lily",
				2,
				nil,
				"bane",
				"john",
				"lily",
				3,
			},
		},
		{
			operations: []struct {
				method string
				data   any
			}{
				{"pushToTail", "bane"},
				{"pushToTail", "lily"},
				{"pushToTail", "john"},
				{"getNodeAtPosition", 0},
				{"getNodeAtPosition", 1},
				{"getNodeAtPosition", 2},
				{"getSize", nil},
				{"popFromPosition", 1},
				{"getNodeAtPosition", 0},
				{"getNodeAtPosition", 1},
				{"getSize", nil},
			},
			expected: []any{
				nil,
				nil,
				nil,
				"bane",
				"lily",
				"john",
				3,
				"lily",
				"bane",
				"john",
				2,
			},
		},
	}

	for _, c := range cases {
		dl := doublyLinkedList{}
		for i, o := range c.operations {
			switch o.method {
			case "pushToHead":
				dl.pushToHead(o.data)
			case "popFromHead":
				got, expected := dl.popFromHead().data, c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			case "pushToTail":
				dl.pushToTail(o.data)
			case "popFromTail":
				got, expected := dl.popFromTail().data, c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			case "pushToPosition":
				d := o.data.([]any)
				dl.pushToPosition(d[0], d[1].(int))
			case "popFromPosition":
				got, expected := dl.popFromPosition(o.data.(int)).data, c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			case "getSize":
				got, expected := dl.getSize(), c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			case "getNodeAtPosition":
				got, expected := dl.getNodeAtPosition(o.data.(int)).data, c.expected[i]
				if got != expected {
					testHelper(t, i, got, expected)
				}
			}
		}
	}
}

func testHelper(t *testing.T, idx int, got, expected any) {
	t.Errorf("expected IDX: %v | got: %v != expected: %v", idx, got, expected)
}
