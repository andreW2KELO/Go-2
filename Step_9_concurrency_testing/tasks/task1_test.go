package tasks

import (
	"reflect"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"обычный случай", []int{5, 3, 8, 1, 2}, []int{1, 2, 3, 5, 8}},
		{"уже отсортирован", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"обратный порядок", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"повторяющиеся элементы", []int{4, 2, 2, 1}, []int{1, 2, 2, 4}},
		{"один элемент", []int{10}, []int{10}},
		{"пустой слайс", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			SortIntegers(inputCopy)

			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("SortIntegers(%v) = %v, ожидалось %v",
					tt.input, inputCopy, tt.expected)
			}
		})
	}
}
