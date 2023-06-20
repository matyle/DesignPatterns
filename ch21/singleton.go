package ch21

import (
	"fmt"
	"sync"
)

type Singleton[T any] struct {
	instance *Singleton[T]
	value    T
	lock     *sync.Mutex
}

// private constructor
func newSingleton[T any]() *Singleton[T] {
	return &Singleton[T]{
		lock: &sync.Mutex{},
	}
}

// public method
func (s *Singleton[T]) GetInstance() *Singleton[T] {
	// Double check
	if s.instance == nil {
		s.lock.Lock()
		//check again
		if s.instance == nil {
			s.instance = newSingleton[T]()
		}
		s.lock.Unlock()
	}
	return s.instance
}

func SingletonMain() {
	s := newSingleton[int]()
	s1 := s.GetInstance()
	s2 := s.GetInstance()
	if s1 == s2 {
		fmt.Println("s1 and s2 are same instance")
	}
}
