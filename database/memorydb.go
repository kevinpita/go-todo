package database

import (
	"fmt"
	"sync"
)

type Todos map[int]string

type Database struct {
	mu     sync.RWMutex
	data   Todos
	lastId int
}

func CreateDatabase() *Database {
	return &Database{data: make(Todos), lastId: -1}
}

func (db *Database) FetchTodo(id int) (string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	val, exists := db.data[id]

	if !exists {
		return "", fmt.Errorf("id %v doesn't exist", id)
	}

	return val, nil
}

func (db *Database) FetchAll() Todos {
	db.mu.RLock()
	defer db.mu.RUnlock()
	todoList := make(Todos, len(db.data))

	for k, v := range db.data {
		todoList[k] = v
	}

	return todoList
}

func (db *Database) CreateTodo(todoText string) int {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.lastId++
	db.data[db.lastId] = todoText

	return db.lastId
}

func (db *Database) UpdateTodo(id int, todoText string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.data[id]

	if !exists {
		return fmt.Errorf("id %v doesn't exist", id)
	}

	db.data[id] = todoText
	return nil
}

func (db *Database) DeleteTodo(id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.data[id]

	if !exists {
		return fmt.Errorf("id %v doesn't exist", id)
	}

	return nil
}
