package debouncer

import (
	"sync"
	"time"
)

type Debouncer[T comparable] struct {
	duration time.Duration
	timers   map[T]*time.Timer
	mu       sync.Mutex
}

func NewDebouncer[T comparable](d time.Duration) *Debouncer[T] {
	return &Debouncer[T]{
		timers:   make(map[T]*time.Timer),
		duration: d,
	}
}

func (d *Debouncer[T]) Debounce(key T, fn func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if timer, exists := d.timers[key]; exists {
		timer.Stop()
	}

	timer := time.AfterFunc(d.duration, func() {
		fn()

		d.mu.Lock()
		delete(d.timers, key)
		d.mu.Unlock()
	})

	d.timers[key] = timer
}

func (d *Debouncer[T]) Cancel(key T) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if timer, exists := d.timers[key]; exists {
		timer.Stop()
		delete(d.timers, key)
	}
}
