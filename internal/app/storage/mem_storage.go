package storage

import (
	"errors"
	"sync"
)

type MemStorage struct {
	storage map[string]string
	mutex   sync.RWMutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		storage: make(map[string]string),
		mutex:   sync.RWMutex{},
	}
}

func (repo *MemStorage) Save(url, id string) error {
	repo.mutex.RLock()
	_, ok := repo.storage[id]
	repo.mutex.RUnlock()

	if ok {
		return errors.New("not unique id")
	}

	repo.mutex.Lock()
	repo.storage[id] = url
	repo.mutex.Unlock()

	return nil
}

func (repo *MemStorage) GetByID(id string) (string, error) {
	repo.mutex.RLock()
	url, ok := repo.storage[id]
	repo.mutex.RUnlock()

	if !ok {
		return "", errors.New("can't find full url by id")
	}

	return url, nil
}
