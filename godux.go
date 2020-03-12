package godux

import (
	"reflect"

	"github.com/gopherjs/vecty"
)

type Component interface {
	vecty.Component
	Connect(*Store) map[interface{}]interface{}
}

type componentWrapper struct {
	Component
	store *Store
}

func (c *componentWrapper) Mount() {
	c.store.subscribe(c)
	if mounter, ok := c.Component.(vecty.Mounter); ok {
		mounter.Mount()
	}
}

func (c *componentWrapper) Unmount() {
	c.store.unsubscribe(c)
	if mounter, ok := c.Component.(vecty.Unmounter); ok {
		mounter.Unmount()
	}
}

type Store struct {
	State     interface{}
	storeMap  map[Component]map[interface{}]interface{}
	counter   int
	callbacks map[int]func(interface{})
}

func (s *Store) Init() {
	s.storeMap = map[Component]map[interface{}]interface{}{}
}

func (s *Store) Register(callback Handler) int {
	s.counter++
	id := s.counter
	s.callbacks[id] = callback

	return id
}

func (s *Store) Unregister(id int) {
	delete(s.callbacks, id)
}

func (s *Store) Dispatch(action interface{}) {
	for _, c := range s.callbacks {
		c(action)
	}
}

func (s *Store) Connect(c Component) vecty.Component {
	println("connected", c, reflect.TypeOf(c).Elem().Name())
	return &componentWrapper{
		Component: c,
		store:     s,
	}
}

func (s *Store) updateComponents() {
	for c, m := range s.storeMap {
		changed := false
		for k, v := range m {
			if reflect.ValueOf(k).Interface() != reflect.ValueOf(v).Interface() {
				changed = true
				reflect.Indirect(reflect.ValueOf(k)).Set(reflect.Indirect(reflect.ValueOf(v)))
			}
		}
		if changed {
			println("rerender", c)
			vecty.Rerender(c)
		}
	}
}

func (s *Store) subscribe(c Component) {
	m := c.Connect(s)
	s.storeMap[c] = m
}

func (s *Store) unsubscribe(c Component) {
	delete(s.storeMap, c)
}
