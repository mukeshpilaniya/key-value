package db

import (
	"errors"
)

// Store a type for inmemory database
type Store struct {
	store map[string]string
}

// NewStore return a new instance pointing to Store type
func (db *Store) NewStore() *Store {
	return &Store{
		store: make(map[string]string),
	}
}

// AddKey store key and value into the database
func (db *Store) AddKey(key string, value string) {
	db.store[key] = value
}

// GetValue return a value corresponding to key
func (db *Store) GetValue(key string) (string, error) {
	if _, ok := db.store[key]; !ok {
		return "", errors.New("key is not present")
	} else {
		return db.store[key], nil
	}
}

// GetAllKeys return all the keys form database
func (db *Store) GetAllKeys() []string {
	var keys []string

	for k, _ := range db.store {
		keys = append(keys, k)
	}
	return keys
}

// GetAllValue return all the values form database
func (db *Store) GetAllValue() []string {
	var values []string

	for _, v := range db.store {
		values = append(values, v)
	}
	return values
}
