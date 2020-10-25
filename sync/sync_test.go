package sync_

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCounter := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCounter)

		for i := 0; i < wantedCounter; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCounter)

	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
