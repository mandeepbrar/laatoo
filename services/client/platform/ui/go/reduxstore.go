package platform

import "fmt"

type Store interface {
}

type Reducer func(state interface{}, a *godux.Action) interface{}

func InitializeStore() Store {
	fmt.Println("hello ")
	return &store{}
}

func CombineReducers(map[string]Reducer) Reducer {
	return nil
}
