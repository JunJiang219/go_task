package tools

type Stack struct {
	data []interface{}
}

func (s *Stack) Length() int {
	return len(s.data)
}

func (s *Stack) Push(val interface{}) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	length := len(s.data)
	if length == 0 {
		return nil, false
	}

	val := s.data[length-1]
	s.data = s.data[:length-1]
	return val, true
}
