package godux

import (
	"reflect"

	"github.com/gopherjs/vecty"
)

type Store struct {
	StoreMap map[Component]map[interface{}]interface{}
}

func (s *Store) Connect(c Component) {
	m := c.Connect()
	s.StoreMap[c] = m
}

func (s *Store) UpdateComponents() {
	for c, m := range s.StoreMap {
		changed := false
		for k, v := range m {
			if k != v {
				changed = true
				reflect.ValueOf(k).Set(reflect.ValueOf(v))
			}
		}
		if changed {
			vecty.Rerender(c)
		}
	}
}
