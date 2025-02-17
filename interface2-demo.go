package main

import "fmt"

var cache = map[string]any{}

type Storage interface {
	Get(key string) (any, error)
	Set(key string, value string) error
}
type MemoryStorage struct{}

func (m *MemoryStorage) Get(key string) (any, error) {
	fmt.Printf("the get method key is %s\n", key)
	return cache[key], nil
}
func (m *MemoryStorage) Set(key string, value string) error {
	fmt.Printf("the set method key is %s and value is %s\n", key, value)
	cache[key] = value
	return nil
}

type Client struct {
	storage Storage
}

//
//func main() {
//	s := &Client{&MemoryStorage{}}
//	s.storage.Set("k1", "v1")
//	val, err := s.storage.Get("k1")
//	if err == nil {
//		fmt.Println("k1 -> value is ", val)
//	}
//}
