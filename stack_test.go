package fund

import (
    "testing"
)

func TestPush(t *testing.T) {
    s := new(Stack)
    in := []int{1,2,3,4,5,6,7,8,9,10}
    for i := 0; i < 10; i++ {
        err := s.Push(in[i])
        if err != nil {
            t.Errorf("Cannot push value %d", in[i]);
        }
    }
}

func TestPop(t *testing.T) {
    s := new(Stack)
    in := []int{1,2,3,4,5,6,7,8,9,10}
    out := []int{10,9,8,7,6,5,4,3,2,1}
    for i := 0; i < 10; i++ {
        err := s.Push(in[i])
        if err != nil {
            t.Errorf("Cannot push value %d", in[i]);
        }
    }
    for i := 0; i< 10; i++ {
        val, err := s.Pop()
        if val != out[i] {
            t.Errorf("Popped value %d does not agree with expected value %d", val, out[i])
        }
        if err != nil {
            t.Errorf("Cannot pop %d. value", i)
        }
    }
}

