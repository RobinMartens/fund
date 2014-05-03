package fund

import (
    "errors"
)

type Elm struct {
    val int
    rest *Elm
}

type Stack struct {
    top *Elm
}

func (s *Stack) Push(val int) error {
    e := new(Elm)
    e.val = val
    e.rest = s.top
    s.top = e
    return nil
}

func (s *Stack) Pop() (int, error) {
    e := s.top
    if e == nil {
       return 0, errors.New("Stack empty!")
    }
    s.top = e.rest
    return e.val, nil
}
