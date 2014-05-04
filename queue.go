package fund

import (
    "errors"
)

type Queue struct {
    in Stack
    out Stack
    n_in int
    n_out int
}

func (q *Queue) Enqueue(val int) error {
    err := q.in.Push(val)
    q.n_in++
    if err != nil {
        return errors.New("Cannot push value to in stack")
    }
    return nil
}

func (q *Queue) Dequeue() (int, error) {
    if q.n_out == 0 {
        for ; q.n_in > 0; q.n_in-- {
            e, inErr := q.in.Pop()
            outErr := q.out.Push(e)
            q.n_out++
            if inErr != nil {
                return 0, errors.New("Cannot pop from in stack")
            }
            if outErr != nil {
                return 0, errors.New("Cannot push to out stack")
            }
        }
    }
    val, err := q.out.Pop()
    q.n_out--
    if err != nil {
        errors.New("Cannot pop from out stack")
    }
    return val, nil
}

