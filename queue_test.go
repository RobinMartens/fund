package fund

import (
    "testing"
)

func TestEnqueue(t *testing.T) {
    q := new(Queue)
    in := []int{5,4,3,7,43,234,654,6,231,2}
    for i := 0; i < 10; i++ {
        err := q.Enqueue(in[i])
        if err != nil {
            t.Errorf("Cannot enqueue value %d", in[i])
        }
    }
}

func TestDequeue(t *testing.T) {
    q := new(Queue)
    in := []int{5,4,3,7,43,234,654,6,231,2}
    for i := 0; i < 10; i++ {
        err := q.Enqueue(in[i])
        if err != nil {
            t.Errorf("Cannot enqueue value %d", in[i])
        }
    }
    for i := 0; i < 10; i++ {
        val, err := q.Dequeue()
        if err != nil {
            t.Errorf("Cannot dequeue %d. value")
        }
        if val != in[i] {
            t.Errorf("Dequeued value %d does not agree with expected value %d", val, in[i])
        }
    }
}
