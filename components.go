package godux

import (
	"github.com/gopherjs/vecty"
)

type Component interface {
	vecty.Component

	Connect() map[interface{}]interface{}
}
