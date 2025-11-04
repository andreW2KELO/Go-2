package main

import "log"

type DataRetriever interface {
	Retrieve(ID string) (*Data, error)
}

type Data struct {
	ID string // для упрощения содержит только ID
}

type Cache struct {
	// данные будем хранить здесь
	m  map[string]*Data
	dr DataRetriever
}

func NewCache(dr DataRetriever) *Cache {
	return &Cache{
		m:  make(map[string]*Data),
		dr: dr,
	}
}

func (c *Cache) Get(ID string) (Data, bool) {
	// проверим, есть ли данные в кэше
	data, exists := c.m[ID]
	// нашли в мапе — вернём значение
	if exists {
		return *data, true
	}
	// данные не нашли — нужно запросить
	data, err := c.dr.Retrieve(ID)
	if err != nil {
		// ошибка получения данных — запишем в лог
		log.Printf("c.dr.Retrieve(ID): %s", err)
		// вернём пустое значение
		return Data{}, false
	}
	// получили значение — запомним
	c.m[data.ID] = data
	// вернём полученное значение
	return *data, true
}

// Это вполне рабочий вариант, но только если мы используем его в одной горутине.
// При чтении кэша и записи в него несколькими горутинами мы получим состояние гонки (race condition).

// В состоянии гонки (race condition) несколько потоков (или процессов) одновременно пытаются выполнить операции чтения и записи.
