package fund

import (
    "testing"
)

func TestStack(t *testing.T) {

    in0 := 13
    in1 := 42
    in2 := 273

    s := new(Stack)
    _ = s.Push(in0)
    _ = s.Push(in1)
    _ = s.Push(in2)

    out0, _ := s.Pop()
    out1, _ := s.Pop()
    out2, _ := s.Pop()

    if in0 != out2 || in1 != out1 || in2 != out0 {
       t.Errorf("Error!")
    }


}
