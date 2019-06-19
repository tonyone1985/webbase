// test.go
package main

import (
	"sync"
)

type ConcurrentHashMap interface {
	Set(k interface{}, v interface{})
	Get(k interface{}) interface{}
}

type _ConcurrentHashMap struct {
	l *sync.Mutex
	m map[interface{}]interface{}
}

func (this *_ConcurrentHashMap) Set(k interface{}, v interface{}) {
	this.l.Lock()
	defer this.l.Unlock()
	this.m[k] = v

}
func (this *_ConcurrentHashMap) Get(k interface{}) interface{} {
	this.l.Lock()
	defer this.l.Unlock()
	v, ok := this.m[k]
	if !ok {
		return nil
	}
	return v
}
func (this *_ConcurrentHashMap) Delete(k interface{}) {
	this.l.Lock()
	defer this.l.Unlock()
	delete(this.m, k)
}

func NewConcurrentHashMap() ConcurrentHashMap {
	return &_ConcurrentHashMap{
		m: make(map[interface{}]interface{}),
		l: new(sync.Mutex),
	}

}
