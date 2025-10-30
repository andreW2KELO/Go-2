package main

import (
	"log"
	"sync"
)

type DataRetriever interface {
	Retrieve(ID string) (*Data, error)
}

type Data struct {
	ID string // для упрощения содержит только ID
}

// Представьте: мы хотим посчитать какое-то значение, прибавить к нему 1, а потом перезаписать.
// Если эту операцию произвести в двух горутинах, которые работают параллельно, они одновременно считают старое
// значение, прибавят к нему 1, получат одинаковый результат, запишут его, и значение увеличится только на 1.
// Если бы состояния гонки не было, горутины отработали бы последовательно и прибавили к предыдущему результату 2.
// Кэш обычно используется многими горутинами, поэтому он должен быть потокобезопасным. Для этого добавим мьютекс:

func NewCache(dr DataRetriever) *Cache {
	return &Cache{
		m:  make(map[string]*Data),
		dr: dr,
	}
}

type Cache struct {
	m  map[string]*Data
	dr DataRetriever
	mu sync.Mutex
}

func (c *Cache) Get(ID string) (Data, bool) {
	c.mu.Lock()
	data, exists := c.m[ID] // теперь доступ к мапе внутри критической секции
	c.mu.Unlock()
	// нашли в мапе — вернём значение
	if exists {
		return *data, true
	}
	// запрос данных из базы — не в критической секции
	data, err := c.dr.Retrieve(ID)
	if err != nil {
		// ошибка получения данных — запишем в лог
		log.Printf("c.dr.Retrieve(ID): %s", err)
		// вернём пустое значение
		return Data{}, false
	}
	// перед обращением к мапе снова заблокируем мьютекс
	c.mu.Lock()
	// разблокируем при выходе из функциии
	defer c.mu.Unlock()
	// внутри критической секции нужно снова проверить на наличие значения в мапе
	data, exists = c.m[data.ID]
	if exists {
		return *data, true
	}
	// получили значение — запомним
	c.m[data.ID] = data
	// вернём полученное значение
	return *data, true
}
