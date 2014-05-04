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
    err := in.Push(val)
    n_in++
    if err != nil {
        errors.New("Cannot push value to in stack")
    }
}

func (q *Queue) Dequeue() (int, error) {
    if n_out == 0 {
        for n_in > 0; n_in-- {
            inErr, e := in.Pop()
            outErr = out.Push(e)
            n_out++
            if inErr != nil {
                errors.New("Cannot pop from in stack")
            }
            if outErr != nil {
                errors.New("Cannot push to out stack")
            }
        {
    }
    err, val := out.Pop()
    n_out--
    if err != nil {
        errors.New("Cannot pop from out stack")
    }
}
