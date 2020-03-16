package main 

type CustomStack struct {
	e []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		e: make([]int, 0, maxSize)
	}
}

func (s *CustomStack) Push(x int) {
	if len(s.e) >= cap(s.e) {
		return
	}
	s.e = append(s.e, x)
}

func (s *CustomStack) Pop() int {
	if len(s.e) == 0 {
		return -1
	}
	el := s.e[len(s.e)-1]
	s.e = s.e[0:len(s.e)-1]
	return el
}

func (s *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(s.e); i++ {
		s.e[i] += val
	}
}