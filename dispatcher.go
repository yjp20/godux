package godux

type Handler func(interface{})

type Dispatcher struct {
	counter   int
	callbacks map[int]Handler
}

func (d *Dispatcher) Init() {
	d.callbacks = make(map[int]Handler)
}

func (d *Dispatcher) Register(callback Handler) int {
	d.counter++
	id := d.counter
	d.callbacks[id] = callback

	return id
}

func (d *Dispatcher) Dispatch(action interface{}) {
	for _, c := range d.callbacks {
		c(action)
	}
}

func (d *Dispatcher) Unregister(id int) {
	delete(d.callbacks, id)
}
