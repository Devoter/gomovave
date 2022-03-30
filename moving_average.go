package gomovave

import "errors"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// ErrMovingAverageQueueIsEmpty means that the front value of the moving average cannot be read, because the queue is empty.
var ErrMovingAverageQueueIsEmpty = errors.New("Moving average queue is empty")

// MovingAverage implementation.
type MovingAverage[T Number] struct {
	maxLength int
	dirty     bool
	queue     []T
	value     T
}

// NewMovingAverage returns an instance of MovingAverage with default private fields.
func NewMovingAverage[T Number](maxLength int) *MovingAverage[T] {
	return &MovingAverage[T]{maxLength: maxLength, queue: []T{}}
}

// Value returns the current moving average value.
// This method uses computed value.
func (ma *MovingAverage[T]) Value() T {
	if ma.dirty {
		sum := T(0)

		for _, v := range ma.queue {
			sum += v
		}

		ma.value = sum / T(len(ma.queue))
		ma.dirty = false
	}

	return ma.value
}

// Len returns the current queue length.
func (ma *MovingAverage[T]) Len() int {
	return len(ma.queue)
}

// MaxLen returns the maximum queue length.
func (ma *MovingAverage[T]) MaxLen() int {
	return ma.maxLength
}

// Front returns a front value of the queue.
func (ma *MovingAverage[T]) Front() (T, error) {
	if len(ma.queue) == 0 {
		return 0, ErrMovingAverageQueueIsEmpty
	}

	return ma.queue[0], nil
}

// Queue returns the queue slice.
func (ma *MovingAverage[T]) Queue() []T {
	return ma.queue
}

// Push appends a value to the queue.
func (ma *MovingAverage[T]) Push(value T) {
	if len(ma.queue) >= ma.maxLength {
		ma.queue = append(ma.queue[1:], value)
	} else {
		ma.queue = append(ma.queue, value)
	}

	ma.dirty = true
}

// Clear resets the instance to the initial state.
func (ma *MovingAverage[T]) Clear() {
	ma.queue = []T{}
	ma.value = 0
	ma.dirty = false
}
