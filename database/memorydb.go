package database

import (
	"fmt"
	"sync"
)

type todos map[int]string

type Database struct {
	mu     sync.Mutex
	data   todos
	lastId int
}

func CreateDatabase() *Database {
	return &Database{data: make(todos), lastId: -1}
}

func (db *Database) FetchTodo(id int) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	val, exists := db.data[id]

	if !exists {
		return "", fmt.Errorf("id %v doesn't exist", id)
	}

	return val, nil
}

func (db *Database) FetchAll() map[int]string {
	db.mu.Lock()
	defer db.mu.Unlock()
	todoList := make(map[int]string, len(db.data))

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

func (db *Database) UpdateTodo(id int, todoText string) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	val, exists := db.data[id]

	if !exists {
		return "", fmt.Errorf("id %v doesn't exist", id)
	}

	db.data[id] = todoText
	return val, nil // returns old value as the new one is already known by the caller
}

func (db *Database) DeleteTodo(id int) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	val, exists := db.data[id]

	if !exists {
		return "", fmt.Errorf("id %v doesn't exist", id)
	}

	return val, nil
}
