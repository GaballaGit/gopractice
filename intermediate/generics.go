package intermediate

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[A any] struct {
	elements []A
}

func (s *Stack[A]) push(elem A) {
	s.elements = append(s.elements, elem)
}

func (s *Stack[A]) pop() (A, bool) {
	if len(s.elements) == 0 {
		var zero A
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[A]) isEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	intStack := Stack[int]{}
	intStack.push(5)
	intStack.push(8)
	intStack.push(3)
	intStack.pop()
	fmt.Println(intStack.elements)

}
