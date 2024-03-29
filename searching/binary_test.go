package searching

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actual, actualErr := BinarySearch(test.data, test.key)
		if actual != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actual)
		}
		if actualErr != test.expectedError {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualErr)
		}
	}
}

type searchTest struct {
	data          []int
	key           int
	expected      int
	expectedError error
	name          string
}

var searchTests = []searchTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 8, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 3, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil, "Sanity"},
	//Absent
	{[]int{1, 4, 5, 6, 7, 10}, -25, -1, ErrNotFound, "Absent"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrNotFound, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, ErrNotFound, "Empty"},
}

var lowerBoundTests = []searchTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -25, 0, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9, nil, "Sanity"},
	{[]int{1, 2, 2, 2, 2, 6, 7, 8, 9, 10}, 2, 1, nil, "Sanity"},
	{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 2, 0, nil, "Sanity"},
	//Absent
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrNotFound, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, ErrNotFound, "Empty"},
}

var uppperBoundTests = []searchTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -25, 0, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 1, nil, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 5, nil, "Sanity"},
	{[]int{1, 2, 2, 2, 2, 6, 7, 8, 9, 10}, 2, 5, nil, "Sanity"},
	//Absent
	{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 2, -1, ErrNotFound, "Sanity"},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, -1, ErrNotFound, "Sanity"},
	{[]int{1, 4, 5, 6, 7, 10}, 25, -1, ErrNotFound, "Absent"},
	//Empty slice
	{[]int{}, 2, -1, ErrNotFound, "Empty"},
}

func Test_calcsqrt2(t *testing.T) {
	// tests := []struct {
	// 	name string
	// 	want float64
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := calcsqrt2(); got != tt.want {
	// 			t.Errorf("calcsqrt2() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }

	res := Calcsqrt2()
	fmt.Println(res)
}
