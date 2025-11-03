package tasks

import (
	"testing"
)

// TestContains содержит набор тестовых случаев для функции Contains.
func TestContains(t *testing.T) {
	// Структура для определения тестового случая
	type testCase struct {
		name     string // Название теста
		numbers  []int  // Входной срез чисел
		target   int    // Искомое число
		expected bool   // Ожидаемый результат
	}

	// Таблица тестовых случаев
	tests := []testCase{
		{
			name:     "Содержит_число_в_середине",
			numbers:  []int{1, 5, 10, 15, 20},
			target:   10,
			expected: true,
		},
		{
			name:     "Содержит_число_в_начале",
			numbers:  []int{1, 5, 10, 15, 20},
			target:   1,
			expected: true,
		},
		{
			name:     "Содержит_число_в_конце",
			numbers:  []int{1, 5, 10, 15, 20},
			target:   20,
			expected: true,
		},
		{
			name:     "Не_содержит_число",
			numbers:  []int{1, 5, 10, 15, 20},
			target:   7,
			expected: false,
		},
		{
			name:     "Пустой_срез",
			numbers:  []int{},
			target:   5,
			expected: false,
		},
		{
			name:     "Срез_из_одного_элемента_на_совпадение",
			numbers:  []int{42},
			target:   42,
			expected: true,
		},
		{
			name:     "Срез_из_одного_элемента_на_отсутствие",
			numbers:  []int{42},
			target:   1,
			expected: false,
		},
	}

	// Итерация по тестовым случаям
	for _, tc := range tests {
		// Запуск каждого теста как подтеста для лучшей изоляции и отчетности
		t.Run(tc.name, func(t *testing.T) {
			// Вызов тестируемой функции
			actual := Contains(tc.numbers, tc.target)

			// Проверка результата
			if actual != tc.expected {
				t.Errorf("Contains(%v, %d) = %v; ожидалось %v", tc.numbers, tc.target, actual, tc.expected)
			}
		})
	}
}