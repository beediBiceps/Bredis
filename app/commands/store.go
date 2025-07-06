package commands

import (
    "sync"
    "time"
)

type StoreItem struct{
    value string
    expiresAt *time.Time
}

type Store struct{
    store map[string]*StoreItem
    mutex sync.RWMutex
}

func NewStore() *Store{
    return &Store{store: make(map[string]*StoreItem)}
}

func (s *Store) Set(key string, value string) error{
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.store[key] = &StoreItem{value: value}
    return nil
}

func (s *Store) SetWithExpiry(key string, value string, expiryMs int64) error{
    s.mutex.Lock()
    defer s.mutex.Unlock()
    expiresAt := time.Now().Add(time.Duration(expiryMs) * time.Millisecond)
    s.store[key] = &StoreItem{value: value, expiresAt: &expiresAt}
    return nil
}

func (s *Store) Get(key string) (string, error){
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    item, exists := s.store[key]
    if !exists {
        return "", nil
    }
    
    if item.expiresAt != nil && time.Now().After(*item.expiresAt) {
        delete(s.store, key)
        return "", nil
    }
    
    return item.value, nil
}