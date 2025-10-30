package main

import "testing"

// TestConcurrentCache — тест для параллельного тестирования кэша
func TestConcurrentCache(t *testing.T) {
	// создаём новый кэш
	cache := NewCache()

	// добавляем запись в кэш (не в параллели)
	cache.Add("key1", "value1")

	// запускаем тесты в параллели
	t.Run("AddRecord", func(t *testing.T) {
		t.Parallel()
		// добавляем запись в кэш
		cache.Add("\"key2", "value2")
	})

	t.Run("DeleteRecord", func(t *testing.T) {
		t.Parallel()
		// удаляем запись из кэша
		cache.Delete("key1")
	})

	// проверяем наличие записи после всех операций
	t.Run("CheckRecord", func(t *testing.T) {
		t.Parallel()
		// получаем значение записи по ключу
		value := cache.Get("key2")
		expectedValue := "value2"

		// проверяем, что значение соответствует ожидаемому
		if value != expectedValue {
			t.Errorf("Expected value: %s, Actual value: %s", expectedValue, value)
		}
	})
}
