package main

type Slice struct {
	items []int
}

func (s *Slice) Enqueuee(v int) {
	s.items = append(s.items, v)
}

func (s *Slice) Dequeuee() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	v := s.items[0]
	s.items = s.items[1:]
	return v, true
}

func main()  {
	slq := Slice{}

	slq.Enqueuee(5)
	
}