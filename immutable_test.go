package queue_test

import (
	"testing"

	queue "github.com/lobotomist/immutable-queue"
)

func TestImmutableQueue(t *testing.T) {
	q1 := queue.NewImmutableQueue()
	q2 := q1.EnQueue(1)
	q3 := q2.EnQueue(2)

	if _, err := q2.Head(); err != queue.ErrEmptyQueue {
		t.Error("must be an error")
	}
	{
		q2 = q2.DeQueue()
		assertHeadInQueue(t, *q2, 1)
	}

	{
		q3 = q3.DeQueue()
		assertHeadInQueue(t, *q3, 1)
		q3 = q3.DeQueue()
		assertHeadInQueue(t, *q3, 2)

		q3 = q3.DeQueue()
		if _, err := q3.Head(); err != queue.ErrEmptyQueue {
			t.Error("must be an error")
		}
	}
}

// Running tool: /usr/local/bin/go test -benchmem -run=^$ training/basic/queue -bench ^(BenchmarkImmutableStack10Ops)$ -v

// goos: darwin
// goarch: amd64
// pkg: training/basic/queue
// BenchmarkImmutableStack10Ops-4   	  707620	      1717 ns/op	    1776 B/op	      35 allocs/op
// PASS
// ok  	training/basic/queue	1.244s
// Success: Benchmarks passed.

func BenchmarkImmutableStack10Ops(b *testing.B) {
	q := queue.NewImmutableQueue()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			q = q.EnQueue(1)
		}
		for j := 0; j < 10; j++ {
			q = q.DeQueue()
		}
	}
}

func assertHeadInQueue(t *testing.T, q queue.ImmutableQueue, expected int) {
	t.Helper()
	if v, err := q.Head(); err == nil {
		if v.(int) == expected {
			return
		}
	}
	t.Fail()
}
