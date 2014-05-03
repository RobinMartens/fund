package fund

type Elm struct {
    val int
    rest *Elm
}

type Stack struct {
    top *Elm
}

func (s *Stack) Push(val int) {
    e := new(Elm)
    e.val = val
    e.rest = s.top
    s.top = e
}

func (s *Stack) Pop() int {
    e := s.top
    s.top = e.rest
    return e.val
}
