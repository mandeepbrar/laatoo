package commonobjects

import (
	"fmt"
	"github.com/pmylund/go-cache"
)

type MemoryStorer struct {
	items *cache.Cache
	name  string
}

func NewMemoryStorer(name string) *MemoryStorer {
	memStor := &MemoryStorer{name: name}
	memStor.items = cache.New(cache.NoExpiration, 0)
	return memStor
}

func (ms *MemoryStorer) GetName() string {
	return ms.name
}

func (ms *MemoryStorer) Put(id string, item interface{}) error {
	ms.items.Set(id, item, cache.NoExpiration)
	return nil
}

func (ms *MemoryStorer) Get(conditions interface{}) (interface{}, error) {
	return nil, fmt.Errorf("Not Implemented")
}

func (ms *MemoryStorer) GetById(id string) (interface{}, error) {
	item, ok := ms.items.Get(id)
	if !ok {
		return nil, fmt.Errorf("Object %s not found", id)
	}
	return item, nil
}

func (ms *MemoryStorer) Delete(id string) error {
	ms.items.Delete(id)
	return nil
}

func (ms *MemoryStorer) GetList() (interface{}, error) {
	cItems := ms.items.Items()
	items := make([]interface{}, len(cItems))
	idx := 0
	for _, value := range cItems {
		items[idx] = value
		idx++
	}
	return items, nil
}
