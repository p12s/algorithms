package parentheses

const (
	OPEN_BRACKET    = '('
	CLOSING_BRACKET = ')'
)

type Stack struct {
	depth int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push() {
	s.depth += 1
}

func (s *Stack) Pop() {
	if s.depth > 0 {
		s.depth -= 1
	}
}

func (s *Stack) IsEmpty() bool {
	return s.depth == 0
}

func IsParenthesesValid(input string) bool {
	stack := NewStack()
	for _, char := range []rune(input) {
		switch char {
		case OPEN_BRACKET:
			stack.Push()
		case CLOSING_BRACKET:
			stack.Pop()
		default:
			break
		}
	}

	return stack.IsEmpty()
}
